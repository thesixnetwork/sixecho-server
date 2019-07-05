const Caver = require('caver-js')
const caver = new Caver('HTTP://127.0.0.1:8551')
// enter your smart contract address
const contractAddress = '0x063cb35c2ba39b7dc8cb3ffc148190fdaaeb28a6'
const callerAddr = '0x9df799fed9eb39dfc1beb32bad4303d0990725f3'
const APIv100 = require('../build/contracts/APIv100.json')
const apiv100 = new caver.klay.Contract(APIv100.abi, contractAddress)
const events = apiv100.events

getBookById(1)
  .then(console.log)
  .catch(console.error)

// addBook('AnimalFarm 2', '978-3-16-148410-1')
//  .then(console.log)
//  .catch(console.error)

// addWriter('0xd87d72c191f3476640e943b1bc54067d96db1710')
//   .then(console.log)
//   .catch(console.error)

function addBook(name, isbn) {
  return new Promise((resolve, reject) => {
    apiv100.methods
      .addBook(name, isbn)
      .send({ from: callerAddr, gas: 200000 }, (err, transactionHash) => {
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

