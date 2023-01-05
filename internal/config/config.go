package config

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"puffin_account_manager/pkg/global"
	"puffin_account_manager/pkg/util"
)

type configStruct struct {
	PrivateKey                       string `json:"private_key"`
	Port                             string `json:"port"`
	MongoDbURI                       string `json:"mongo_db_uri"`
}

func GetConfig() global.Config {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.WithFields(log.Fields{"error": err.Error(), "file": "config:init"}).Fatal("Failed to generate a private key")
		}
		privateKeyBytes := crypto.FromECDSA(privateKey)

		file, _ := json.MarshalIndent(configStruct{
			PrivateKey:                       fmt.Sprintf("%v", hexutil.Encode(privateKeyBytes)[2:]),
			Port:                             "80",
		}, "", "  ")
		_ = ioutil.WriteFile("config.json", file, 0644)
		log.WithFields(log.Fields{ "file": "config:init"}).Fatal("Generated config.json | Fill in empty data and run again")

	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "config:init"}).Fatal("Config file is invalid")
	}

	var config configStruct
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "config:init"}).Fatal("Could not unmarshal config file")
	}

	var c global.Config

	c.PrivateKey = config.PrivateKey

	if config.PrivateKey != "" {
		_publicKey, _ := util.GenerateECDSAKey(config.PrivateKey)
		c.PublicKey = _publicKey
	}

	c.MongoDbURI = config.MongoDbURI
	return c
}

