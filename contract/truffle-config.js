// truffle-config.js
module.exports = {
  networks: {
    develop: {
      host: '127.0.0.1',
      port: 7545,
      //from: '0x5a40651fa12c5c69152d216de4570ddb50321b30', // enter your account address
      network_id: '*' // Baobab network id
    }
    //  klaytn: {
    //    host: '127.0.0.1',
    //    port: 8551,
    //    from: '0x5a40651fa12c5c69152d216de4570ddb50321b30', // enter your account address
    //    network_id: '1001', // Baobab network id
    //    gas: 20000000, // transaction gas limit
    //    gasPrice: 25000000000 // gasPrice of Baobab is 25 Gpeb
    //  }
  },
  compilers: {
    solc: {
      version: '0.4.24' // Specify compiler's version to 0.4.24
    }
  }
}
