const Handler = require('../middleware/handler.middleware')

class EchoCtrl {
  static newBook(req, res, next) {
    const [
      id,
      {
        title,
        author,
        origin,
        lang,
        paperback,
        publisher_id: publisherId,
        publish_date: publishDate
      }
    ] = [req.params.id, req.body]
    const handler = new Handler()
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
      .send({ from: callerAddr, gas: 6721975 }, (err, txID) => {
        if (err) {
          handler.setErrorMessage(err).setStatusCode(400)
        } else {
          handler.setResponseBody(txID).setStatusCode(200)
        }
        next(handler)
      })
  }

  static getBook(req, res, next) {
    const id = req.params.id
    const handler = new Handler()
    const echo = handler.getEchoAPI()
    echo
      .getBook(id)
      .call()
      .then(r => {
        handler.setResponseBody(r).setStatusCode(200)
        next(handler)
      })
      .catch(err => {
        handler.setErrorMessage(err).setStatusCode(200)
        next(handler)
      })
  }

  static getDigest(req, res, next) {
    //TODO @IMPLEMENT this
  }

  static newDigest(req, res, next) {
    //TODO @IMPLEMENT this
  }
}

module.exports = EchoCtrl
