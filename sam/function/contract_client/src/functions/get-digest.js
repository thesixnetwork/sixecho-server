const Joi = require('joi')
const Handler = require('../middleware/handler.middleware')

class Function {
  constructor(body, callback) {
    const handler = new Handler()
    const echo = handler.getEchoAPI()
    echo
      .downloadDigest(body.id)
      .call()
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

module.exports = Function
