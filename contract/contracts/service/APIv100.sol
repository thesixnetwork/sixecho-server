pragma solidity ^0.4.24;

import "./API.sol";
import "../controller/MetaDataController.sol";
import "../security/AccessRestriction.sol";
import "../common/Event.sol";

contract APIv100 is API ,AccessRestriction,Event {

    string constant _version = "v1.0";

    address _storageAddress;

    MetaDataController _metaDataController;

    constructor (address storageAddress) public {
        _owner = msg.sender;
        _metaDataController = new MetaDataController(storageAddress);
        _storageAddress = storageAddress;
    }

    function getDataControllerAddress() public view returns (address) {
        return address(_metaDataController);
    }

    function getVersion() public view returns (string) {
        return _version;
    }

    function addBook(string newKey,string name,string isbn) public onlyBy(_owner) returns(string key) {
        key = _metaDataController.addBook(newKey,name,isbn);
        // emit OutputBytes32("returnKey",key);
        return key;
    }

    function getBook(string key) public view returns (string,string) {
        
        string memory name;
        string memory isbn;

        (name,isbn) = _metaDataController.getBookByKey(key);

         return (name,isbn);
    }
}