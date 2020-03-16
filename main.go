package main

import (
	"log"
	"net/http"
	"time"

	"github.com/aakordas/creature_manager/pkg/server"
	"github.com/gorilla/mux"
)

// #A TODO: Create a database index on the player name.
// #A TODO: Add some flag to set the host for the API and the database.

// #B TODO: Make some subrouter for /{name}/

// #C TODO: If the Mongo driver fails to connect to the Mongo daemon, provide
// only roll functionality.

var (
	address         = "127.0.0.1:8080"
	databaseAddress = "mongodb://127.0.0.1:27017" // MongoDB's default port.

	sides   = "{sides:[0-9]+}"
	dsides  = "{sides:[d|D][0-9]+}"
	count   = "{count:[0-9]+}"
	name    = "{name:[a-zA-Z ]+}"
	number  = "{number:[0-9]+}"
	ability = "{ability:[a-zA-Z]+}"
	skill   = "{skill:[a-zA-Z_]+}"
	save    = "{save:[a-zA-Z]+}"
)

func main() {
	server.Connect(databaseAddress)
	defer server.Disconnect()

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1/").Subrouter()

	// Rolls
	api.HandleFunc("/roll", server.Roll)
	api.Queries("sides", sides, "count", count).HandlerFunc(server.Roll).Methods(http.MethodGet)

	roll := api.PathPrefix("/roll/").Subrouter()
	roll.HandleFunc("/"+dsides, server.RollN).Methods(http.MethodGet)

	dRoll := roll.PathPrefix("/" + dsides + "/").Subrouter()
	dRoll.HandleFunc("/"+count, server.DRollN).Methods(http.MethodGet)
	dRoll.Queries("count", count).HandlerFunc(server.RollN)

	// Player
	player := api.PathPrefix("/player").Subrouter()
	player.HandleFunc("/"+name, server.AddPlayer).Methods(http.MethodPut)
	player.HandleFunc("/"+name, server.GetPlayer).Methods(http.MethodGet)
	player.HandleFunc("/"+name, server.DeletePlayer).Methods(http.MethodDelete)

	// Cannot (?) create subrouters with variables, like `name'.
	playerName := "/" + name + "/"
	player.HandleFunc(playerName+"hitpoints/"+number, server.SetHitPoints).Methods(http.MethodPut)
	player.HandleFunc(playerName+"level/"+number, server.SetLevel).Methods(http.MethodPut)
	player.HandleFunc(playerName+"armor/"+number, server.SetArmorClass).Methods(http.MethodPut)

	// Player's abilities
	player.HandleFunc(playerName+"abilities/"+ability+"/"+number, server.SetAbility).Methods(http.MethodPut)
	player.HandleFunc(playerName+"abilities", server.GetAbilities).Methods(http.MethodGet)

	// Player's skills
	player.HandleFunc(playerName+"skills/"+skill, server.SetSkill).Methods(http.MethodPut)
	player.HandleFunc(playerName+"skills", server.GetSkills).Methods(http.MethodGet)

	// Player's saving throws
	player.HandleFunc(playerName+"saving_throws/"+save, server.SetSave).Methods(http.MethodPut)
	player.HandleFunc(playerName+"saving_throws", server.GetSaves).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server...")
	log.Fatal(srv.ListenAndServe())
}
