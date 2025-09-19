package utils

import (
	"crypto/ecdsa"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SetupOnchain() (deployedContract string, conn *ethclient.Client, chainId int64, privateKey *ecdsa.PrivateKey, auth *bind.TransactOpts) {

	var err error
	privateKeyHex := Getenv("PRIVATE_KEY", "")
	deployedContract = Getenv("DEPLOYED_CONTRACT", "")

	blockchainNode := Getenv("BLOCKCHAIN_NODE", "")
	backupNode := Getenv("BACKUP_NODE", "")

	conn, err = MakeEthClient(blockchainNode, backupNode)
	if err != nil {
		log.Fatalf("MakeEthClient: %v", err)
	}

	chainId, err = strconv.ParseInt(Getenv("CHAIN_ID", "100640"), 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse chainId: %v", err)
	}

	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	privateKey, err = crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	return

}
