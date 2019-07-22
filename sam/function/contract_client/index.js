'use strict'
const awsServerlessExpress = require('aws-serverless-express')
const app = require('./src/app')

if (process.env.NODE_ENV === 'test') {
  const port = 3000
  // eslint-disable-next-line no-console
  app.listen(port, () =>
    console.log(`Debugging app listening on port ${port}!`)
  )
} else {
  const server = awsServerlessExpress.createServer(app)

  exports.handler = (event, context) =>
    awsServerlessExpress.proxy(server, event, context)
}
