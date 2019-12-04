const EchoAPI = require("../abi/contracts/APIv100.json");
const Caver = require("caver-js");
const caver = new Caver("https://api.cypress.klaytn.net:8651");

const echoAPI = new caver.klay.Contract(
  EchoAPI.abi,
  "0xad67c0115b1dbb8ba9a263ef49c2a8b14ccf8138"
);
const account = caver.klay.accounts.wallet.add(
  "28df7859de12a1b9700deba0ee173312ad1dd9fadeffb2074b9722859d100211"
);

const data = echoAPI.methods
  .addAsset(
    "9fa223bc80a81da667dfa2bf6c44f000bbdd62e7c9dcb23d1e4aba6e2e26a295",
    "1"
  )
  .encodeABI();
console.log(echoAPI.options.address);
console.log(data);
/*
caver.klay
  .getTransactionCount("0xb989d084019a899d11e66853688649307f8d5070")
  .then(nonce => {
    console.log(nonce);
    caver.klay
      .sendTransaction({
        type: "SMART_CONTRACT_EXECUTION",
        from: account.address,
        to: "0xad67c0115b1dbb8ba9a263ef49c2a8b14ccf8138",
        data: data,
        gas: 10000000,
        nonce: nonce
      })
      .then(r => {
        console.log(r);
      })
      .catch(e => {
        console.log(e);
      });
  });
*/
