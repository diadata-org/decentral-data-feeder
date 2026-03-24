package onchain

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type NonceManager struct {
	mu        sync.Mutex
	nextNonce uint64
	initDone  bool
}

func (nm *NonceManager) NextNonce(conn *ethclient.Client, from common.Address) (uint64, error) {
	nm.mu.Lock()
	defer nm.mu.Unlock()

	if !nm.initDone {
		nonce, err := conn.PendingNonceAt(context.Background(), from)
		if err != nil {
			return 0, err
		}
		nm.nextNonce = nonce
		nm.initDone = true
	}

	nonce := nm.nextNonce
	nm.nextNonce++
	return nonce, nil
}

func (nm *NonceManager) BuildTxOpts(conn *ethclient.Client, auth *bind.TransactOpts) (*bind.TransactOpts, error) {
	nonce, err := nm.NextNonce(conn, auth.From)
	if err != nil {
		return nil, err
	}

	return &bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Nonce:  big.NewInt(int64(nonce)),
	}, nil
}
