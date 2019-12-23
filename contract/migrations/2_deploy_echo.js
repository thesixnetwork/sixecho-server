const EchoApp = artifacts.require('EchoApp')
const Storage = artifacts.require('Storage')
const APIv100 = artifacts.require('APIv100')
const fs = require('fs')
const argv = require('minimist')(process.argv.slice(2));

module.exports = async function(deployer, network) {
  // await deployer.deploy(EchoApp,{ overwrite:true });
  // await deployer.deploy(Storage, EchoApp.address, {overwrite:true})
  // await deployer.deploy(APIv100,Storage.address,{ overwrite:true });

  // var appInstance = await EchoApp.deployed();
  // await appInstance.setUpStorage(Storage.address);
  // console.log(argv)

  if (argv['s'] == 'init') {
    deployer.deploy(EchoApp).then(function() {
      return deployer.deploy(Storage, EchoApp.address).then(() => {
        return deployer.deploy(APIv100, Storage.address).then(async () => {
          var appInstance = await EchoApp.deployed()
          await appInstance.setUpStorage(Storage.address)
          await appInstance.addNewAPI('APIv100', APIv100.address)
          addressBookContent =
            'addresses.Storage=' +
            Storage.address +
            '\n' +
            'addresses.EchoApp=' +
            EchoApp.address
          fs.writeFile(
            'address_book_' + network + '.properties',
            addressBookContent,
            err => {
              if (err) console.log(err)
              console.log('Successfully Written to File.')
            }
          )
        })
      })
    })
  } else {
    await console.log('Skipped.')
  }
}
