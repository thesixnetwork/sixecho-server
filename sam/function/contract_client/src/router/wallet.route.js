const express = require('express')
const Web3 = require('web3')
const web3 = new Web3(process.env.NETWORK_PROVIDER_URL)
const router = express.Router()
const Handler = require('../middleware/handler.middleware')

router.route('/accounts').get((req, res, next) => {
  const handler = new Handler()
  web3.eth
    .getAccounts()
    .then(e => {
      handler.setResponseBody(e).setStatusCode(200)
      next(handler)
    })
    .catch(e => {
      handler.setErrorMessage(e.message).setStatusCode(400)
      next(handler)
    })
})

module.exports = router
