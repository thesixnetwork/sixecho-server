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

    function addAsset(string h, string blockNo) public onlyOwner returns (string) {
        _storage.setString(h,"hash",h);
        _storage.setString(blockNo,"block_number",blockNo);
        return h; 
    }

    function addWriter(address addr) public onlyOwner returns (address) {
        _storage.addWriter(addr);
        return addr;
    }

    function addBook(string newKey,string title,string author,string origin,string lang,uint256 paperback,string publisherId,uint256 publishDate) public onlyOwner returns(string returnKey) {
        // Validate data
        // newKey = _storage.getNewKey();
        returnKey = newKey;
        // emit OutputString("title",title);
        // emit OutputString("author",author);
        _storage.setString(newKey,"title",title);
        _storage.setString(newKey,"author",author);
        _storage.setString(newKey,"origin",origin);
        _storage.setString(newKey,"lang",lang);
        _storage.setUint(newKey,"paperback",paperback);
        _storage.setString(newKey,"publisher_id",publisherId);
        _storage.setUint(newKey,"publish_date",publishDate);
        return returnKey;
    }

    function uploadDigest(string newKey,uint256[] digest) public onlyOwner returns(string returnKey) {
        _storage.setUintArray(newKey,"digest",digest);

        returnKey = newKey;
    }

    function getBookByKey(string key) public view returns(string title,string author,string lang,string publisherId,uint256 publishDate) {
         title = _storage.getString(key,"title");
         author = _storage.getString(key,"author");
         lang = _storage.getString(key,"lang");
         publisherId = _storage.getString(key,"publisher_id");
         publishDate = _storage.getUint(key,"publish_date");
    }

    function downloadDigest(string newKey) public view returns(uint256[] digest) {
        digest = _storage.getUintArray(newKey,"digest");
    }

    function getAdditionalBookData(string key) public view returns(string origin,uint256 paperback) {
        origin = _storage.getString(key,"origin");
        paperback = _storage.getUint(key,"paperback");
    }

}