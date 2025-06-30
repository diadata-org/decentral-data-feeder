// compiled using solidity 0.8.19
pragma solidity 0.8.19;

contract DIAOracleRandomness {

    struct BytesEntry {
        string[] randomness;
        string seed;
        string signature;
        bool exists;
    }

    struct IntRangeEntry {
        uint256[] randomInts;
        string seed;
        string signature;
        bool exists;
    }

    struct IntArrayEntry {
        int256[] randomInts;
        string seed;
        string signature;
        bool exists;
    }

    mapping(int256 => BytesEntry) private bytesEntries;
    mapping(int256 => IntRangeEntry) private intRangeEntries;
    mapping(int256 => IntArrayEntry) private intArrayEntries;
   
    address oracleUpdater;
    
    event BytesSet(int256 indexed round, string seed, string signature);
    event IntRangeSet(int256 indexed round, string seed, string signature);
    event IntArraySet(int256 indexed round, string seed, string signature);
    event UpdaterAddressChange(address newUpdater);
    
    constructor() {
        oracleUpdater = msg.sender;
    }
    
    // SetBytes
    function setBytes(
        string[] calldata randomness,
        int64 round,
        string calldata seed,
        string calldata signature
    ) external {
        bytesEntries[round] = BytesEntry({
            randomness: randomness,
            seed: seed,
            signature: signature,
            exists: true
        });
        emit BytesSet(round, seed, signature);
    }

    // ReadBytes: returns full array, seed, and signature
    function getBytes(int64 round) external view returns (
        string[] memory randomness,
        string memory seed,
        string memory signature
    ) {
        require(bytesEntries[round].exists, "No BytesEntry for this round");
        BytesEntry storage entry = bytesEntries[round];
        return (entry.randomness, entry.seed, entry.signature);
    }

    // SetIntRange
    function setIntRange(
        uint256[] calldata randomInts,
        int64 round,
        string calldata seed,
        string calldata signature
    ) external {
        intRangeEntries[round] = IntRangeEntry({
            randomInts: randomInts,
            seed: seed,
            signature: signature,
            exists: true
        });
        emit IntRangeSet(round, seed, signature);
    }

    // ReadIntRange
    function getIntRange(int64 round) external view returns (
        uint256[] memory randomInts,
        string memory seed,
        string memory signature
    ) {
        require(intRangeEntries[round].exists, "No IntRangeEntry for this round");
        IntRangeEntry storage entry = intRangeEntries[round];
        return (entry.randomInts, entry.seed, entry.signature);
    }

    // SetIntArray
    function setIntArray(
        int256[] calldata randomInts,
        int64 round,
        string calldata seed,
        string calldata signature
    ) external {
        intArrayEntries[round] = IntArrayEntry({
            randomInts: randomInts,
            seed: seed,
            signature: signature,
            exists: true
        });
        emit IntArraySet(round, seed, signature);
    }

    // ReadIntArray
    function getIntArray(int64 round) external view returns (
        int256[] memory randomInts,
        string memory seed,
        string memory signature
    ) {
        require(intArrayEntries[round].exists, "No IntArrayEntry for this round");
        IntArrayEntry storage entry = intArrayEntries[round];
        return (entry.randomInts, entry.seed, entry.signature);
    }
    
    function updateOracleUpdaterAddress(address newOracleUpdaterAddress) public {
        require(msg.sender == oracleUpdater);
        oracleUpdater = newOracleUpdaterAddress;
        emit UpdaterAddressChange(newOracleUpdaterAddress);
    }
}

