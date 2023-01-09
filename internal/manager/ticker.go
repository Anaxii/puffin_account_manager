package manager

import "time"

func ticker(seconds int) *time.Ticker {
	return time.NewTicker(time.Second * time.Duration(seconds-time.Now().Second()))
}

func (m *Manager) startVerificationTimer() {
	requestsTicker := ticker(m.Interval)
	clients, _ := m.getAllClients()
	go m.listenForClientChanges(&clients)
	m.verifyUsers(clients)

	for {
		<-requestsTicker.C
		m.verifyUsers(clients)
		requestsTicker = ticker(m.Interval)
	}
}