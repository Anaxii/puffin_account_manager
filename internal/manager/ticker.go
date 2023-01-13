package manager

import "time"

func ticker(seconds int) *time.Ticker {
	return time.NewTicker(time.Second * time.Duration(seconds))
}

func (m *Manager) startVerificationTimer() {
	requestsTicker := ticker(m.Interval)
	clients, _ := m.getAllClients()
	go m.listenForClientChanges(&clients)
	m.verifyUsers(clients)

	for {
		<-requestsTicker.C
		clients, _ = m.getAllClients()
		m.verifyUsers(clients)
		requestsTicker = ticker(m.Interval)
	}
}