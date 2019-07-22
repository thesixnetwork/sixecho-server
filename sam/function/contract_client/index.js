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

/*
const Web3 = require('web3')
const web3 = new Web3(process.env.NETWORK_PROVIDER_URL)

exports.handler = async (event) => {
  // TODO implement

  web3.eth.getAccounts().then(e => { 
    console.log("len: " + e.length);
   }) 

  const response = {
      statusCode: 200,
      body: JSON.stringify('Hello from Lambda!'),
  };
  return response;
};
*/
