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

var (
	address         = "127.0.0.1:8080"
	databaseAddress = "mongodb://127.0.0.1:27017" // MongoDB's default port.

	sides  = "{sides:[0-9]+}"
	dsides = "{sides:[d|D][0-9]+}"
	count  = "{count:[0-9]+}"
	name   = "{name:[a-zA-Z ]+}"
)

func main() {
	server.Connect(databaseAddress)
	defer server.Disconnect()

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1/").Subrouter()

	// Rolls
	api.HandleFunc("/roll", server.Roll)
	api.Queries("sides", sides, "count", count).HandlerFunc(server.Roll)

	roll := api.PathPrefix("/roll/").Subrouter()
	roll.HandleFunc("/"+dsides, server.RollN)

	dRoll := roll.PathPrefix("/" + dsides + "/").Subrouter()
	dRoll.HandleFunc("/"+count, server.DRollN)
	dRoll.Queries("count", count).HandlerFunc(server.RollN)

	// Player
	player := api.PathPrefix("/player").Subrouter()
	player.HandleFunc("/"+name, server.AddPlayer).Methods("PUT")
	player.Queries("name", name).HandlerFunc(server.AddPlayer).Methods("PUT")
	player.HandleFunc("/"+name, server.GetPlayer)
	player.Queries("name", name).HandlerFunc(server.GetPlayer)

	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server...")
	log.Fatal(srv.ListenAndServe())
}
