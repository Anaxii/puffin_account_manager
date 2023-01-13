package manager

import (
	log "github.com/sirupsen/logrus"
	"puffin_account_manager/internal/api"
	"puffin_account_manager/internal/config"
	"puffin_account_manager/internal/database"
	"puffin_account_manager/pkg/global"
)

type Manager struct {
	Config   global.Config
	Interval int
	DB       database.Database
}

func StartManager() {
	c := config.GetConfig()
	log.Println(c)

	d := database.Database{DBURI: c.MongoDbURI}
	m := Manager{
		Config:   c,
		Interval: 3,
		DB: d,
	}

	go api.StartAPI("8081", d)

	m.startVerificationTimer()
}
