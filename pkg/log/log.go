package log

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func SetupLogs() {
	f, err := os.OpenFile("logs.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	//defer f.Close()
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true})

}
