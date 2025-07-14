// compiled using solidity 0.8.30
pragma solidity 0.8.30;

contract DIAOracleRandomness {

    struct BytesEntry {
        string requestId;
        string[] randomness;
        string seed;
        string signature;
        bool exists;
    }

    struct IntRangeEntry {
        string requestId;
        uint256[] randomInts;
        string seed;
        string signature;
        bool exists;
    }

    struct IntArrayEntry {
        string requestId;
        int256[] randomInts;
        string seed;
        string signature;
        bool exists;
    }

    mapping(string => BytesEntry) private bytesEntries;
    mapping(string => IntRangeEntry) private intRangeEntries;
    mapping(string => IntArrayEntry) private intArrayEntries;
   
    address oracleUpdater;
    
    event BytesSet(string requestId, int256 indexed round, string seed, string signature);
    event IntRangeSet(string requestId, int256 indexed round, string seed, string signature);
    event IntArraySet(string requestId, int256 indexed round, string seed, string signature);
    event UpdaterAddressChange(address newUpdater);
    
    constructor() {
        oracleUpdater = msg.sender;
    }
    
    // SetBytes
    function setBytes(
        string calldata requestId,
        string[] calldata randomness,
        int64 round,
        string calldata seed,
        string calldata signature
    ) external {
        bytesEntries[requestId] = BytesEntry({
            requestId: requestId,
            randomness: randomness,
            seed: seed,
            signature: signature,
            exists: true
        });
        emit BytesSet(requestId, round, seed, signature);
    }

    // ReadBytes: returns full array, seed, and signature
    function getBytes(string calldata requestId_) external view returns (
        string memory requestId,
        string[] memory randomness,
        string memory seed,
        string memory signature
    ) {
        require(bytesEntries[requestId_].exists, "No BytesEntry for this round");
        BytesEntry storage entry = bytesEntries[requestId_];
        return (entry.requestId, entry.randomness, entry.seed, entry.signature);
    }

    // SetIntRange
    function setIntRange(
        string calldata requestId,
        uint256[] calldata randomInts,
        int64 round,
        string calldata seed,
        string calldata signature
    ) external {
        intRangeEntries[requestId] = IntRangeEntry({
            requestId: requestId,
            randomInts: randomInts,
            seed: seed,
            signature: signature,
            exists: true
        });
        emit IntRangeSet(requestId, round, seed, signature);
    }

    // ReadIntRange
    function getIntRange(string calldata requestId_) external view returns (
        string memory requestId,
        uint256[] memory randomInts,
        string memory seed,
        string memory signature
    ) {
        require(intRangeEntries[requestId_].exists, "No IntRangeEntry for this round");
        IntRangeEntry storage entry = intRangeEntries[requestId_];
        return (entry.requestId, entry.randomInts, entry.seed, entry.signature);
    }

    // SetIntArray
    function setIntArray(
        string calldata requestId,
        int256[] calldata randomInts,
        int64 round,
        string calldata seed,
        string calldata signature
    ) external {
        intArrayEntries[requestId] = IntArrayEntry({
            requestId: requestId,
            randomInts: randomInts,
            seed: seed,
            signature: signature,
            exists: true
        });
        emit IntArraySet(requestId, round, seed, signature);
    }

    // ReadIntArray
    function getIntArray(string calldata requestId_) external view returns (
        string memory requestId,
        int256[] memory randomInts,
        string memory seed,
        string memory signature
    ) {
        require(intArrayEntries[requestId_].exists, "No IntArrayEntry for this round");
        IntArrayEntry storage entry = intArrayEntries[requestId_];
        return (entry.requestId, entry.randomInts, entry.seed, entry.signature);
    }
    
    function updateOracleUpdaterAddress(address newOracleUpdaterAddress) public {
        require(msg.sender == oracleUpdater);
        oracleUpdater = newOracleUpdaterAddress;
        emit UpdaterAddressChange(newOracleUpdaterAddress);
    }
}

