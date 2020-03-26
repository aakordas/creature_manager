package main

import (
	"log"
	"net/http"
	"time"

	"github.com/aakordas/creature_manager/pkg/server"
)

// #A TODO: Create a database index on the player name.
// #A TODO: Add some flag to set the host for the API and the database.

// #B TODO: Make some subrouter for /{name}/

// #C TODO: If the Mongo driver fails to connect to the Mongo daemon, provide
// only roll functionality.

func main() {
	var (
		serverPort      = `:8080`
		serverAddress   = `127.0.0.1` + serverPort
		databasePort    = `:27017` // Mongo's default port
		databaseAddress = `mongodb://127.0.0.1` + databasePort
	)

	r := server.Connect(databaseAddress)
	defer server.Disconnect()

	srv := &http.Server{
		Handler:      r,
		Addr:         serverAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server...")
	log.Fatal(srv.ListenAndServe())
}
