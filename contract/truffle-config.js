// truffle-config.js
const HDWalletProvider = require('truffle-hdwallet-provider-klaytn')
module.exports = {
  networks: {
    testnet: {
      // host: '127.0.0.1',
      // port: 8551,
      provider: () => new HDWalletProvider(
        '0xe8151654d9ba440883fd9c15bbd8e3c0d5eabb614ed0ffda2051dea92a95fb9b', 
        'https://api.baobab.klaytn.net:8651'),
      from: '0x9df799fed9eb39dfc1beb32bad4303d0990725f3', // enter your account address
      network_id: '1001', // Baobab network id
      gas: 20000000, // transaction gas limit
      gasPrice: 25000000000 // gasPrice of Baobab is 25 Gpeb
    },
    ganache: {
      host: '127.0.0.1',
      port: 7545,
      //from: '0x5a40651fa12c5c69152d216de4570ddb50321b30', // enter your account address
      network_id: '*' // Baobab network id
    },
    mainnet: {
      // host: 'api.cypress.klaytn.net',
      // port: 8651,
      provider: () =>
        new HDWalletProvider(
          '[[WALLET_SECRET]]',
          'https://api.cypress.klaytn.net:8651'
        ),
      from: '0xb989d084019a899d11e66853688649307f8d5070', // enter your account address
      network_id: '8217', // Baobab network id
      gas: 20000000, // transaction gas limit
      gasPrice: 25000000000 // gasPrice of Baobab is 25 Gpeb
    }
  },
  compilers: {
    solc: {
      version: '0.4.24' // Specify compiler's version to 0.4.24
    }
  }
}
