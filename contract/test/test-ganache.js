const Web3 = require('web3')
const web3 = new Web3('http://localhost:7545')
// enter your smart contract address

const argv = require('minimist')(process.argv.slice(2));
const network = argv['network']

const propReader = require('properties-reader')
const addressBooks = propReader('../address_book_'+network+'.properties')

const echoAppAddress = addressBooks.getRaw('addresses.EchoApp')
const callerAddr = web3.eth.accounts[0]
const EchoApp = require('../build/contracts/EchoApp.json')
const APIv100 = require('../build/contracts/APIv100.json')
//enter your smart contract address
// const apiv100 = new web3.eth.Contract(APIv100.abi, contractAddress)

const echoApp = new web3.eth.Contract(EchoApp.abi, echoAppAddress)


getAPIAddressByName('APIv100')
  .then(console.log)
  .catch(console.error)

// addBook('0000001','AnimalFarm 2', '978-3-16-148410-1')
//  .then(console.log)
//  .catch(console.error)

// addWriter('0xd87d72c191f3476640e943b1bc54067d96db1710')
//  .then(console.log)
//  .catch(console.error)

// getDataControllerAddress()
//   .then(console.log)
//   .catch(console.error)

// function addBook(key,name, isbn) {
//   return new Promise((resolve, reject) => {
//     apiv100.methods
//       .addBook(key,name, isbn)
//       .send({ from: callerAddr }, (err, transactionHash) => {
//         if (err) {
//           return reject(err)
//         }
//         return resolve(transactionHash)
//       })
//   })
// }

function getAPIAddressByName(apiname) {
  return echoApp.methods.getAPIAddressByName(apiname).call()
}

// function getDataControllerAddress() {
//   return new Promise((resolve, reject) => {
//     apiv100.methods
//       .getDataControllerAddress()
//       .send({ from: callerAddr }, (err, transactionHash) => {
//         if (err) {
//           return reject(err)
//         }
//         return resolve(transactionHash)
//       })
//   })
// }