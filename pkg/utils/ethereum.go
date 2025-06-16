package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

func MakeEthClient(mainNode string, backupNode string) (conn *ethclient.Client, err error) {
	var useBackupNode bool

	conn, err = ethclient.Dial(mainNode)
	if err != nil {
		useBackupNode = true
	} else {
		syncProgress, err := conn.SyncProgress(context.Background())
		if err != nil {
			log.Warnf("main rpc node not able to get sync state: %v. Switch to backup node.", err)
			useBackupNode = true
		} else if syncProgress != nil {
			log.Warnf("main rpc node not synced. Switch to backup node.")
			useBackupNode = true
		}
	}

	if useBackupNode {
		conn, err = ethclient.Dial(backupNode)
		if err != nil {
			err = fmt.Errorf("failed to connect to the backup rpc node: %v", err)
			return
		}
		var syncProgress *ethereum.SyncProgress
		syncProgress, err = conn.SyncProgress(context.Background())
		if err != nil {
			err = fmt.Errorf("backup rpc node not able to get sync state: %v", err)
			return
		}
		if syncProgress != nil {
			err = errors.New("backup rpc node not synced")
			return
		}
	}

	return
}
