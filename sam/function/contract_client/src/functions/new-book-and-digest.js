const Joi = require('joi')
const NewBook = require('./new-book')
const NewDigest = require('./new-digest')
const Handler = require('../middleware/handler.middleware')

class Function {
  constructor(body, callback) {
    const newBook = new Promise((resolve, reject) => {
      new NewBook(body, (handlerError, handler) => {
        if (handlerError) {
          reject(handlerError.getErrorMessage())
          return
        }
        resolve(handler.getResponseBody())
      })
    })
    const newDigest = new Promise((resolve, reject) => {
      new NewDigest(body, (handlerError, handler) => {
        if (handlerError) {
          reject(handlerError.getErrorMessage())
          return
        }
        resolve(handler.getResponseBody())
      })
    })
    Promise.all([newBook, newDigest])
      .then(resps => {
        new Handler().then(handler => {
          handler.setResponseBody(resps).setStatusCode(200)
          callback(null, handler)
        })
      })
      .catch(e => {
        callback(e)
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
        .required(),
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
