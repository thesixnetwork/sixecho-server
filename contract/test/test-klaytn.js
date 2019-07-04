const Caver = require('caver-js')
const caver = new Caver('HTTP://127.0.0.1:7545')
// enter your smart contract address
const contractAddress = '0x69800b7Edb8776682f508FC612ae77e3914AB48D'
const EchoApp = require('../build/contracts/EchoApp.json')
// enter your smart contract address
const echoApp = new caver.klay.Contract(EchoApp.abi, contractAddress)
const events = echoApp.events
console.log(events)
events.allEvents(function(error, result) {
  // result will contain various information
  // including the argumets given to the `Deposit`
  // call.

  if (error) {
    console.log('error')
    console.log(error)
    console.log('asdf')
  }
  console.log(result)
})
echoApp.events
  .allEvents({}, (error, event) => {
    console.log(event)
  })
  .on('data', function(event) {
    console.log(event) // same results as the optional callback above
  })
  .on('error', console.error)
setTimeout(() => {
  echoApp.methods
    .initial()
    .call()
    .then(r => {
      console.log(r)
    })
}, 1000)
/*
echoApp.getPastEvents(
  'allEvents',
  {
    filter: {}, // Using an array means OR: e.g. 20 or 23
    fromBlock: 0,
    toBlock: 'latest'
  },
  function(error, events) {
    console.log(error)
    console.log(events)
  }
)
*/
