package blockchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
)

func getAuth(rpcURL string, chainID *big.Int, _privateKey string) (*ethclient.Client, *bind.TransactOpts, error) {

	conn, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "rpcURL": rpcURL, "chainID": chainID, "file": "Blockchain:getAuth"}).Error("Failed to connect to the Ethereum client")
		return &ethclient.Client{}, &bind.TransactOpts{}, err
	}

	privateKey, err := crypto.HexToECDSA(_privateKey)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "rpcURL": rpcURL, "chainID": chainID, "file": "Blockchain:getAuth"}).Error("Failed to convert private key string to ECDSA")
		return &ethclient.Client{}, &bind.TransactOpts{}, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "rpcURL": rpcURL, "chainID": chainID, "file": "Blockchain:getAuth"}).Error("Failed to create authorized transactor")
		return &ethclient.Client{}, &bind.TransactOpts{}, err
	}

	return conn, auth, err
}
