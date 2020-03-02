package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/aakordas/creature_manager/pkg/creature"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// #C TODO: If the Mongo driver fails to connect to the Mongo daemon, provide
// only roll functionality.

var (
	client          *mongo.Client
	playersDatabase *mongo.Database
	dbContext       *context.Context

	database = "creatures"
	players  = "players"
)

// Connect initializes the interface and connects an application to the provided
// database.
func Connect(db string) {
	client, err := mongo.NewClient(options.Client().ApplyURI(db))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	playersDatabase = client.Database(database)
}

// Disconnect disconnecs the client from the database.
func Disconnect() {
	err := client.Disconnect(*dbContext)
	if err != nil {
		log.Fatal(err)
	}
}

// AddPlayer is the handler that creates new players in the database.
func AddPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerName := vars["player"]
	if playerName == "" {
		unexpectedError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	playersCollection := playersDatabase.Collection(players)

	enc := json.NewEncoder(w)
	// Ensure that the provided name does not already exist in the databae.
	findResult := playersCollection.FindOne(ctx, bson.M{
		"name": playerName,
	})
	// Can this even happen?
	if findResult == nil {
		log.Println("Error in FindOne")
		response := &errorResponse{
			"database error",
			"An error was encountered while accessing the database.",
		}
		jsonEncode(w, enc, response)
		return
	}

	err := findResult.Err()
	if err != mongo.ErrNoDocuments {
		response := &errorResponse{
			"player exists",
			"A player with the provided name already exists in the database.",
		}
		jsonEncode(w, enc, response)
		return
	}

	// A brand new player only has his name associated with them. Everything
	// else will have to be added by subsequent requests.
	entry, err := bson.Marshal(creature.Creature{Name: playerName})
	if err != nil {
		log.Println(err)
		response := &errorResponse{
			"server error",
			"An error was encountered while processing the request.",
		}
		jsonEncode(w, enc, response)
	}
	_, err = playersCollection.InsertOne(ctx, entry)
	if err != nil {
		log.Println(err)
		response := &errorResponse{
			"database error",
			"There was an error inserting the new entry in the database.",
		}
		jsonEncode(w, enc, response)
	}
}
