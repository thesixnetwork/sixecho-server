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

    function addAsset(string h, string blockNo) returns (string) {

        string memory returnKey;
        returnKey = _metaDataController.addAsset(h, blockNo);

        return returnKey;

    }

    function addBook(string newKey,
                string title,
                string author,
                string origin,
                string lang,
                uint256 paperback,
                string publisherId,
                uint256 publishDate
                ) public onlyBy(_owner) returns(string) {
        string memory returnKey;
        returnKey = _metaDataController.addBook(newKey,title,author,origin,lang,paperback,publisherId,publishDate);
        // emit OutputBytes32("returnKey",key);

        return returnKey;
    }

    function uploadDigest(string newKey,uint256[] digest) public onlyBy(_owner) returns(string) {
        string memory returnKey;
        returnKey = _metaDataController.uploadDigest(newKey,digest);
        return returnKey;
    }

    function downloadDigest(string newKey) public view returns (uint256[] digest) {
        digest = _metaDataController.downloadDigest(newKey);
    }

    function getBook(string key) public view returns (string title,string author,string lang,string publisherId,uint256 publishDate) {

        (title,author,lang,publisherId,publishDate) = _metaDataController.getBookByKey(key);

        //  return (title,author,lang,publisherId,publishDate);
    }

    function getAdditionalBookData(string key) public view returns(string origin,uint256 paperback) {
        (origin,paperback) = _metaDataController.getAdditionalBookData(key);
    }
}