package blockchain

import (
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
)

func CheckIfIsApproved(walletAddress string, rpcURL string, chainId *big.Int) bool {
	conn, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to connect to the Ethereum client")
	}

	verify, err := abi.NewPuffinApprovedAccounts(common.HexToAddress(config.AvaxChainApprovedAccountsAddress), conn)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to instantiate PuffinApprovedAccounts contract")
	}

	isApproved, err := verify.IsApproved(nil, common.HexToAddress(walletAddress))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to check if user is approved")
		return false
	}

	conn, err = ethclient.Dial(config.PuffinRpcURL)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to connect to puffin RPC")
		return false
	}

	verifyPuffin, err := abi.NewAllowListInterface(common.HexToAddress(config.PuffinAllowListInterfaceURL), conn)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to instantiate AllowListInterface")
		return false
	}

	isEnabled, err := verifyPuffin.ReadAllowList(nil, common.HexToAddress(walletAddress))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to call ReadAllowList")
		return false
	}

	return isApproved && isEnabled != big.NewInt(0)
}

func ApproveAddress(walletAddress string, rpcURL string, chainId *big.Int) error {
	conn, auth, err := getAuth(config.AvaxRpcURL, config.AvaxChainId)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:ApproveAddress"}).Error("Failed to get auth")
		return err
	}

	verify, err := abi.NewPuffinApprovedAccounts(common.HexToAddress(config.AvaxChainApprovedAccountsAddress), conn)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:ApproveAddress"}).Error("Failed to instantiate PuffinApprovedAccounts contract")
		return err
	}

	_, err = verify.Approve(auth, common.HexToAddress(walletAddress))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:ApproveAddress"}).Error("Failed to call approve")
		return err
	}

	return nil
}
