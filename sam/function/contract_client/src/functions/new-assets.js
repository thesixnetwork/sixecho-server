const Joi = require('joi');
const Promise = require('bluebird');
const Handler = require('../middleware/handler.middleware');
const AWS = require('aws-sdk');

class Function {
  constructor(body, callback) {
    new Handler().then(handler => {
      const kms = new AWS.KMS();
      const echo = handler.getEchoAPI();
      const caver = handler.getCaverAPI();
      let { address: senderAddress } = handler.getAccountDefault();
      const accounts = {};

      // set default account to find nonce
      accounts[senderAddress] = 0;
      const kmsPromises = [];
      for (let i = 0; i < body.length; i++) {
        accounts[body[i].account] = 0;
        let a = new Promise(function(resolve, reject) {
          if (body[i].private_key === '') {
            resolve('');
          }
          const decryptBase64 = Buffer.from(body[i].private_key, 'base64');
          var params = {
            CiphertextBlob: decryptBase64
          };
          kms.decrypt(params, function(err, data) {
            if (err) {
              reject(err);
              return;
            }
            resolve(Buffer.from(data.Plaintext).toString());
          });
        });
        kmsPromises.push(a);
      }

      Promise.all(kmsPromises).then(result => {
        result.forEach((data, index) => {
          body[index].private_key = data;
        });

        const getNonces = [];
        Object.keys(accounts).forEach((account, index) => {
          if (account !== '') {
            let a = caver.klay.getTransactionCount(account).then(nonce => {
              return { nonce, account };
            });
            getNonces.push(a);
          }
        });
        Promise.all(getNonces).then(result => {
          result.forEach(nonce => {
            accounts[nonce.account] = nonce.nonce;
          });

          const promises = [];
          for (let i = 0; i < body.length; i++) {
            let {
              privateKey: senderPrivateKey,
              address: senderAddress
            } = handler.getAccountDefault();
            if (body[i].private_key !== '') {
              senderPrivateKey = body[i].private_key;
              senderAddress = body[i].account;
            }
            // console.log('-----------------------');
            // console.log(senderAddress);
            // console.log(senderPrivateKey);
            // console.log(handler.getContractAddress());
            // console.log('-----------------------');

            function getCurrentNonce(account) {
              const tmp = accounts[account];
              accounts[account] = accounts[account] + 1;
              return tmp;
            }
            const data = echo
              .addAsset(body[i].hash, body[i].block_number)
              .encodeABI();
            const klayRequest = caver.klay.accounts
              .signTransaction(
                {
                  type: 'FEE_DELEGATED_SMART_CONTRACT_EXECUTION',
                  from: senderAddress,
                  to: handler.getContractAddress(),
                  data,
                  gas: '300000',
                  nonce: getCurrentNonce(senderAddress)
                },
                senderPrivateKey
              )
              .then(
                result => {
                  const { rawTransaction: senderRawTransaction } = result;
                  // console.log('Logic feePayer');
                  // console.log(handler.getFeePayer());
                  // console.log('-----------------------');
                  return caver.klay
                    .sendTransaction({
                      senderRawTransaction: senderRawTransaction,
                      feePayer: handler.getFeePayer()
                    })
                    .then(
                      result => {
                        return result;
                      },
                      err => {
                        console.log(err);
                        return err;
                      }
                    );
                },
                error => {
                  console.log(error);
                  return error;
                }
              );
            promises.push(klayRequest);
          }

          Promise.all(promises)
            .then(r => {
              handler.setResponseBody(r).setStatusCode(200);
              callback(null, handler);
            })
            .catch(err => {
              // console.log("EEEEEEEEEEEEEEEEEEEEEE")
              handler.setErrorMessage(err);
              callback(handler);
            });
        });
      });
    });
  }

  static schema(body) {
    const hashSchema = Joi.object().keys({
      hash: Joi.string().required(),
      block_number: Joi.string().required(),
      account: Joi.string().empty(''),
      private_key: Joi.string().empty('')
    });
    const schema = Joi.array().items(hashSchema);
    // Return result.
    const result = Joi.validate(body, schema);
    if (result.error === null) {
      return;
    }
    return result;
  }
}
module.exports = Function;
