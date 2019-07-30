const Joi = require('joi')
const Handler = require('../middleware/handler.middleware')

class Function {
  constructor(body, callback) {
    const handler = new Handler()
    const echo = handler.getEchoAPI()
    const callerAddr = handler.getCallerAddress()
    echo
      .uploadDigest(body.id, body.digest)
      .send({ from: callerAddr, gas: 2000000 }, (err, txID) => {
        if (err) {
          handler.setErrorMessage(err)
          callback(handler)
          return
        }
        handler.setResponseBody(txID).setStatusCode(200)
        callback(null, handler)
        return
      })
  }

  static schema(body) {
    const schema = Joi.object().keys({
      id: Joi.string().required(),
      digest: Joi.array().required()
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
