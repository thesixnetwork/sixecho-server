const Caver = require('caver-js')
const caver = new Caver('HTTP://127.0.0.1:8551')
// enter your smart contract address
const contractAddress = '0x9e8d100e7579cc14d9ff70cbe82863bfe7b24c70'
const callerAddr = '0x402fc4eecfefbe25325c7fc98a423bcbcb7ba4a4'
const APIv100 = require('../build/contracts/APIv100.json')
const apiv100 = new caver.klay.Contract(APIv100.abi, contractAddress)
const events = apiv100.events

//getBookById(1)
//  .then(console.log)
//  .catch(console.error)

addBook('11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111', '11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111')
  .then(console.log)
  .catch(console.error)

// addWriter('0xd87d72c191f3476640e943b1bc54067d96db1710')
//   .then(console.log)
//   .catch(console.error)

function addBook(name, isbn) {
  return new Promise((resolve, reject) => {
    apiv100.methods
      .addBook(name, isbn)
      .send({ from: callerAddr, gas: 2000000 }, (err, transactionHash) => {
        if (err) {
          return reject(err)
        }
        return resolve(transactionHash)
      })
  })
}

function getBookById(id) {
  return apiv100.methods.getBook(id).call()
}

function addWriter(addr) {
  return new Promise((resolve, reject) => {
    apiv100.methods
      .addWriter(addr)
      .send({ from: callerAddr, gas: 200000 }, (err, transactionHash) => {
        if (err) {
          return reject(err)
        }
        return resolve(transactionHash)
      })
  })
}

