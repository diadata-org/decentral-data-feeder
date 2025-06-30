package onchain

import (
	"time"

	diaOracleRandomness "github.com/diadata-org/decentral-data-feeder/contracts/DIAOracleRandomness"
	diaOracleV2MultiupdateService "github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleV2MultiupdateService"
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
		case diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService:
			contract, err = diaOracleV2MultiupdateService.NewDiaOracleV2MultiupdateService(common.HexToAddress(deployedContract), conn)
		case diaOracleRandomness.DIAOracleRandomness:
			contract, err = diaOracleRandomness.NewDIAOracleRandomness(common.HexToAddress(deployedContract), conn)
		}
		if err != nil {
			return nil, err
		}

	} else {

		// deploy contract
		var addr common.Address
		var tx *types.Transaction
		switch contract.(type) {
		case diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService:
			addr, tx, contract, err = diaOracleV2MultiupdateService.DeployDiaOracleV2MultiupdateService(auth, conn)
		case diaOracleRandomness.DIAOracleRandomness:
			addr, tx, contract, err = diaOracleRandomness.DeployDIAOracleRandomness(auth, conn)
		}
		if err != nil {
			log.Fatalf("could not deploy contract: %v", err)
			return nil, err
		}

		log.Infof("Contract pending deploy: 0x%x.", addr)
		log.Infof("Transaction waiting to be mined: 0x%x.", tx.Hash())
		time.Sleep(180000 * time.Millisecond)

	}

	return contract, nil
}

func DeployOrBindContractOld(
	deployedContract string,
	conn *ethclient.Client,
	auth *bind.TransactOpts,
	contract **diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
) error {
	var err error
	if deployedContract != "" {
		// bind primary and backup
		*contract, err = diaOracleV2MultiupdateService.NewDiaOracleV2MultiupdateService(common.HexToAddress(deployedContract), conn)
		if err != nil {
			return err
		}
	} else {
		// deploy contract
		var addr common.Address
		var tx *types.Transaction
		addr, tx, *contract, err = diaOracleV2MultiupdateService.DeployDiaOracleV2MultiupdateService(auth, conn)
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
