const express = require('express')
const router = express.Router()

// routers
const WalletRoute = require('./wallet.route')
const EchoRoute = require('./echo.route')

// report path
router.use('/wallet', WalletRoute)
router.use('/echo', EchoRoute)

module.exports = router
