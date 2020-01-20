const argv = require('minimist')(process.argv.slice(2));
const apiName = argv['a']
const storageAddress = '0x'+argv['r']
const newAPI = artifacts.require(apiName)
const fs = require('fs')

module.exports = async function(deployer) {
  if (argv['s'] == 'new_api') {
        console.log("new_api")
        // console.log(newAPI)
        console.log(storageAddress)
        await deployer.deploy(newAPI, storageAddress)
  } else {
    await console.log('Skipped.')
  }
}
