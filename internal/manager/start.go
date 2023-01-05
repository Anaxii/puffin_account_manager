package manager

import (
	log "github.com/sirupsen/logrus"
	"puffin_account_manager/internal/config"
	"puffin_account_manager/pkg/global"
)

type Manager struct {
	Config   global.Config
	Interval int
}

func StartManager() {
	c := config.GetConfig()
	log.Println(c)

	m := Manager{Config: c, Interval: 60}
	m.startTicker()
}
