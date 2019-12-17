const Caver = require('caver-js');
const caver = new Caver(process.env.NETWORK_PROVIDER_URL);
const _ = require('lodash');
exports.handler = (event, context, callback) => {
  let accounts = [];
  const number = _.get(event, 'number', 1);
  for (var i = 0; i < number; i++) {
    const account = caver.klay.accounts.create();
    accounts.push(account);
  }
  callback(null,accounts);
};
