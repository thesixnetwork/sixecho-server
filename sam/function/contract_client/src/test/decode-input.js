const EchoAPI = require('../abi/contracts/APIv100.json');
const abiDecoder = require('abi-decoder');
const Caver = require('caver-js');
const caver = new Caver('https://api.cypress.klaytn.net:8651');
abiDecoder.addABI(EchoAPI.abi);
const data =
  '0xda2824a80000000000000000000000005260266f23b6b35df1151e6019316c6b7ad6c63e';
const decodedData = abiDecoder.decodeMethod(data);
console.log(decodedData);
