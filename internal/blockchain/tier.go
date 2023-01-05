package blockchain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"puffin_account_manager/pkg/abi"
)

func GetTier(walletAddress string, contractAddress string, rpcurl string, chainID *big.Int) (int64, bool) {

	conn, err := ethclient.Dial(rpcurl)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to connect to the Ethereum client")
	}

	core, err := abi.NewPuffinCore(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to instantiate PuffinApprovedAccounts contract")
	}

	tier, err := core.Tier(nil, common.HexToAddress(walletAddress))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to check if user is approved")
		return 0, false
	}

	isKYC, err := core.IsKYC(nil, common.HexToAddress(walletAddress))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:CheckIfIsApproved"}).Error("Failed to check if user is approved")
		return 0, false
	}
	return tier.Int64(), isKYC
}

//func SetTier(walletAddress string, tier *big.Int) error {
//
//	conn, auth, err := getAuth(config.AvaxRpcURL, config.AvaxChainId)
//	if err != nil {
//		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:EnableOnPuffin"}).Error("Failed to get auth")
//		return err
//	}
//
//	verify, err := abi.NewPuffinCore(common.HexToAddress(config.PuffinCoreAddress), conn)
//	if err != nil {
//		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:EnableOnPuffin"}).Error("Failed to initialize AllowListInterface")
//		return err
//	}
//
//	_, err = verify.SetTier(auth, common.HexToAddress(walletAddress), tier)
//	if err != nil {
//		log.WithFields(log.Fields{"error": err.Error(), "file": "Blockchain:EnableOnPuffin"}).Error("Failed to call SetEnabled")
//		return err
//	}
//
//	return nil
//}
