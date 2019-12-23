pragma solidity ^0.4.24;

contract AccessRestriction {

    address internal _owner;
    mapping(address => bool) internal writers;

    constructor() internal {
        _owner = msg.sender;
        writers[msg.sender] = true;
    }

    modifier onlyOwner() {
        require(isOwner(), "Sender is not authorized with owner policy.");
        _;
    }

    modifier onlyWriter()
    {
        require(
            writers[msg.sender] == true || isOwner(),
            "Sender not authorized to write storage."
        );
        _;
    }

    modifier onlyBy(address _account)
    {
        require(
            msg.sender == _account,
            "Sender not authorized."
        );
        // Do not forget the "_;"! It will
        // be replaced by the actual function
        // body when the modifier is used.
        _;
    }

    modifier oneOfTwo(address _a1, address _a2)
    {
        require(
            msg.sender == _a1 || msg.sender == _a2,
            "Sender not authorized."
        );
        // Do not forget the "_;"! It will
        // be replaced by the actual function
        // body when the modifier is used.
        _;
    }

    /**
    * @return the address of the owner.
    */
    function owner() public view returns(address) {
        return _owner;
    }

    function isOwner() public view returns(bool) {
        return msg.sender == _owner;
    }

}