pragma solidity ^0.4.24;

import "../persistence/Storage.sol";
import "../security/AccessRestriction.sol";
import "../common/Event.sol";

contract MetaDataController is AccessRestriction,Event{

    Storage _storage;

    constructor (address storageAddress) public{
        _owner = msg.sender;
        _storage = Storage(storageAddress);
    }

    function addBook(string newKey,string name,string isbn) public returns(string returnKey) {
        // Validate data
        // newKey = _storage.getNewKey();
        returnKey = newKey;
        emit OutputString("name",name);
        emit OutputString("isbn",isbn);
        _storage.setString(newKey,"name",name);
        _storage.setString(newKey,"isbn",isbn);
        return returnKey;
    }

    function getBookByKey(string key) public view returns(string name,string isbn) {
         name = _storage.getString(key,"name");
         isbn = _storage.getString(key,"isbn");
    }

}