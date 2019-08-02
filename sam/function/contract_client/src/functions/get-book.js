const _ = require('underscore')
const Joi = require('joi')
const Handler = require('../middleware/handler.middleware')

class Function {
  constructor(body, callback) {
    new Handler().then(handler => {
      const echo = handler.getEchoAPI()
      echo
        .getBook(body.id)
        .call()
        .then(r =>
          echo
            .getAdditionalBookData(body.id)
            .call()
            .then(additionalInfo =>
              _.extend(
                filterOutNumberKeyInObject(r),
                filterOutNumberKeyInObject(additionalInfo)
              )
            )
        )
        .then(r => {
          handler.setResponseBody(r).setStatusCode(200)
          callback(null, handler)
          return
        })
        .catch(err => {
          handler.setErrorMessage(err)
          callback(handler)
          return
        })
    })
  }

  static schema(body) {
    const schema = Joi.object().keys({
      id: Joi.string().required()
    })
    // Return result.
    const result = Joi.validate(body, schema)
    if (result.error === null) {
      return
    }

    return result
  }
}

function filterOutNumberKeyInObject(r) {
  const buildResp = {}
  _.keys(r)
    .filter(key => isNaN(parseInt(key)))
    .forEach(key => {
      buildResp[key] = r[key]
    })
  return buildResp
}

module.exports = Function
