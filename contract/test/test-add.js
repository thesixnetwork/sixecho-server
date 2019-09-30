const argv = require('minimist')(process.argv.slice(2))
const network = argv['network']
const apiName = argv['apiname']
const apiUpdate = argv['apiupdate']

const propReader = require('properties-reader')
const props = propReader('../address_book_' + network + '.properties')
props.append('../network.properties')
const echoAppAddress = props.getRaw('addresses.EchoApp')

const Web3 = require('web3')
const web3 = new Web3(props.getRaw('network.' + network + '.url'))

const EchoApp = require('../build/contracts/EchoApp.json')
const echoApp = new web3.eth.Contract(EchoApp.abi, echoAppAddress)

const callerAddr = props.getRaw('network.' + network + '.from')

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
    console.log('API adddress : ' + curApiAddress)

    const apiBuild = require('../build/contracts/' + apiName + '.json')
    const apiInst = new web3.eth.Contract(apiBuild.abi, curApiAddress)

    // var randomDigest = [];
    // for ( i=0;i<=128; i++){
    //   randomDigest.push(Math.floor(Math.random()*1000));
    // }

    // addBook(
    //   apiInst,
    //   '000004',
    //   'Book4',
    //   'Bank4',
    //   'THA',
    //   'th',
    //   100,
    //   '1',
    //   1563362788
    // )
    //   .then(console.log)
    //   .catch(console.error)

    // console.log(randomDigest);

    randomDigest = [
      13442035,
      150261961,
      67132350,
      103561595,
      182914072,
      75546008,
      68829758,
      25996082,
      154257997,
      64861726,
      81323651,
      9840801,
      34600070,
      78014731,
      5757667,
      8294693,
      59544370,
      145847892,
      85377301,
      46044126,
      20469469,
      73805333,
      81250880,
      9420512,
      93354697,
      17709743,
      61662889,
      18058311,
      173414506,
      32591289,
      141745270,
      41376982,
      128807942,
      212656921,
      23432208,
      110810588,
      31130091,
      6758885,
      18647083,
      20794948,
      123596416,
      22998340,
      84401280,
      6071202,
      34071252,
      74893762,
      21344463,
      45529261,
      120201373,
      23045511,
      7440208,
      173133351,
      21378680,
      23181250,
      30739931,
      128991460,
      188636076,
      145839650,
      46008764,
      223496323,
      168470566,
      57839075,
      85069589,
      20415169,
      63752447,
      14599548,
      33266903,
      93779578,
      63849683,
      44607934,
      9565948,
      10173181,
      9125874,
      71175489,
      149419894,
      95751201,
      36800030,
      108267293,
      17006543,
      20420349,
      134666605,
      43480155,
      86779709,
      69243351,
      30415323,
      33855823,
      5296051,
      52691529,
      73805374,
      8229621,
      24574405,
      76628343,
      1911171,
      36661455,
      41530573,
      117588054,
      16400560,
      13187931,
      18489598,
      112544168,
      7084345,
      87476238,
      63020317,
      149916,
      56911580,
      37286141,
      32557847,
      38236552,
      20458540,
      1867804,
      9740879,
      224096484,
      67101539,
      228620272,
      29303961,
      58886521,
      45282749,
      4877049,
      32031921,
      1376386,
      87099640,
      44134447,
      134594875,
      56997857,
      79634914,
      15436036,
      129915177,
      56623235
    ]

    // await uploadDigest(apiInst, '123456789', randomDigest)
    //   .then(console.log)
    //   .catch(console.error)

    // await getBookById(apiInst,"1234")
    // .then(console.log)
    // // .catch(console.error)

    // await getAdditionalBookData(apiInst,"1234")
    // .then(console.log)
    // .catch(console.error)

    await addAsset(
      apiInst,
      '0xC40548e6bE91162f16791740713514eB5a417DF7',
      '1234'
    )
      .then(console.log)
      .catch(console.error)

    // await downloadDigest(apiInst, '123456789')
    //   .then(console.log)
    //   .catch(console.error)

    return
  } catch (e) {
    console.log(e)
    return
  }
}

start()

function addBook(
  inst,
  key,
  title,
  author,
  origin,
  lang,
  paperback,
  publisherId,
  publishDate
) {
  return new Promise((resolve, reject) => {
    inst.methods
      .addBook(
        key,
        title,
        author,
        origin,
        lang,
        paperback,
        publisherId,
        publishDate
      )
      .send(
        { from: callerAddr, gasLimit: '10000000' },
        (err, transactionHash) => {
          if (err) {
            return reject(err)
          }
          return resolve(transactionHash)
        }
      )
  })
}

function uploadDigest(inst, key, digest) {
  return new Promise((resolve, reject) => {
    inst.methods
      .uploadDigest(key, digest)
      .send(
        { from: callerAddr, gasLimit: '10000000' },
        (err, transactionHash) => {
          if (err) {
            return reject(err)
          }
          return resolve(transactionHash)
        }
      )
  })
}

function addAsset(inst, h, blockNo) {
  return new Promise((resolve, reject) => {
    inst.methods
      .addAsset(h, blockNo)
      .send(
        { from: callerAddr, gasLimit: '10000000' },
        (err, transactionHash) => {
          if (err) {
            return reject(err)
          }
          return resolve(transactionHash)
        }
      )
  })
}

function getBookById(inst, id) {
  return inst.methods.getBook(id).call()
}

function getAdditionalBookData(inst, id) {
  return inst.methods.getAdditionalBookData(id).call()
}

function downloadDigest(inst, id) {
  return inst.methods.downloadDigest(id).call()
}
