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

class Handler {
  constructor() {
    this._echo_api = echoAPI.methods
    this._body = {}
    this._message = ''
    this._status = 200
    this._is_error = false
    this._error_message = 'error'
    if (account) {
      this._account = account
      return Promise.resolve(this)
    } else {
      return setSK2Account().then(acc => {
        this._account = acc
        return Promise.resolve(this)
      })
    }
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
  static Response(h) {
    if (h instanceof Handler) {
      if (h._is_error === true) {
        return {
          status: h._status,
          message: h._error_message
        }
      }
      return {
        status: h._status || 200,
        message: h._message,
        body: h._body
      }
    } else if (h instanceof Error) {
      const errBody = { message: h.message }
      errBody.status = 400
      if (process.env.NODE_ENV === 'test') errBody.stack = h.stack
      return errBody
    }
    return {
      status: 500,
      message: 'unknown error occur.'
    }
  }
}

function setSK2Account() {
  const options = { Name: 'SK_ECHO_WALLET', WithDecryption: true }
  return new Promise((resolve, reject) => {
    ssm.getParameter(options, (err, data) => {
      if (err) {
        reject(err)
        return
      }
      caver.klay.accounts.wallet.clear()
      account = caver.klay.accounts.wallet.add(data.Parameter.Value)
      resolve(account)
    })
  })
}

module.exports = Handler
