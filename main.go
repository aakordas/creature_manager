package main

import (
	"log"
	"net/http"
	"time"

	"github.com/aakordas/creature_manager/pkg/server"
	"github.com/gorilla/mux"
)

var address = "127.0.0.1:8080"

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1/").Subrouter()

	api.HandleFunc("/roll", server.Roll)
	api.Queries("sides", "{sides:[0-9]+}", "count", "{count:[0-9]+}").HandlerFunc(server.Roll)

	roll := api.PathPrefix("/roll/").Subrouter()
	roll.HandleFunc("/{sides:[d|D][0-9]+}", server.RollN)

	dRoll := roll.PathPrefix("/{sides:[d|D][0-9]+}/").Subrouter()
	dRoll.HandleFunc("/{count:[0-9]+}", server.DRollN)
	dRoll.Queries("count", "{count:[0-9]+}").HandlerFunc(server.RollN)

	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server...")
	log.Fatal(srv.ListenAndServe())
}
