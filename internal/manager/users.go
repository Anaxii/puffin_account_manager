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
	toSet := map[string]int{}
	for _, c := range clients {
		for _, product := range c.PackageOptions {
			if product == "geo_block" {
				toSet, err = m.handleGeoBlock(c)
				if err != nil {
					log.Error(err)
					continue
				}
			}
		}
		log.Println(toSet)
		toSet = map[string]int{}
	}
	return nil
}

func (m *Manager) handleGeoBlock(c global.ClientSettings) (map[string]int, error) {
	users, err := m.DB.GetClientUsers(c.UUID)
	toSet := map[string]int{}
	if err != nil {
		return toSet, nil
	}
	for _, v := range users {
		v.Country = strings.ToLower(v.Country)
		userTier, isKYC := blockchain.GetTier(v.User, c.PuffinGeoAddress, c.RPCURL, big.NewInt(c.ChainID))
		for tier, countries := range c.BlockedCountries {
			for _, country := range countries {
				log.Println(tier, userTier, isKYC, country)
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
