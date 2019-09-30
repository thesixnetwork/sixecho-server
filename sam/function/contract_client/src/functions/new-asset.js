const Joi = require('joi')
const Handler = require('../middleware/handler.middleware')

class Function {
  constructor(body, callback) {
    const [{ hash, block_number: blockNo }] = [body]
    new Handler().then(handler => {
      const echo = handler.getEchoAPI()
      const callerAddr = handler.getCallerAddress()
      echo
        .addAsset(hash, blockNo)
        .send({ from: callerAddr, gas: 10000000 })
        .then(r => {
          handler.setResponseBody(r).setStatusCode(200)
          callback(null, handler)
        })
        .catch(err => {
          console.error(err)
          handler.setErrorMessage(err)
          callback(handler)
        })
    })
  }

  static schema(body) {
    const schema = Joi.object().keys({
      hash: Joi.string().required(),
      block_number: Joi.string().required()
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
