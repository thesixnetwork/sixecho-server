const Caver = require('caver-js');
const AWS = require('aws-sdk');
const EchoAPI = require('../abi/contracts/APIv100.json');
const caver = new Caver(process.env.NETWORK_PROVIDER_URL);

let account;
let feePayer;

const echoAPI = new caver.klay.Contract(
  EchoAPI.abi,
  process.env.API_CONTRACT_ADDRESS
);
var ssm = new AWS.SSM({
  apiVersion: '2014-11-06'
});

class Handler {
  constructor() {
    this._caver = caver;
    this._echo_api = echoAPI.methods;
    this._echo = echoAPI;
    this._body = {};
    this._message = '';
    this._status = 200;
    this._is_error = false;
    this._error_message = 'error';
    if (account) {
      this._account = account;
      this._feepayer = feePayer;
      return Promise.resolve(this);
    } else {
      return setFeePayerAccount().then(acc => {
        this._account = acc[0];
        this._feepayer = acc[1];
        return Promise.resolve(this);
      });
    }
  }

  getContractAddress() {
    return this._echo.options.address;
  }

  getCaverAPI() {
    return this._caver;
  }

  getEchoAPI() {
    return this._echo_api;
  }

  getAccount() {
    return this._account;
  }

  getCallerAddress() {
    return this._account.address;
  }

  getAccountDefault() {
    return {
      address: this._account.address,
      privateKey: this._account.privateKey
    };
  }

  getResponseBody() {
    return this._body;
  }

  getErrorMessage() {
    return this._error_message;
  }

  setResponseBody(body) {
    this._body = body;
    return this;
  }

  setMessage(body) {
    this._message = body;
    return this;
  }

  setStatusCode(body) {
    this._status = body;
    if (body > 300) this._is_error = true;

    return this;
  }
  getFeePayer() {
    return this._feepayer;
  }
  setErrorMessage(body) {
    let msg = '';
    if (body instanceof Error) {
      msg = `${body.name}: ${body.message}`;
    } else if (typeof body === 'object') {
      msg = body.message ? body.message : body;
    }
    this._error_message = msg;
    this._is_error = true;
    if (this._status < 300) this._status = 400;
    return this;
  }

  addPrivateKey(privateKey, account) {
    caver.klay.accounts.wallet.add(privateKey, account);
  }

  // eslint-disable-next-line no-unused-vars
  static Response(h) {
    if (h instanceof Handler) {
      if (h._is_error === true) {
        return {
          status: h._status,
          message: h._error_message
        };
      }
      return {
        status: h._status || 200,
        message: h._message,
        body: h._body
      };
    } else if (h instanceof Error) {
      const errBody = { message: h.message };
      errBody.status = 400;
      if (process.env.NODE_ENV === 'test') errBody.stack = h.stack;
      return errBody;
    }
    return {
      status: 500,
      message: 'unknown error occur.'
    };
  }
}

function setSK2Account() {
  const options = { Name: 'SK_ECHO_WALLET', WithDecryption: true };
  return new Promise((resolve, reject) => {
    ssm.getParameter(options, (err, data) => {
      if (err) {
        reject(err);
        return;
      }
      caver.klay.accounts.wallet.clear();
      account = caver.klay.accounts.wallet.add(data.Parameter.Value);
      resolve(account);
    });
  });
}

function setFeePayerAccount() {
  const options = { Name: 'SK_ECHO_FEE_PAYER_ACCOUNT', WithDecryption: true };
  const options2 = { Name: 'SK_ECHO_FEE_PAYER_PRIVATE', WithDecryption: true };
  const options3 = { Name: 'SK_ECHO_WALLET', WithDecryption: true };

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
    caver.klay.accounts.wallet.clear();
    caver.klay.accounts.wallet.add(values[1], values[0]);
    console.log("FeePayer Key");
    console.log(values[1]+" "+ values[0]);
    console.log("----------------------");
    account = caver.klay.accounts.wallet.add(values[2]);
    feePayer = values[0]
    return [account, values[0]];
  });
}

module.exports = Handler;
