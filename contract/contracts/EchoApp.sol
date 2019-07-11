pragma solidity ^0.4.24;

import "./persistence/Storage.sol";
import "./security/AccessRestriction.sol";
import "./service/API.sol";
import "./service/APIv100.sol";
import "./common/Event.sol";

contract EchoApp is AccessRestriction ,Event{

    Storage private _echoData;

    mapping(string => API) apis;

    API latest;

    constructor() public  {
        _owner = msg.sender;
    }

    function setUpStorage(address addressStorage) public onlyBy(_owner) {
        _echoData = Storage(addressStorage);
    }

    // function switchApp(address newAppAddress) public onlyBy(_owner) {
    //     // Change new owner of Storage
    //     // Remove All Writers from Storage
    // }

    function addNewAPI(string apiName,address apiAddress) public onlyBy(_owner) returns (string) {
        
        API api = API(apiAddress);
        // emit OutputString("API","API Load OK");

        address controllerAddress = api.getDataControllerAddress();
        // emit OutputAddress("Controller",controllerAddress);
        _echoData.addWriter(controllerAddress);
        // apis.push(api);
        apis[apiName] = api;
        latest = api;

        return "OK";
    }

    function getLatestAPIAddress() public view returns (address) {
        return address(latest);
    }

    function getAPIAddressByName(string apiName) public view returns (address) {
        return address(apis[apiName]);
    }

    // function initial() {
    //     _echoData = new Storage(address(this));
    //     emit OutputAddress("Storage address",address(_echoData));

    //     APIv100 apiv100 = new APIv100(address(_echoData));
    //     emit OutputAddress("APIv100 address",address(apiv100));
    // }
}