const argv = require('minimist')(process.argv.slice(2));
const network = argv['network']
const apiName = argv['apiname']
const apiUpdate = argv['apiupdate']

const propReader = require('properties-reader')
const props = propReader('../address_book_'+network+'.properties')
props.append('../network.properties')
const echoAppAddress = props.getRaw('addresses.EchoApp')

const Web3 = require('web3')
const web3 = new Web3(props.getRaw('network.'+network+'.url'))

const EchoApp = require('../build/contracts/EchoApp.json')
const echoApp = new web3.eth.Contract(EchoApp.abi, echoAppAddress)

const callerAddr = props.getRaw("network."+network+".from")

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
        curApiAddress = await getAPIAddressByName(apiUpdate)
        console.log("API adddress : "+curApiAddress);

        const apiBuild = require('../build/contracts/'+apiName+'.json')
        const apiInst = new web3.eth.Contract(apiBuild.abi, curApiAddress)

        // await addBook(apiInst,"000002","Book2","123456789")
        //   .then(console.log)
        //   .catch(console.error)
        await getBookById(apiInst,"000002")
        .then(console.log)
        .catch(console.error)

        return
    }
    catch(e){
        console.log(e);
        return 
    }
   
}

start()


function addBook(inst,key,name, isbn) {
  return new Promise((resolve, reject) => {
    inst.methods
      .addBook(key,name, isbn)
      .send({ from: callerAddr }, (err, transactionHash) => {
        if (err) {
          return reject(err)
        }
        return resolve(transactionHash)
      })
  })
}

function getBookById(inst,id) {
  return inst.methods.getBook(id).call()
}