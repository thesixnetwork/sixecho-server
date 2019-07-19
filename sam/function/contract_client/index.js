const Web3 = require('web3')
const web3 = new Web3(process.env.NETWORK_PROVIDER_URL)

exports.handler = async (event) => {
  // TODO implement
  const response = {
      statusCode: 200,
      body: JSON.stringify('Hello from Lambda!'),
  };
  return response;
};