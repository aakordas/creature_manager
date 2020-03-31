package server

import (
	"context"
	"log"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect initializes the interface and connects an application to the provided
// database.
func Connect(db string) *mux.Router {
	opts := options.Client().ApplyURI(db)
	v := opts.Validate()
	if v != nil {
		log.Fatal(v)
	}
	client, err := mongo.NewClient(opts)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: context.WithClose. Return the close function and pass it in Disconnect.
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	playersDatabase = client.Database(database)

	r := mux.NewRouter()
	r = diceRoutes(r)
	r = playerRoutes(r)

	return r
}

// Disconnect disconnecs the client from the database.
func Disconnect() {
	err := client.Disconnect(*dbContext)
	if err != nil {
		log.Fatal(err)
	}
}
