const express = require('express')
const cors = require('cors')
const awsServerlessExpressMiddleware = require('aws-serverless-express/middleware')
const bodyParser = require('body-parser')
const Handler = require('./middleware/handler.middleware')
const debug = require('debug')('http')

// routers
const router = require('./router/index')

// libs

const app = express()

app.use(cors())

// middleware

app.use(bodyParser.json())
app.use(bodyParser.urlencoded({ extended: true }))
app.use(awsServerlessExpressMiddleware.eventContext())

app.use((req, res, next) => {
  debug(req.method + ' ' + req.url)
  debug(req.body)
  next()
})

app.use('/', router)

app.use('/*', (req, res, next) => {
  const handler = new Handler()
  handler.setErrorMessage('Error: path is not exists.').setStatusCode(404)
  next(handler)
})

app.use(Handler.Response)

module.exports = app
