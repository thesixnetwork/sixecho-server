const express = require('express')
const router = express.Router()

// routers
const WalletRoute = require('./wallet.route')

// report path
router.use('/wallet', WalletRoute)

module.exports = router
