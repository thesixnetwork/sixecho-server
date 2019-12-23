const argv = require('minimist')(process.argv.slice(2));
const apiName = argv['a']
const storageAddress = '0x'+argv['r']
const newAPI = artifacts.require(apiName)
const fs = require('fs')

module.exports = async function(deployer, network) {
  if (argv['s'] == 'new_api') {
        await deployer.deploy(newAPI, storageAddress,{overwrite: true})
  } else {
    await console.log('Skipped.')
  }
}
