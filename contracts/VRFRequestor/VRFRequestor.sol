// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract RandomRequestManager {
    event RequestReceived(uint256 indexed requestId, string seed, uint256 numOfValues);

    // TO DO: Include type of request once we include all 3 randamu api endpoints.
    enum RequestType {
        Bytes,
        IntRange,
        IntArray
    }

    uint256 public counter;
    mapping(uint256 => Request) public requests;

    struct Request {
        string seed;
        uint256 numOfValues;
    }

    function generateRequestId(address requester) internal view returns (uint256) {
        return uint256(keccak256(abi.encode(requester, counter)));
    }

    function request(uint256 numOfValues) external returns (uint256 requestId) {
        requestId = uint256(keccak256(abi.encode(msg.sender, counter++)));
        requests[requestId] = Request("", numOfValues);
        emit RequestReceived(requestId, "", numOfValues);

        return requestId;
    }

    function request(string calldata seed, uint256 numOfValues) external returns (uint256 requestId) {
        requestId = uint256(keccak256(abi.encode(msg.sender, counter++)));
        requests[requestId] = Request(seed, numOfValues);
        emit RequestReceived(requestId, seed, numOfValues);

        return requestId;
    }
}
