package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aakordas/creature_manager/pkg/abilities"
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

	contextTimeout = 5 * time.Second
)

const (
	serverError            = "server error"
	databaseError          = "database error"
	playerNotFoundError    = "player not found"
	invalidPlayerNameError = "invalid player name"
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

// sendErrorResponse creates and sends a custom error response.
func sendErrorResponse(w http.ResponseWriter, enc *json.Encoder, s, l string, h int) {
	w.WriteHeader(h)
	response := &errorResponse{s, l}
	jsonEncode(w, enc, response)
}

// findPlayer returns the SingleResult of looking up the name of a single player
// in the database.
func findPlayer(ctx context.Context, col *mongo.Collection, name string) *mongo.SingleResult {
	findResult := col.FindOne(ctx, bson.M{
		"name": name,
	})

	return findResult
}

// AddPlayer is the handler that creates new players in the database.
func AddPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)
	playerName := vars["name"]
	if playerName == "" {
		sendErrorResponse(w, enc, invalidPlayerNameError,
			"A player's name should contain only characters and spaces.",
			http.StatusBadRequest,
		)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	playersCollection := playersDatabase.Collection(players)

	// Ensure that the provided name does not already exist in the databae.
	findResult := findPlayer(ctx, playersCollection, playerName)
	// Can this even happen?
	if findResult == nil {
		log.Println("Error in FindOne")
		sendErrorResponse(w, enc, databaseError,
			"An error was encountered while accessing the database",
			http.StatusInternalServerError,
		)
		return
	}

	if findResult.Err() != mongo.ErrNoDocuments {
		sendErrorResponse(w, enc,
			"player exists",
			"A player with the provided name already exists in the database.",
			http.StatusBadRequest,
		)
		return
	}

	// A brand new player is of first level, with the initial proficiency
	// bonus of +2 and only has their name associated with them. Everything
	// else will have to be added by subsequent requests.
	entry, err := bson.Marshal(creature.Creature{
		Name:             playerName,
		Level:            1,
		ProficiencyBonus: creature.ProficiencyBonusPerLevel[1],
	})
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, enc, serverError,
			"An error was encountered while processing the request.",
			http.StatusBadRequest,
		)
		return
	}
	_, err = playersCollection.InsertOne(ctx, entry)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, enc, databaseError,
			"There was an error inserting the new entry in the database",
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// emptyResponse models a response with an empty object.
type emptyResponse struct {
}

// GetPlayer is the handler that returns the information about a player in the
// database.
func GetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)
	playerName := vars["name"]
	if playerName == "" {
		sendErrorResponse(w, enc, invalidPlayerNameError,
			"A player's name should contain only characters and spaces.",
			http.StatusBadRequest,
		)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	playersCollection := playersDatabase.Collection(players)

	findResult := findPlayer(ctx, playersCollection, playerName)
	// Can this even happen?
	if findResult == nil {
		log.Println("Error in FindOne")
		sendErrorResponse(w, enc, databaseError,
			"An error was encountered while accessing the database",
			http.StatusInternalServerError,
		)
		return
	}

	var response creature.Creature
	err := findResult.Decode(&response)
	// If no player with such name was found, return an empty object.
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		jsonEncode(w, enc, emptyResponse{})
		return
	}

	w.WriteHeader(http.StatusFound)
	jsonEncode(w, enc, response)
}

// GetAbilities is the handler that returns the abilities information of a
// player in the database.
func GetAbilities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)
	playerName := vars["name"]
	if playerName == "" {
		sendErrorResponse(w, enc, invalidPlayerNameError,
			"A player's name should contain only characters and spaces.",
			http.StatusBadRequest,
		)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	playersCollection := playersDatabase.Collection(players)

	findResult := findPlayer(ctx, playersCollection, playerName)
	// Can this even happen?
	if findResult == nil {
		log.Println("Error in FindOne")
		sendErrorResponse(w, enc, databaseError,
			"An error was encountered while accsesing the database.",
			http.StatusInternalServerError,
		)
		return
	}

	var player creature.Creature
	err := findResult.Decode(&player)
	// If no player with such name was found, return an empty object.
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		jsonEncode(w, enc, emptyResponse{})
		return
	}

	w.WriteHeader(http.StatusFound)
	jsonEncode(w, enc, player.Abilities)
}

// validAbility checks if the provided value if a valid ability.
func validAbility(a string) bool {
	switch strings.ToLower(a) {
	case abilities.Strength, abilities.Dexterity, abilities.Constitution,
		abilities.Intelligence, abilities.Wisdom, abilities.Charisma:
		return true
	default:
		return false
	}
}

