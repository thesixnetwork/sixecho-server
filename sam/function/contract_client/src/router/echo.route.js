const express = require('express')
const router = express.Router()
const EchoController = require('../controller/echo.controller')

router.route('/book/:id').get(EchoController.getBook)
router.route('/book/:id').post(EchoController.newBook)

router.route('/digest/:id').get(EchoController.getDigest)
router.route('/digest/:id').post(EchoController.newDigest)

module.exports = router
