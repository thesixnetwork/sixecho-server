'use strict'

class Handler {
  constructor() {
    this._body = {}
    this._message = ''
    this._status = 200
    this._is_error = false
    this._error_message = 'error'
  }

  setResponseBody(body) {
    this._body = body
    return this
  }

  setMessage(body) {
    this._message = body
    return this
  }

  setStatusCode(body) {
    this._status = body
    if (body > 300) this._is_error = true

    return this
  }

  setErrorMessage(body) {
    this._error_message = body
    this._is_error = true
    if (this._status < 300) this._status = 400
    return this
  }

  // eslint-disable-next-line no-unused-vars
  static Response(h, req, res, next) {
    if (h instanceof Handler) {
      if (h._is_error === true) {
        res.status(h._status).json({ message: h._error_message })
        return
      }
      const respBody = { body: h._body }
      if (h._message.length > 0) respBody.message = h._message
      res.status(h._status).json(respBody)
      return
    } else if (h instanceof Error) {
      res.status(400).json({ message: h.message })
      return
    }
    res.status(500).json({ message: 'unknown error occur.' })
  }
}

module.exports = Handler