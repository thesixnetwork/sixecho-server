const Caver = require('caver-js')
// const caver = new Caver('HTTP://127.0.0.1:8551')

const caver = new Caver('https://api.cypress.klaytn.net:8651')
caver.klay.accounts.wallet.add('28df7859de12a1b9700deba0ee173312ad1dd9fadeffb2074b9722859d100211')

// enter your smart contract address
const contractAddress = '0x83b57c43c3e0f17639627e2f3822b187bf9b314f'
const callerAddr = '0xb989d084019a899d11e66853688649307f8d5070'
const APIv101 = require('../build/contracts/APIv101.json')
const apiv101 = new caver.klay.Contract(APIv101.abi, contractAddress)

const EchoApp = require('../build/contracts/EchoApp.json')
const echoApp = new caver.klay.Contract(EchoApp.abi, '0xa8534940836d81025a1a899c3de22affa548be74')
// const events = apiv101.events

// addNewAPI('APIv101','0x83b57c43c3e0f17639627e2f3822b187bf9b314f')
  // .then(console.log)
  // .catch(console.error)

// getDataControllerAddress()
//  .then(console.log)
//  .catch(console.error)

// getLatestAPIAddress()
//  .then(console.log)
//  .catch(console.error)

// addBook('11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111', '11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111')
//   .then(console.log)
//   .catch(console.error)

// addWriter('0xb989d084019a899d11e66853688649307f8d5070')
//   .then(console.log)
//   .catch(console.error)

// addAsset('1234','5678')
//   .then(console.log)
//   .catch(console.error)

// getLatestAPIAddress()
//      .then(console.log)
//      .catch(console.error)

function addBook(name, isbn) {
  return new Promise((resolve, reject) => {
    apiv101.methods
      .addBook(name, isbn)
      .send({ from: callerAddr, gas: 2000000 }, (err, transactionHash) => {
        if (err) {
          return reject(err)
        }
        return resolve(transactionHash)
      })
  })
}

function getLatestAPIAddress(){
  return echoApp.methods.getLatestAPIAddress().call()
}

function getDataControllerAddress() {
  return apiv101.methods.getDataControllerAddress().call()
}

function getBookById(id) {
  return apiv101.methods.getBook(id).call()
}

function addWriter(addr) {
  return new Promise((resolve, reject) => {
    apiv101.methods
      .addWriter(addr)
      .send({ from: callerAddr, gas: 2000000 }, (err, transactionHash) => {
        if (err) {
          return reject(err)
        }
        return resolve(transactionHash)
      })
  })
}

function addAsset(h,b) {
  return new Promise((resolve, reject) => {
    apiv101.methods
      .addAsset(h,b)
      .send({ from: callerAddr, gas: 2000000 }, (err, transactionHash) => {
        if (err) {
          return reject(err)
        }
        return resolve(transactionHash)
      })
  })
}

function addNewAPI( apiName, apiAddress) {
  return new Promise((resolve, reject) => {
    echoApp.methods
      .addNewAPI(apiName,apiAddress)
      .send({ from: callerAddr, gas: 20000000 }, (err, transactionHash) => {
        if (err) {
          return reject(err)
        }
        return resolve(transactionHash)
      })
  })
}

