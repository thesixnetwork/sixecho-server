const Migrations = artifacts.require('Migrations')
const APIv101 = artifacts.require('APIv101')
const EchoApp = artifacts.require('EchoApp')

// console.log(EchoApp)

module.exports = function(deployer) {
  // deployer.deploy(Migrations)
  // deployer.deploy(APIv101, '0xe3448491e64604d6e9032794d70f6325921f8247',{overwrite: true})

  // deployer.then(function() {
  //   // appInstance = EchoApp.new();
  //   echoApp = EchoApp.at('0x19bbbf5b6d1c7f5403411525aaf2f71b2c6ca68f');
  //   return echoApp;
  // }).then(function(echoApp){
  //   echoApp.setUpStorage('0xe3448491e64604d6e9032794d70f6325921f8247');
  //   return echoApp
  // }).then(function(echoApp){
  //   echoApp.addNewAPI('APIv101', '0x24797fe286403176030d091d363c1a8e414c8ac0')
  // })
}
