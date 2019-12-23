const argv = require('minimist')(process.argv.slice(2));

const EchoApp = artifacts.require("EchoApp");
const Storage = artifacts.require("Storage");


module.exports = async function(deployer,network) {
  
  // console.log(argv)
  
  if (argv['s'] == 'update_api') {
    const updateAPI = artifacts.require(argv['name']);
    deployer.deploy(updateAPI,Storage.address,{ overwrite:true }).then(async () => {
          var appInstance = await EchoApp.deployed();
          console.log("udpate api to " + argv['name']+'u'+argv['update']);
          var result = await appInstance.addNewAPI(argv['name']+'u'+argv['update'],updateAPI.address);
          console.log(result);
          
        });
  }else {
    console.log("Skipped.")
  }
};
