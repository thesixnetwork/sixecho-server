const Joi = require('joi');
const Promise = require('bluebird');
const Handler = require('../middleware/handler.middleware');

class Function {
  constructor(body, callback) {
    console.log('-------------------------------------');
    new Handler().then(handler => {
      const echo = handler.getEchoAPI();
      const caver = handler.getCaverAPI();
      caver.klay.getTransactionCount(handler.getCallerAddress()).then(nonce => {
        const promises = [];
        for (let i = 0; i < body.length; i++) {
          let {
            privateKey: senderPrivateKey,
            address: senderAddress
          } = handler.getAccountDefault();
          if (body[i].private_key != '') {
            senderPrivateKey = body[i].private_key;
            senderAddress = body[i].account;
          }
          // console.log("-----------------------") ;
          // console.log(senderAddress);
          // console.log(senderPrivateKey);
          // console.log(handler.getFeePayer())
          // console.log("-----------------------") ;
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
                gas: 10000000,
                nonce: nonce + i
              },
              senderPrivateKey
            )
            .then(
              result => {
                const { rawTransaction: senderRawTransaction } = result;
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
                      return err;
                    }
                  );
              },
              error => {
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
            console.error(err);
            handler.setErrorMessage(err);
            callback(handler);
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
