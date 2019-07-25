'use strict'
const Caver = require('caver-js')
const AWS = require('aws-sdk')
const EchoAPI = require('../abi/contracts/APIv100.json')
const caver = new Caver(process.env.NETWORK_PROVIDER_URL)

let account

const echoAPI = new caver.klay.Contract(
  EchoAPI.abi,
  process.env.API_CONTRACT_ADDRESS
)
var ssm = new AWS.SSM({
  apiVersion: '2014-11-06'
})

setSK2Account()

class Handler {
  constructor() {
    this._echo_api = echoAPI.methods
    this._body = {}
    this._message = ''
    this._status = 200
    this._is_error = false
    this._error_message = 'error'
    this._account = account
  }

  getEchoAPI() {
    return this._echo_api
  }

  getAccount() {
    return this._account
  }

  getCallerAddress() {
    return this._account.address
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
      const errBody = { message: h.message }
      if (process.env.NODE_ENV === 'test') errBody.stack = h.stack
      res.status(400).json(errBody)
      return
    }
    res.status(500).json({ message: 'unknown error occur.' })
  }
}

function setSK2Account() {
  const options = { Name: 'SK_ECHO_WALLET' }
  return new Promise((resolve, reject) => {
    ssm.getParameter(options, (err, data) => {
      if (err) {
        reject(err)
        return
      }
      account = caver.klay.accounts.wallet.add(data.Parameter.Value)
      resolve()
    })
  })
}

module.exports = Handler
