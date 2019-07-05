pragma solidity ^0.4.24;

import "../persistence/Storage.sol";

contract MetaDataController is Storage {

    /*
    constructor(address storageAddress) {
        _owner = msg.sender;
        _storage = Storage(storageAddress);
    }
    */

    function addNewBook(string _name, string _isbn) internal returns(uint256) {
        // Validate data
        uint256 newKey = getNewKey();
        setString(newKey, "name", _name);
        setString(newKey, "isbn", _isbn);
        return newKey;
    }

    function getBookByKey(uint256 _key) internal view returns(string,string) {
        return (
          getString(_key, "name"),
          getString(_key,"isbn")
        );
    }

}