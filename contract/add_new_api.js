const argv = require('minimist')(process.argv.slice(2));
console.log(argv)
const network_url = argv['n']
const wallet = '0x'+argv['w']
const apiAddr = '0x'+argv['a']
const apiName = argv['p']
const callerAddr = '0x'+argv['c']
const echoAppAdr = '0x'+argv['e']

const Caver = require('caver-js')
const caver = new Caver(network_url)
caver.klay.accounts.wallet.add(wallet)

// enter your smart contract address
const API = require('./build/contracts/'+apiName+'.json')
const api = new caver.klay.Contract(API.abi, apiAddr)

const EchoApp = require('./build/contracts/EchoApp.json')
const echoApp = new caver.klay.Contract(EchoApp.abi, echoAppAdr)

async function start() {
  try{

    await addNewAPI(apiName,apiAddr)
    .then(console.log)
    .catch(console.error)

    await getLatestAPIAddress()
     .then(console.log)
     .catch(console.error)

    await addWriter(callerAddr)
    .then(console.log)
    .catch(console.error)

  } catch (e) {
    console.log(e)
    return
  }
}

start()

function addWriter(addr) {
  return new Promise((resolve, reject) => {
    api.methods
      .addWriter(addr)
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

//  getLatestAPIAddress()
//  .then(console.log)
//  .catch(console.error)

function getLatestAPIAddress(){
  return echoApp.methods.getLatestAPIAddress().call()
}