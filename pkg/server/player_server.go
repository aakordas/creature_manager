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
	"github.com/aakordas/creature_manager/pkg/saves"
	"github.com/aakordas/creature_manager/pkg/skills"
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

	// TODO: context.WithClose. Return the close function and pass it in Disconnect.
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
	// w.WriteHeader(h)
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

func getInfo(w http.ResponseWriter, r *http.Request, v interface{}) {
	enc := json.NewEncoder(w)

	player, err := getPlayer(w, r)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusFound)

	var res interface{}
	switch v.(type) {
	case creature.Creature:
		res = player
	case abilities.Abilities:
		res = player.Abilities
	case skills.Skills:
		res = player.Skills
	case saves.SavingThrows:
		res = player.SavingThrows
	}

	if res == nil {
		jsonEncode(w, enc, emptyResponse{})
		return
	}

	jsonEncode(w, enc, res)
}

// GetPlayer is the handler that returns the information about a player in the
// database.
func GetPlayer(w http.ResponseWriter, r *http.Request) {
	var p creature.Creature
	getInfo(w, r, p)
}

type missingPlayerError struct {
	Name string
}

// Error implements the Error interface for missingPlayerError.
func (e missingPlayerError) Error() string {
	return "No player found with name " + e.Name
}

// getPlayer returns a creature.Creature.
func getPlayer(w http.ResponseWriter, r *http.Request) (*creature.Creature, error) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)
	playerName := vars["name"]
	if playerName == "" {
		sendErrorResponse(w, enc, invalidPlayerNameError,
			"A player's name should contain only characters and spaces.",
			http.StatusBadRequest,
		)
		return nil, missingPlayerError{playerName}
	}

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	playersCollection := playersDatabase.Collection(players)

	findResult := findPlayer(ctx, playersCollection, playerName)
	// Can this even happen?
	if findResult == nil {
		log.Println(findResult.Err())
		sendErrorResponse(w, enc, databaseError,
			"An error was encountered while accessing the database.",
			http.StatusInternalServerError,
		)
		return nil, missingPlayerError{playerName}
	}

	var player creature.Creature
	err := findResult.Decode(&player)
	// If no player was found, return an empty object.
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		jsonEncode(w, enc, emptyResponse{})
		return nil, missingPlayerError{playerName}
	}

	return &player, nil
}

// GetAbilities is the handler that returns the abilities information of a
// player in the database.
func GetAbilities(w http.ResponseWriter, r *http.Request) {
	var a abilities.Abilities
	getInfo(w, r, a)
}

// validAbility checks if the provided value is a valid ability.
func validAbility(a string) bool {
	switch strings.ToLower(a) {
	case abilities.Strength, abilities.Dexterity, abilities.Constitution,
		abilities.Intelligence, abilities.Wisdom, abilities.Charisma:
		return true
	default:
		return false
	}
}