// SetAbility is the handler that sets the requested ability of a player to
// the provided value.
func SetAbility(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)
	playerName := vars["name"]
	if playerName == "" {
		sendErrorResponse(w, enc, invalidPlayerNameError,
			"A player's name should contain only characters and spaces.",
			http.StatusBadRequest,
		)
		return
	}
	ability := vars["ability"]
	if !validAbility(ability) {
		sendErrorResponse(w, enc,
			"invalid ability name",
			"Please provide a valid ability name.",
			http.StatusBadRequest,
		)
		return
	}
	v := vars["number"]
	value, err := strconv.Atoi(v)
	if err != nil {
		sendErrorResponse(w, enc,
			"invalid ability value",
			"Please provide a valid value for the ability.",
			http.StatusBadRequest,
		)
		return
	}
	if abilities.OutOfRange(value) {
		sendErrorResponse(w, enc,
			"ability value out of range",
			"Please provide a value within range.",
			http.StatusBadRequest,
		)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	playersCollection := playersDatabase.Collection(players)

	findResult := findPlayer(ctx, playersCollection, playerName)
	// Can this even happen?
	if findResult == nil {
		log.Println("Error in FindOne.")
		sendErrorResponse(w, enc, databaseError,
			"An error was encountered while accessing the database.",
			http.StatusInternalServerError,
		)
		return
	} else if findResult.Err() == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound)
		jsonEncode(w, enc, emptyResponse{})
		return
	}

	_, err = playersCollection.UpdateOne(ctx, bson.M{
		"name": playerName,
	}, bson.M{
		"$set": bson.M{
			"abilities." + ability:               value,
			"abilities." + ability + "_modifier": abilities.AbilityScoresAndModifiers[value],
		}})
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, enc, databaseError,
			"There was an error updating the entry in the database.",
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// setCreatureAttribute sets the provided attribute to the provided value.
func setCreatureAttribute(w http.ResponseWriter, r *http.Request, query bson.M) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)
	playerName := vars["name"]
	if playerName == "" {
		sendErrorResponse(w, enc, invalidPlayerNameError,
			"A player's name should contain only characteres and spaces.",
			http.StatusBadRequest,
		)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	playersCollection := playersDatabase.Collection(players)

	findResult := findPlayer(ctx, playersCollection, playerName)
	// Can this even happen?
	if findResult == nil {
		log.Println("Error in FindOne")
		sendErrorResponse(w, enc, databaseError,
			"An error was encountered while accessing the database.",
			http.StatusInternalServerError,
		)
		return
	} else if findResult.Err() == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound)
		jsonEncode(w, enc, emptyResponse{})
		return
	}

	_, err := playersCollection.UpdateOne(ctx, bson.M{
		"name": playerName,
	}, bson.M{
		"$set": query,
	})
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, enc, databaseError,
			"There was an error updating the entry in the database.",
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// SetHitPoints is the handler that sets the hitpoints of the requested creature
// to the provided value.
func SetHitPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v := vars["number"]
	value, err := strconv.Atoi(v)
	if err != nil {
		sendErrorResponse(
			w,
			json.NewEncoder(w),
			"invalid value",
			"Please provide a valid numeric value.",
			http.StatusBadRequest,
		)
	}

	setCreatureAttribute(w, r, bson.M{
		"hitpoints": value,
	})
}

// SetLevel is the handler that sets the hitpoints of the requested creature to
// the provided value.
func SetLevel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v := vars["number"]
	value, err := strconv.Atoi(v)
	if err != nil {
		sendErrorResponse(
			w,
			json.NewEncoder(w),
			"invalid value",
			"Please provide a valid numeric value.",
			http.StatusBadRequest,
		)
	}

	setCreatureAttribute(w, r, bson.M{
		"level":             value,
		"proficiency_bonus": creature.ProficiencyBonusPerLevel[value],
	})
}

// SetArmorClass is the handler that sets the armor class of the requested
// creature to the provided value.
func SetArmorClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v := vars["number"]
	value, err := strconv.Atoi(v)
	if err != nil {
		sendErrorResponse(
			w,
			json.NewEncoder(w),
			"invalid value",
			"Please provide a valid numeric value.",
			http.StatusBadRequest,
		)
	}

	setCreatureAttribute(w, r, bson.M{
		"armor_class": value,
	})
}
