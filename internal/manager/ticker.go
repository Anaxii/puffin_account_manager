package manager

import "time"

func ticker(seconds int) *time.Ticker {
	return time.NewTicker(time.Second * time.Duration(seconds-time.Now().Second()))
}

func (m *Manager) startTicker() {
	requestsTicker := ticker(m.Interval)
	for {
		<-requestsTicker.C
		verifyUserStatuses()
		requestsTicker = ticker(m.Interval)
	}
}