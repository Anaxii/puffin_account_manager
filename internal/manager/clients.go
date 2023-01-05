package manager

import (
	"puffin_account_manager/pkg/global"
)

func (m *Manager) getAllClients() ([]global.ClientSettings, error) {
	activeClients, err := m.DB.GetClients()
	return activeClients, err
}
