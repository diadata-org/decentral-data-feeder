// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract RandomRequestManager {
   
    event RequestReceived(
        uint256 indexed requestId,
        RequestType requestType,
        string seed,
        bytes params
    );
    
    // Overall event for any fulfilled request.
    event RequestFulfilled(uint256 indexed requestId, RequestType requestType);

    // Types of requested randomness.
    enum RequestType {
        Bytes,
        IntRange,
        IntArray
    }

    uint256 public counter;

    // --- Requests ---
    struct RandamuBytesRequest {
        uint256 requestId;
        string seed;
        uint8 lenBytes;
        uint8 numBytes;
        string[] randomness;
        bool fulfilled;
    }

    struct RandamuIntRangeRequest {
        uint256 requestId;
        string seed;
        uint8 min;
        uint8 max;
        uint8 numInts;
        uint256[] randomInts;
        bool fulfilled;
    }

    struct RandamuIntRequest {
        uint256 requestId;
        string seed;
        uint8 bitSize;
        uint8 numInts;
        int256[] randomInts;
        bool fulfilled;
    }

    // Storage of requestType and requests depending on the requestID for later retrieval.
    mapping(uint256 => RequestType) public requestTypes;
    mapping(uint256 => RandamuBytesRequest) public bytesRequests;
    mapping(uint256 => RandamuIntRangeRequest) public intRangeRequests;
    mapping(uint256 => RandamuIntRequest) public intArrayRequests;

    // Helper function for request IDs.
    function _createRequestId(address requester) internal view returns (uint256) {
        return uint256(keccak256(abi.encode(requester, counter)));
    }

    // TO DO: Should we add seed?
    // Dispatcher for all request types.
    function baseRequest(RequestType rType, bytes calldata inputParams)
        external
        returns (uint256 requestId)
    {
        requestId = _createRequestId(msg.sender);
        counter++;
        requestTypes[requestId] = rType;

        if (rType == RequestType.Bytes) {
            (string memory seed, uint8 lenBytes, uint8 numBytes) =
                abi.decode(inputParams, (string, uint8, uint8));

            bytesRequests[requestId] = RandamuBytesRequest({
                requestId: requestId,
                seed: seed,
                lenBytes: lenBytes,
                numBytes: numBytes,
                randomness: new string[](uint256(numBytes)) ,
                fulfilled: false
            });

        } else if (rType == RequestType.IntRange) {
            (string memory seed, uint8 min, uint8 max, uint8 numInts) =
                abi.decode(inputParams, (string, uint8, uint8, uint8));

            intRangeRequests[requestId] = RandamuIntRangeRequest({
                requestId: requestId,
                seed: seed,
                min: min,
                max: max,
                numInts: numInts,
                randomInts: new uint256[](uint256(numInts)),
                fulfilled: false
            });

        } else if (rType == RequestType.IntArray) {
            (string memory seed, uint8 bitSize, uint8 numInts) =
                abi.decode(inputParams, (string, uint8, uint8));

            intArrayRequests[requestId] = RandamuIntRequest({
                requestId: requestId,
                seed: seed,
                bitSize: bitSize,
                numInts: numInts,
                randomInts: new int256[](uint256(numInts)),
                fulfilled: false
            });

        } else {
            revert("Unsupported request type");
        }

        emit RequestReceived(requestId, rType, "", inputParams);
    }

    // Generic fulfillment function.
    function fulfillRequest(uint256 requestId, bytes calldata fulfillmentData)
        external
    {
        RequestType rType = requestTypes[requestId];
        require(uint8(rType) <= 2, "Unknown request type");

        if (rType == RequestType.Bytes) {
            // Expect (string[] randomness)
            string[] memory randomness = abi.decode(fulfillmentData, (string[]));
            RandamuBytesRequest storage req = bytesRequests[requestId];
            require(!req.fulfilled, "Already fulfilled");
            req.randomness = randomness;
            req.fulfilled = true;

        } else if (rType == RequestType.IntRange) {
            // Expect (uint256[] randomInts)
            uint256[] memory randomInts = abi.decode(fulfillmentData, (uint256[]));
            RandamuIntRangeRequest storage req = intRangeRequests[requestId];
            require(!req.fulfilled, "Already fulfilled");
            req.randomInts = randomInts;
            req.fulfilled = true;

        } else if (rType == RequestType.IntArray) {
            // Expect (int256[] randomInts)
            int256[] memory randomInts = abi.decode(fulfillmentData, (int256[]));
            RandamuIntRequest storage req = intArrayRequests[requestId];
            require(!req.fulfilled, "Already fulfilled");
            req.randomInts = randomInts;
            req.fulfilled = true;
        }

        emit RequestFulfilled(requestId, rType);
    }
}
