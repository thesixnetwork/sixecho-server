'use strict'
const Handler = require('./src/middleware/handler.middleware')
const functions = {
  'new-book': require('./src/functions/new-book'),
  'new-book-and-digest': require('./src/functions/new-book-and-digest'),
  'get-book': require('./src/functions/get-book'),
  'get-digest': require('./src/functions/get-digest'),
  'new-digest': require('./src/functions/new-digest')
}

exports.handler = (event, context, callback) => {
  context.callbackWaitsForEmptyEventLoop = false
  const name = event.name
  const body = event.body || {}
  if (functions[name]) {
    const Fn = functions[name]
    const invalid = Fn.schema(body)
    if (invalid) {
      callback(invalid.error)
      return
    }
    new Fn(body, (err, resp) => {
      if (err) {
        callback(Handler.Response(err))
        return
      }
      callback(null, Handler.Response(resp))
    })
  } else {
    callback(new Error(`No function name ${name}.`))
  }
}
