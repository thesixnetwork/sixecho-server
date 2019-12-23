const Caver = require('caver-js')
// const caver = new Caver('HTTP://127.0.0.1:8551')

const caver = new Caver('https://api.cypress.klaytn.net:8651')
caver.klay.accounts.wallet.add('0x28df7859de12a1b9700deba0ee173312ad1dd9fadeffb2074b9722859d100211')

// enter your smart contract address
const contractAddress = '0x2bb5996678451804b2872a5dfa8b91932e37a3d4'
const callerAddr = '0xb989d084019a899d11e66853688649307f8d5070'
const APIv101 = require('../build/contracts/APIv101.json')
const apiv101 = new caver.klay.Contract(APIv101.abi, contractAddress)

const EchoApp = require('../build/contracts/EchoApp.json')
const echoApp = new caver.klay.Contract(EchoApp.abi, '0xa8534940836d81025a1a899c3de22affa548be74')
// const events = apiv101.events

// addNewAPI('APIv101','0xfe7ef9d8073e3b7aa685d27fd513244d74e55f62')
//   .then(console.log)
//   .catch(console.error)

// getDataControllerAddress()
//  .then(console.log)
//  .catch(console.error)

// getLatestAPIAddress()
//  .then(console.log)
//  .catch(console.error)

// addBook('11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111', '11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111')
//   .then(console.log)
//   .catch(console.error)

// addWriter('0x9df799fed9eb39dfc1beb32bad4303d0990725f3')
//   .then(console.log)
//   .catch(console.error)

addAsset('1234','5678')
  .then(console.log)
  .catch(console.error)

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

