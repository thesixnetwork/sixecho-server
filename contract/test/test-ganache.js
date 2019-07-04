const Web3 = require('web3')
const web3 = new Web3('http://localhost:7545')
// enter your smart contract address

const contractAddress = '0x4c810ef2f324ee35b19555e7ba97836ac1554fe1'
const callerAddr = '0x9df799fed9eb39dfc1beb32bad4303d0990725f3'
const APIv100 = require('../build/contracts/APIv100.json')
//enter your smart contract address
const apiv100 = new web3.eth.Contract(APIv100.abi, contractAddress)

getBookById(1)
  .then(console.log)
  .catch(console.error)

// addBook('AnimalFarm 2', '978-3-16-148410-1')
//  .then(console.log)
//  .catch(console.error)

// addWriter('0xd87d72c191f3476640e943b1bc54067d96db1710')
//  .then(console.log)
//  .catch(console.error)

function addBook(name, isbn) {
  return new Promise((resolve, reject) => {
    apiv100.methods
      .addBook(name, isbn)
      .send({ from: callerAddr }, (err, transactionHash) => {
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
      .send({ from: callerAddr }, (err, transactionHash) => {
        if (err) {
          return reject(err)
        }
        return resolve(transactionHash)
      })
  })
}
