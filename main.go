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

	number  = "{number:[0-9]+}"
	dsides  = "{sides:[d|D][0-9]+}"
	count   = "{count:[0-9]+}"
	name    = "{name:[a-zA-Z ]+}"
	ability = "{ability:[a-zA-Z]+}"
	skill   = "{skill:[a-zA-Z_]+}"
	save = "{save:[a-zA-Z]+}"
)

func main() {
	server.Connect(databaseAddress)
	defer server.Disconnect()

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1/").Subrouter()

	// Rolls
	api.HandleFunc("/roll", server.Roll)
	api.Queries("sides", number, "count", count).HandlerFunc(server.Roll)

	roll := api.PathPrefix("/roll/").Subrouter()
	roll.HandleFunc("/"+dsides, server.RollN)

	dRoll := roll.PathPrefix("/" + dsides + "/").Subrouter()
	dRoll.HandleFunc("/"+count, server.DRollN)
	dRoll.Queries("count", count).HandlerFunc(server.RollN)

	// Player
	player := api.PathPrefix("/player").Subrouter()
	player.HandleFunc("/"+name, server.AddPlayer).Methods("PUT")
	player.HandleFunc("/"+name, server.GetPlayer)

        // Cannot (?) create subrouters with variables, like `name'.
        playerName := "/"+name+"/"
	player.HandleFunc(playerName+"hitpoints/"+number, server.SetHitPoints).Methods("PUT")
	player.HandleFunc(playerName+"level/"+number, server.SetLevel).Methods("PUT")
	player.HandleFunc(playerName+"armor/"+number, server.SetArmorClass).Methods("PUT")

	// Player's abilities
	player.HandleFunc(playerName+"abilities/"+ability+"/"+number, server.SetAbility).Methods("PUT")
	player.HandleFunc(playerName+"abilities", server.GetAbilities)

	// Player's skills
	player.HandleFunc(playerName+"skills/"+skill, server.SetSkill).Methods("PUT")
	player.HandleFunc(playerName+"skills", server.GetSkills)

        // Player's saving throws
        player.HandleFunc(playerName+"saving_throw/"+save, server.SetSave).Methods("PUT")
        player.HandleFunc(playerName+"saving_throw", server.GetSaves)

	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server...")
	log.Fatal(srv.ListenAndServe())
}
