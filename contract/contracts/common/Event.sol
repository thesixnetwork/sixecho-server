pragma solidity ^0.4.24;

contract Event {
    event OutputString(string key, string value);
    event OutputAddress(string key, address value);
    event OutputBytes32(string key,bytes32 value);
    event OutputUint256(string key,uint256 value);
}