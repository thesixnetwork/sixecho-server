pragma solidity ^0.4.24;

import "../controller/MetaDataController.sol";
import "../common/Event.sol";

contract APIv100 is MetaDataController,Event {

    string constant _version = "v1.0";

    function getVersion() public pure returns (string) {
        return _version;
    }

    function addBook(string _name,string _isbn) public onlyWriter returns(uint256) {
        uint256 key = addNewBook(_name, _isbn);
        emit OutputUint256("returnUint256",key);
        return key;
    }

    function getBook(uint256 _key) public view returns (string,string) {
        return getBookByKey(_key);
    }
}