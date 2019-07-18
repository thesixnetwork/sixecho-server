pragma solidity ^0.4.24;

contract MetaData {
    struct MetaDataModel {
        string title;
        string author;
        string lang;
        string origin;
        uint256 paperback;
        uint256 publihDate;
        uint256 catId;
        uint256 publisherId;
    }
}