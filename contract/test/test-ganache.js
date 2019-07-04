const Web3 = require('web3')
const web3 = new Web3('http://localhost:7545')
// enter your smart contract address

const contractAddress = '0x4c810ef2f324ee35b19555e7ba97836ac1554fe1'
const APIv100 = require('../build/contracts/APIv100.json')
//enter your smart contract address
const apiv100 = new web3.eth.Contract(APIv100.abi, contractAddress)
apiv100.methods
  //  .setString(web3.utils.fromAscii('0'), 'name', 'bookA')
  // .getBook(4)
  // //.totalKey()
  // .call()
  // .then(r => {
  //   console.log(r)
  // })
  .addBook('book-xxx', 'isbn-xxxx')
  .send(
    { from: '0x7414Ee66C066D1aC43c0521d1a7495cdF84c1472' },
    (error, transactionHash) => {
      console.log(error)
      console.log(transactionHash)
    }
  )
// const contractAddress = '0xec329adbff34512ccbecf2a427e14587cb57bb61'
// const APIv100 = require('../build/contracts/APIv100.json')
// // enter your smart contract address
// const apiv100 = new web3.eth.Contract(APIv100.abi, contractAddress)
// apiv100.methods
//   .getBook('bookA')
//   .call()
//   .then(r => {
//     console.log(r)
//   })
//   .catch(e => {
//     console.log(e)
//   })