// calculatePassivePerception returns the passive perception of the creature,
// based on its wisdom and proficiency modifiers
func calculatePassivePerception(c creature.Creature, w int, pb int) int {
	modifier := abilities.AbilityScoresAndModifiers[w]

	perceptionProficiency := c.Skills[skills.Perception]
	if perceptionProficiency != nil {
		return 10 + modifier + pb
	}

	return 10 + modifier
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

	player, err := getPlayer(w, r)
	if err != nil {
		sendErrorResponse(w, enc,
			"Could not fetch the player.",
			"There was an error retrieving the player from the database.",
			http.StatusBadRequest,
		)
		return
	}

	modifier := abilities.AbilityScoresAndModifiers[value]

	query := bson.M{
		"abilities." + ability:               value,
		"abilities." + ability + "_modifier": modifier,
	}

	if ability == abilities.Wisdom {
		query["passive_perception"] = calculatePassivePerception(*player, value, player.Level)
	}

	// Changes the modifier of a skill, if it depends on this ability.
	for skill := range player.Skills {
		_, err = playersCollection.UpdateMany(ctx, bson.M{
			"name":                          playerName,
			"skills." + skill + ".modifier": ability,
		}, bson.M{
			"$set": bson.M{
				"skills." + skill + ".value": modifier + player.ProficiencyBonus,
			}})
		if err != nil {
			log.Println(err)
			sendErrorResponse(w, enc, databaseError,
				"There was an error updating the entry in the database.",
				http.StatusInternalServerError,
			)
			return
		}
	}

	_, err = playersCollection.UpdateOne(ctx, bson.M{
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

// setNoUpsert sets the provided attribute to the provided value.
func setNoUpsert(w http.ResponseWriter, r *http.Request, enc *json.Encoder, filter, update interface{}) {
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	playersCollection := playersDatabase.Collection(players)

	_, err := playersCollection.UpdateOne(ctx, filter, update)
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
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)
	v := vars["number"]
	value, err := strconv.Atoi(v)
	if err != nil {
		sendErrorResponse(w, enc,
			"invalid value",
			"Please provide a valid numeric value.",
			http.StatusBadRequest,
		)
	}
	playerName := vars["name"]
	if playerName == "" {
		sendErrorResponse(w, enc, invalidPlayerNameError,
			"A player's name should contain only characters and spaces.",
			http.StatusBadRequest,
		)
		return
	}

	f := bson.M{
		"name": playerName,
	}
	u := bson.M{
		"$set": bson.M{
			"hit_points": value,
		}}

	setNoUpsert(w, r, enc, f, u)
}

// SetLevel is the handler that sets the hitpoints of the requested creature to
// the provided value.
func SetLevel(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)
	v := vars["number"]
	value, err := strconv.Atoi(v)
	if err != nil {
		sendErrorResponse(w, enc,
			"invalid value",
			"Please provide a valid numeric value.",
			http.StatusBadRequest,
		)
	}
	playerName := vars["name"]
	if playerName == "" {
		sendErrorResponse(w, enc, invalidPlayerNameError,
			"A player's name should contain only characters and spaces.",
			http.StatusBadRequest,
		)
		return
	}

	player, err := getPlayer(w, r)
	if err != nil {
		sendErrorResponse(w, enc,
			"Could not fetch the player.",
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	proficiencyBonus := creature.ProficiencyBonusPerLevel[value]

	f := bson.M{
		"name": playerName,
	}
	u := bson.M{
		"$set": bson.M{
			"level":             value,
			"proficiency_bonus": proficiencyBonus,
			"passive_perception": calculatePassivePerception(
				*player,
				player.Abilities.Wisdom,
				proficiencyBonus,
			),
		}}

	setNoUpsert(w, r, enc, f, u)
}

// SetArmorClass is the handler that sets the armor class of the requested
// creature to the provided value.
func SetArmorClass(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)

	vars := mux.Vars(r)
	v := vars["number"]
	value, err := strconv.Atoi(v)
	if err != nil {
		sendErrorResponse(w, enc,
			"invalid value",
			"Please provide a valid numeric value.",
			http.StatusBadRequest,
		)
	}
	playerName := vars["name"]
	if playerName == "" {
		sendErrorResponse(w, enc, invalidPlayerNameError,
			"A player's name should contain only characters and spaces.",
			http.StatusBadRequest,
		)
		return
	}

	f := bson.M{
		"name": playerName,
	}
	u := bson.M{
		"$set": bson.M{
			"armor_class": value,
		}}

	setNoUpsert(w, r, enc, f, u)
}

// GetSkills is the handler that returns the skills information of a player in
// the database.
func GetSkills(w http.ResponseWriter, r *http.Request) {
	var s skills.Skills
	getInfo(w, r, s)
}

// validSkill checks if the provided value is a valid skill.
func validSkill(s string) bool {
	switch strings.ToLower(s) {
	case skills.Acrobatics, skills.AnimalHandling, skills.Arcana,
		skills.Athletics, skills.Deception, skills.History,
		skills.Insight, skills.Intimidation, skills.Investigation,
		skills.Medicine, skills.Nature, skills.Perception,
		skills.Performance, skills.Persuasion, skills.Religion,
		skills.SleightOfHand, skills.Stealth, skills.Survival:
		return true
	default:
		return false
	}
}

func setUpsert(w http.ResponseWriter, r *http.Request, enc *json.Encoder, filter, update interface{}) {
	w.Header().Set("Content-Type", "application/json")

	playersCollection := playersDatabase.Collection(players)

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	_, err := playersCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, enc, databaseError,
			"There was an error updating the entry in the database.",
			http.StatusInternalServerError,
		)
		return
	}
}

// SetSkill is the handler that sets the requested skill of a player to the
// provided value.
func SetSkill(w http.ResponseWriter, r *http.Request) {
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
	skill := vars["skill"]
	if !validSkill(skill) {
		sendErrorResponse(w, enc,
			"invalid skill name",
			"Please provide a valid skill name.",
			http.StatusBadRequest,
		)
		return
	}

	player, err := getPlayer(w, r)
	if err != nil {
		sendErrorResponse(w, enc,
			"Could not fetch the player.",
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	abilityModifier := skills.SkillToAbility[skill]

	var modifier int
	switch abilityModifier {
	case abilities.Strength:
		modifier = player.StrengthModifier
	case abilities.Dexterity:
		modifier = player.DexterityModifier
	case abilities.Constitution:
		modifier = player.ConstitutionModifier
	case abilities.Intelligence:
		modifier = player.IntelligenceModifier
	case abilities.Wisdom:
		modifier = player.WisdomModifier
	case abilities.Charisma:
		modifier = player.CharismaModifier
	}

	f := bson.M{
		"name": playerName,
	}
	u := bson.M{
		"$set": bson.M{
			"skills." + skill: bson.M{
				"value":    modifier + player.ProficiencyBonus,
				"modifier": abilityModifier,
			}}}

	setUpsert(w, r, enc, f, u)
}

// GetSaves is the handler that returns the saving throws information of a
// player in the database.
func GetSaves(w http.ResponseWriter, r *http.Request) {
	var s saves.SavingThrows
	getInfo(w, r, s)
}

// validSave checks if the provided value is a valid saving throw.
func validSave(s string) bool {
	switch strings.ToLower(s) {
	case saves.Strength, saves.Dexterity, saves.Constitution,
		saves.Intelligence, saves.Wisdom, saves.Charisma:
		return true
	default:
		return false
	}
}

// SetSave is the handler that sets the requested saving throw of a player in
// the database.
func SetSave(w http.ResponseWriter, r *http.Request) {
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
	save := vars["save"]
	if !validSave(save) {
		sendErrorResponse(w, enc,
			"invalid saving throw name",
			"Please provide a valid saving throw name.",
			http.StatusBadRequest,
		)
		return
	}

	f := bson.M{
		"name": playerName,
	}
	u := bson.M{
		"$set": bson.M{
			"saving_throws." + save: true,
		}}

	setUpsert(w, r, enc, f, u)
}
