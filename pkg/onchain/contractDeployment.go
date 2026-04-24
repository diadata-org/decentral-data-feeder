package onchain

import (
	"time"

	diaOracleRandomness "github.com/diadata-org/decentral-data-feeder/contracts/DIAOracleRandomness"
	diaoraclev3 "github.com/diadata-org/lumina-library/contracts/lumina/diaoraclev3"
	"github.com/diadata-org/lumina-library/onchain"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployOrBindContract(
	deployedContract string,
	conn *ethclient.Client,
	auth *bind.TransactOpts,
	contract any,
) (any, error) {
	var err error
	if deployedContract != "" {

		// bind primary and backup
		switch contract.(type) {
		case diaoraclev3.DIAOracleV3:
			var contractPtr *diaoraclev3.DIAOracleV3
			var contractBackup *diaoraclev3.DIAOracleV3
			err = onchain.DeployOrBindContract(deployedContract, conn, conn, auth, &contractPtr, &contractBackup, 18)
			if err != nil {
				return nil, err
			}

			contract = *contractPtr
		case diaOracleRandomness.DIAOracleRandomness:
			contract, err = diaOracleRandomness.NewDIAOracleRandomness(common.HexToAddress(deployedContract), conn)
		}

	} else {

		// deploy contract
		switch contract.(type) {
		case diaoraclev3.DIAOracleV3:
			var contractPtr *diaoraclev3.DIAOracleV3
			var contractBackup *diaoraclev3.DIAOracleV3
			err = onchain.DeployOrBindContract("", conn, conn, auth, &contractPtr, &contractBackup, 18)
			if err != nil {
				log.Fatalf("could not deploy contract: %v", err)
				return nil, err
			}
			contract = *contractPtr
		case diaOracleRandomness.DIAOracleRandomness:
			var addr common.Address
			var tx *types.Transaction
			addr, tx, contract, err = diaOracleRandomness.DeployDIAOracleRandomness(auth, conn)
			if err != nil {
				log.Fatalf("could not deploy contract: %v", err)
				return nil, err
			}
			log.Infof("Contract pending deploy: 0x%x.", addr)
			log.Infof("Transaction waiting to be mined: 0x%x.", tx.Hash())
			time.Sleep(180000 * time.Millisecond)
		}

	}

	return contract, nil
}

func DeployOrBindContractOld(
	deployedContract string,
	conn *ethclient.Client,
	auth *bind.TransactOpts,
	contract **diaoraclev3.DIAOracleV3,
) error {
	var err error
	if deployedContract != "" {
		// bind primary and backup
		*contract, err = diaoraclev3.NewDIAOracleV3(common.HexToAddress(deployedContract), conn)
		if err != nil {
			return err
		}
	} else {
		// deploy contract
		var addr common.Address
		var tx *types.Transaction
		addr, tx, *contract, err = diaoraclev3.DeployDIAOracleV3(auth, conn)
		if err != nil {
			log.Fatalf("could not deploy contract: %v", err)
			return err
		}
		log.Infof("Contract pending deploy: 0x%x.", addr)
		log.Infof("Transaction waiting to be mined: 0x%x.", tx.Hash())
		time.Sleep(180000 * time.Millisecond)
	}
	return nil
}
