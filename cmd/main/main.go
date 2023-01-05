package main

import (
	"puffin_account_manager/internal/manager"
	Log "puffin_account_manager/pkg/log"
)

func main() {
	Log.SetupLogs()
	manager.StartManager()
}
