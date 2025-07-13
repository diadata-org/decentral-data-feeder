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

    mapping(int256 => BytesEntry) private bytesEntries;
    mapping(int256 => IntRangeEntry) private intRangeEntries;
    mapping(int256 => IntArrayEntry) private intArrayEntries;
   
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
        bytesEntries[round] = BytesEntry({
            requestId: requestId,
            randomness: randomness,
            seed: seed,
            signature: signature,
            exists: true
        });
        emit BytesSet(requestId, round, seed, signature);
    }

    // ReadBytes: returns full array, seed, and signature
    function getBytes(int64 round) external view returns (
        string memory requestId,
        string[] memory randomness,
        string memory seed,
        string memory signature
    ) {
        require(bytesEntries[round].exists, "No BytesEntry for this round");
        BytesEntry storage entry = bytesEntries[round];
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
        intRangeEntries[round] = IntRangeEntry({
            requestId: requestId,
            randomInts: randomInts,
            seed: seed,
            signature: signature,
            exists: true
        });
        emit IntRangeSet(requestId, round, seed, signature);
    }

    // ReadIntRange
    function getIntRange(int64 round) external view returns (
        string memory requestId,
        uint256[] memory randomInts,
        string memory seed,
        string memory signature
    ) {
        require(intRangeEntries[round].exists, "No IntRangeEntry for this round");
        IntRangeEntry storage entry = intRangeEntries[round];
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
        intArrayEntries[round] = IntArrayEntry({
            requestId: requestId,
            randomInts: randomInts,
            seed: seed,
            signature: signature,
            exists: true
        });
        emit IntArraySet(requestId, round, seed, signature);
    }

    // ReadIntArray
    function getIntArray(int64 round) external view returns (
        string memory requestId,
        int256[] memory randomInts,
        string memory seed,
        string memory signature
    ) {
        require(intArrayEntries[round].exists, "No IntArrayEntry for this round");
        IntArrayEntry storage entry = intArrayEntries[round];
        return (entry.requestId, entry.randomInts, entry.seed, entry.signature);
    }
    
    function updateOracleUpdaterAddress(address newOracleUpdaterAddress) public {
        require(msg.sender == oracleUpdater);
        oracleUpdater = newOracleUpdaterAddress;
        emit UpdaterAddressChange(newOracleUpdaterAddress);
    }
}

