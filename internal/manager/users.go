package manager

import (
	log "github.com/sirupsen/logrus"
	"math/big"
	"puffin_account_manager/internal/blockchain"
	"puffin_account_manager/pkg/global"
	"strings"
)

func (m *Manager) verifyUsers() error {
	clients, err := m.getAllClients()
	if err != nil {
		return err
	}
	for _, c := range clients {
		for _, product := range c.PackageOptions {
			if product == "geo_block" {
				if c.PuffinGeoAddress == "" {
					continue
				}
				toSet, err := m.handleGeoBlock(c)
				if err != nil {
					log.Error(err)
					continue
				}
				for u, t := range toSet {
					_ = blockchain.SetTier(u, big.NewInt(t), c.PuffinGeoAddress, c.RPCURL, big.NewInt(c.ChainID), m.Config.PrivateKey)
				}
			} else if product == "kyc" {
				if c.PuffinGeoAddress != "" || c.PuffinKYCAddress == "" {
					continue
				}

				toSet, err := m.handleGeoBlock(c)
				if err != nil {
					log.Error(err)
					continue
				}
				for u, t := range toSet {
					if t == 0 {
						_ = blockchain.SetTier(u, big.NewInt(t), c.PuffinGeoAddress, c.RPCURL, big.NewInt(c.ChainID), m.Config.PrivateKey)
					} else if t == 1 {
						_ = blockchain.RemoveUser(u, c.PuffinGeoAddress, c.RPCURL, big.NewInt(c.ChainID), m.Config.PrivateKey)
					}
				}
			}
		}
	}
	return nil
}

func (m *Manager) handleGeoBlock(c global.ClientSettings) (map[string]int64, error) {
	users, err := m.DB.GetClientUsers(c.UUID)
	toSet := map[string]int64{}
	if err != nil {
		return toSet, nil
	}
	for _, v := range users {
		v.Country = strings.ToLower(v.Country)
		userTier, isKYC := blockchain.GetTier(v.User, c.PuffinGeoAddress, c.RPCURL, big.NewInt(c.ChainID))
		for tier, countries := range c.BlockedCountries {
			for _, country := range countries {
				country = strings.ToLower(country)
				if v.Country != country && userTier == 0 && isKYC {
					continue
				}  else if v.Country == country && userTier == 0  {
					_, ok := toSet[v.User]; if ok {
						if tier < toSet[v.User] {
							continue
						}
					}

					toSet[v.User] = tier
					continue
				} else if v.Country != country && !isKYC {
					_, ok := toSet[v.User]; if ok {
						continue
					}
					toSet[v.User] = 0
					continue
				}
			}
		}
	}
	return toSet, nil
}

func (m *Manager) handleKYC(c global.ClientSettings) (map[string]int64, error) {
	users, err := m.DB.GetClientUsers(c.UUID)
	toSet := map[string]int64{}
	if err != nil {
		return toSet, nil
	}
	for _, v := range users {
		v.Country = strings.ToLower(v.Country)
		_, isKYC := blockchain.GetTier(v.User, c.PuffinGeoAddress, c.RPCURL, big.NewInt(c.ChainID))
		if !isKYC && v.Status == "approved" {
			toSet[v.User] = 1
		} else if isKYC && v.Status == "blocked" {
			toSet[v.User] = 0
		}
	}
	return toSet, nil
}
