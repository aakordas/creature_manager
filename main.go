package main

import (
	"log"
	"net/http"
	"time"

	"github.com/aakordas/creature_manager/pkg/server"
	"github.com/gorilla/mux"
)

// #A TODO: Add some flag to set the host for the API and the database.

// #C TODO: If the Mongo driver fails to connect to the Mongo daemon, provide
// only roll functionality.

var address = "127.0.0.1:8080"
var databaseAddress = "mongodb://127.0.0.1:27017" // MongoDB's default port.

func main() {
	server.Connect(databaseAddress)
	defer server.Disconnect()

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1/").Subrouter()

	// Rolls
	api.HandleFunc("/roll", server.Roll)
	api.Queries("sides", "{sides:[0-9]+}", "count", "{count:[0-9]+}").HandlerFunc(server.Roll)

	roll := api.PathPrefix("/roll/").Subrouter()
	roll.HandleFunc("/{sides:[d|D][0-9]+}", server.RollN)

	dRoll := roll.PathPrefix("/{sides:[d|D][0-9]+}/").Subrouter()
	dRoll.HandleFunc("/{count:[0-9]+}", server.DRollN)
	dRoll.Queries("count", "{count:[0-9]+}").HandlerFunc(server.RollN)

	// Player
	add := api.PathPrefix("/add/").Subrouter()
	add.HandleFunc("/player/{player:[a-zA-Z ]+}", server.AddPlayer).Methods("PUT")

	get := api.PathPrefix("/get/").Subrouter()
	get.HandleFunc("/player/{player:[a-zA-Z ]+}", server.GetPlayer)

	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server...")
	log.Fatal(srv.ListenAndServe())
}
