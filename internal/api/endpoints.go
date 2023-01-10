package api

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"puffin_account_manager/pkg/global"
	"strconv"
	"strings"
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
		_userClients, err := db.GetUserClients(strings.ToLower(wallet[0]))
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


func joinClient(w http.ResponseWriter, r *http.Request) {
	userClients := map[int]bool{}
	id := ""
	wallet := ""
	_wallet, ok := r.URL.Query()["wallet"]; if ok {
		wallet = strings.ToLower(_wallet[0])
	} else {
		res, _ := json.Marshal(map[string]interface{}{"error": "user did not supply wallet address"})
		w.Write(res)
		return
	}

	userClients, err := db.GetUserClients(wallet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, ok = r.URL.Query()["id"]; if ok {
		id = r.URL.Query()["id"][0]
	} else {
		res, _ := json.Marshal(map[string]interface{}{"error": "user did not supply wallet address"})
		w.Write(res)
		return
	}

	for k, _ := range userClients {
		if id == fmt.Sprintf("%v", k) {
			err := db.DeleteUser(wallet, k)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				res, _ := json.Marshal(map[string]interface{}{"error": "failed to remove user"})
				w.Write(res)
				return
			}
			res, _ := json.Marshal(map[string]interface{}{"status": "successfully left"})
			w.Write(res)
			return
		}
	}

	_id, err := strconv.Atoi(id)
	if err != nil {
		res, _ := json.Marshal(map[string]interface{}{"error": "invalid id"})
		w.Write(res)
		return
	}
	user, err := db.GetUser(wallet)
	if err != nil {
		res, _ := json.Marshal(map[string]interface{}{"error": "invalid user"})
		w.Write(res)
		return
	}
	err = db.AddClientUser(global.ClientUsers{
		User: wallet,
		Client: _id,
		Country: user.Country,
		Status: "approved",
	})

	res, _ := json.Marshal(map[string]interface{}{"status": "success"})
	w.Write(res)
	return
}