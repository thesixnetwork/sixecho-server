const Joi = require('joi');
const Promise = require('bluebird');
const Handler = require('../middleware/handler.middleware');

class Function {
  constructor(body, callback) {
    new Handler().then(handler => {
      const echo = handler.getEchoAPI();
      const caver = handler.getCaverAPI();
      caver.klay.getTransactionCount(handler.getCallerAddress()).then(nonce => {
        const promises = [];
        for (let i = 0; i < body.length; i++) {
          const data = echo
            .addAsset(body[i].hash, body[i].block_number)
            .encodeABI();
          const klayRequest = caver.klay
            .sendTransaction({
              type: 'SMART_CONTRACT_EXECUTION',
              from: handler.getCallerAddress(),
              to: handler.getContractAddress(),
              data,
              gas: 10000000,
              nonce: nonce + i
            })
            .on('error', e => {
              throw e;
            });
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
      account: Joi.string()
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
