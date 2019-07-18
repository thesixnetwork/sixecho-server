pragma solidity ^0.4.24;

import "../security/AccessRestriction.sol";

contract Storage is AccessRestriction {

    address _appAddress;

    mapping(address => bool) internal writers;

    mapping(string => mapping(string => string)) private stringAttributes;
    mapping(string => mapping(string => uint256)) private uintAttributes;
    mapping(string => mapping(string => address)) private addressAttributes;
    mapping(string => mapping(string => bool)) private boolAttributes;


    mapping(string => mapping(string => uint256[])) private uintArrayAttributes;

    
    constructor(address appAddress) public  {
        _owner = msg.sender;
        _appAddress = appAddress;
        writers[_owner] = true;
        writers[appAddress] = true;
    } 

    function addWriter(address newWriter) public oneOfTwo(_owner,_appAddress) {
        writers[newWriter] = true;
    }

    modifier onlyWriter()
    {
        require(
            writers[msg.sender] == true,
            "Sender not authorized to write storage."
        );
        // Do not forget the "_;"! It will
        // be replaced by the actual function
        // body when the modifier is used.
        _;
    }

    // function getNewKey() public returns (string) {
    //     _key = _key + 1;
    //     string newKey = string(_key);
    //     return newKey;
    // }

    function setString(string key,string attriName,string value) public onlyWriter {
        stringAttributes[key][attriName] = value;
    }

    function getString(string key,string attriName) public view returns (string) {
        return stringAttributes[key][attriName];
    }

    function setUint(string key,string attriName,uint256 value) public onlyWriter {
        uintAttributes[key][attriName] = value;
    }

    function getUint(string key,string attriName) public view returns (uint256) {
        return uintAttributes[key][attriName];
    }

    function setBool(string key,string attriName,bool value) public onlyWriter {
        boolAttributes[key][attriName] = value;
    }

    function getBool(string key,string attriName) public view returns (bool) {
        return boolAttributes[key][attriName];
    }

    function setAddress(string key,string attriName,address value) public onlyWriter {
        addressAttributes[key][attriName] = value;
    }

    function getAddress(string key,string attriName) public view returns (address) {
        return addressAttributes[key][attriName];
    }

    function setUintArray(string key,string attriName,uint256[] value) public onlyWriter {
        uintArrayAttributes[key][attriName] = value;
    }

    function getUintArray(string key,string attriName) public view returns (uint256[]) {
        return uintArrayAttributes[key][attriName];
    }
}