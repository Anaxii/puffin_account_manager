package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"puffin_account_manager/internal/database"
)

var db database.Database

func StartAPI(port string, _db database.Database) {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET"},
	})

	db = _db

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/client/all", getAllClients).Methods("GET")

	r.Use(mux.CORSMethodMiddleware(r))
	log.Info(fmt.Sprintf("API listening on port %v", port))
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(r)))
}