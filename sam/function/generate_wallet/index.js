const Caver = require('caver-js');
const AWS = require('aws-sdk');
const caver = new Caver(
  new Caver.providers.HttpProvider(process.env.NETWORK_PROVIDER_URL, {
    timeout: 44000
  })
);
const EchoAPI = require('./APIv102.json');
const _ = require('lodash');
const echoAPI = new caver.klay.Contract(
  EchoAPI.abi,
  process.env.API_CONTRACT_ADDRESS
);

var ssm = new AWS.SSM({
  apiVersion: '2014-11-06'
});

let account;
let feePayer;

function setFeePayerAccount() {
  const options = { Name: 'SK_ECHO_FEE_PAYER_ACCOUNT', WithDecryption: true };
  const options2 = { Name: 'SK_ECHO_FEE_PAYER_PRIVATE', WithDecryption: true };
  const options3 = { Name: 'SK_ECHO_WALLET', WithDecryption: true };
  if (account && feePayer) {
    return new Promise((resolve, reject) => {
      resolve([account, feePayer]);
    });
  }
  const feeaccount = new Promise((resolve, reject) => {
    ssm.getParameter(options, (err, data) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(data.Parameter.Value);
    });
  }).catch(e => {
    return e.code;
  });
  const privateKey = new Promise((resolve, reject) => {
    ssm.getParameter(options2, (err, data) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(data.Parameter.Value);
    });
  }).catch(e => {
    return e.code;
  });

  const defaultAccount = new Promise((resolve, reject) => {
    ssm.getParameter(options3, (err, data) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(data.Parameter.Value);
    });
  });
  return Promise.all([feeaccount, privateKey, defaultAccount]).then(values => {
    console.log('@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@');
    caver.klay.accounts.wallet.clear();
    caver.klay.accounts.wallet.add(values[1], values[0]);
    account = caver.klay.accounts.wallet.add(values[2]);
    feePayer = values[0];
    return [account, values[0]];
  });
}

exports.handler = (event, context, callback) => {
  const number = _.get(event, 'number', 1);
  setFeePayerAccount().then(
    data => {
      const defaultAccount = data[0];
      const feepayer = data[1];
      const promises = [];

      caver.klay.getTransactionCount(defaultAccount.address).then(nonce => {
        for (var i = 0; i < number; i++) {
          const account = caver.klay.accounts.create();
          const data = echoAPI.methods.addWriter(account.address).encodeABI();
          const klayRequest = caver.klay.accounts
            .signTransaction(
              {
                type: 'FEE_DELEGATED_SMART_CONTRACT_EXECUTION',
                from: defaultAccount.address,
                to: process.env.API_CONTRACT_ADDRESS,
                data,
                gas: 10000000,
                nonce: nonce + i
              },
              defaultAccount.privateKey
            )
            .then(
              result => {
                const { rawTransaction: senderRawTransaction } = result;
                return caver.klay
                  .sendTransaction({
                    senderRawTransaction: senderRawTransaction,
                    feePayer: feepayer
                  })
                  .then(
                    result => {
                      return account;
                    },
                    err => {
                      console.log(err);
                      return err;
                    }
                  );
              },
              err => {
                return err;
              }
            );
          promises.push(klayRequest);
        }
        Promise.all(promises).then(accounts => {
          callback(null, accounts);
        });
      });
    },
    error => {
      callback(error);
    }
  );
};
