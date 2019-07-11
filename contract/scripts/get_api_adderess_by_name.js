const argv = require('minimist')(process.argv.slice(2));
const network = argv['network']
const apiName = argv['name']

const propReader = require('properties-reader')
const props = propReader('../address_book_'+network+'.properties')
props.append('../network.properties')
const echoAppAddress = props.getRaw('addresses.EchoApp')

const Web3 = require('web3')
const web3 = new Web3(props.getRaw('network.'+network+'.url'))

const EchoApp = require('../build/contracts/EchoApp.json')
const echoApp = new web3.eth.Contract(EchoApp.abi, echoAppAddress)

// const apiBuild = require('../build/contracts/'+apiName+'.json')

function getAPIAddressByName(apiname) {
    return echoApp.methods.getAPIAddressByName(apiname).call()
}

// async function getAPIAddressByName(apiname) {

//     const address = await new Promise((resolve, reject) => {
//         echoApp.methods.getAPIAddressByName(apiname).call({from: props.getRaw('network.'+network+'.from')}, (error, result) => {
//           if (!error) reject(error)
//           resolve(result)
//         });
        
//     })

//     return address
// }
  

async function start() {
    try {
        curApiAddress = await getAPIAddressByName(apiName)
        console.log(curApiAddress);
        return
    }
    catch(e){
        console.log('error');
        return 
    }
   
}

start()