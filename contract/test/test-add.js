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

        var randomDigest = [];
        for ( i=0;i<=128; i++){
          randomDigest.push(Math.floor(Math.random()*1000));
        }

        // await addBook(apiInst,"000004","Book4","Bank4","THA","th",100,"1",1563362788)
        //   .then(console.log)
        //   .catch(console.error);

        // console.log(randomDigest);
        
        // await uploadDigest(apiInst,"000003",randomDigest)
        //   .then(console.log)
        //   .catch(console.error);

        await getBookById(apiInst,"000004")
        .then(console.log)
        .catch(console.error)

        // // await downloadDigest(apiInst,"000001")
        // .then(console.log)
        // .catch(console.error)

        return
    }
    catch(e){
        console.log(e);
        return 
    }
   
}

start()


function addBook(inst,key,title,author,origin,lang,paperback,publisherId,publishDate) {
  return new Promise((resolve, reject) => {
    inst.methods
      .addBook(key,
        title,
        author,
        origin,
        lang,
        paperback,
        publisherId,publishDate)
      .send({ from: callerAddr ,gasLimit: "10000000"}, (err, transactionHash) => {
        if (err) {
          return reject(err)
        }
        return resolve(transactionHash)
      })
  })
}

function uploadDigest(inst,key,digest) {
  return new Promise((resolve, reject) => {
    inst.methods
      .uploadDigest(key,
        digest)
      .send({ from: callerAddr ,gasLimit: "10000000"}, (err, transactionHash) => {
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

function downloadDigest(inst,id) {
  return inst.methods.downloadDigest(id).call()
}