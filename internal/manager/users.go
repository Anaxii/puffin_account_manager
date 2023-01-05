package manager

import log "github.com/sirupsen/logrus"

func (m *Manager) verifyUsers() error {
	clients, err := m.getAllClients()
	if err != nil {
		return err
	}
	log.Println(clients, err)
	for _, c := range clients {
		users, err := m.DB.GetClientUsers(c.UUID)
		if err != nil {
			return err
		}
		for _, v := range users {
			log.Println(v)
		}
	}
	return nil
}
