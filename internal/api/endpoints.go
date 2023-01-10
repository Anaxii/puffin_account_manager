package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func getAllClients(w http.ResponseWriter, r *http.Request) {
	clients, err := db.GetClients()
	if err != nil {
		log.Warn("Failed to decode request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userClients := map[int]bool{}
	wallet, ok := r.URL.Query()["wallet"]; if ok {
		_userClients, err := db.GetUserClients(wallet[0])
		if err == nil {
			userClients = _userClients
		}
	}
	c := map[string]int{}
	for _, v := range clients {
		c[v.ProjectName] = v.UUID
	}
	res, _ := json.Marshal(map[string]interface{}{"clients": c, "user_clients": userClients})
	w.Write(res)
	return
}