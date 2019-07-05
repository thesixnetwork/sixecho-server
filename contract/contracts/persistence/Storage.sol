pragma solidity ^0.4.24;

import "../security/AccessRestriction.sol";

contract Storage is AccessRestriction {

    uint256 key = 0;

    mapping(uint256 => mapping(string => string)) internal stringAttributes; // use to private
    mapping(uint256 => mapping(string => uint256)) private uintAttributes;
    mapping(uint256 => mapping(string => address)) private addressAttributes;
    mapping(uint256 => mapping(string => bool)) private boolAttributes;

    /*
    constructor(address appAddress) internal {
        writers[msg.sender] = true;
        writers[appAddress] = true;
    }
    */

    function addWriter(address newWriter) public onlyOwner {
        writers[newWriter] = true;
    }


    function totalKey() public view returns (uint256) {
        return key;
    }

    function getNewKey() internal returns (uint256) {
        key++;
        return key;
    }

    function setString(uint256 _key,string _attriName,string _value) internal onlyWriter {
        stringAttributes[_key][_attriName] = _value;
    }

    function getString(uint256 _key,string _attriName) public view returns (string) {
        return stringAttributes[_key][_attriName];
    }

    function _toBytes(uint256 x) private pure returns (bytes32) {
        return bytes32(x);
    }
}