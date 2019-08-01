const Joi = require('joi')
const Handler = require('../middleware/handler.middleware')

class Function {
  constructor(body, callback) {
    const [
      {
        id,
        title,
        author,
        origin,
        lang,
        paperback,
        publisher_id: publisherId,
        publish_date: publishDate
      }
    ] = [body]
    new Handler().then(handler => {
      const echo = handler.getEchoAPI()
      const callerAddr = handler.getCallerAddress()
      echo
        .addBook(
          id,
          title,
          author,
          origin,
          lang,
          paperback,
          publisherId,
          publishDate
        )
        .send({ from: callerAddr, gas: 2000000 })
        .then(r => {
          handler.setResponseBody(r).setStatusCode(200)
          callback(null, handler)
        })
        .catch(err => {
          handler.setErrorMessage(err)
          callback(handler)
        })
    })
  }

  static schema(body) {
    const schema = Joi.object().keys({
      id: Joi.string().required(),
      title: Joi.string().required(),
      author: Joi.string().required(),
      origin: Joi.string().required(),
      lang: Joi.string().required(),
      paperback: Joi.number().required(),
      publisher_id: Joi.string().required(),
      publish_date: Joi.date()
        .timestamp('unix')
        .required()
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
