package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/aakordas/creature_manager/pkg/dice"
	"github.com/gorilla/mux"
)

// TODO: Make the responses get formed in a go routine, instead of everything
// staying in the main server.

// TODO: Testing for the functions that accept ResponseWriters and Requests.

// diceRoutes properly initializes the routes for the dice part of
// the server.
func diceRoutes(r *mux.Router) *mux.Router {
	var (
		sides  = "{sides:[0-9]+}"
		dsides = "{sides:[d|D][0-9]+}"
		count  = "{count:[0-9]+}"
	)

	api := r.PathPrefix("/api/v1/").Subrouter()

	// Rolls
	api.HandleFunc("/roll", Roll)
	api.Queries("sides", sides, "count", count).HandlerFunc(Roll).Methods(http.MethodGet)

	roll := api.PathPrefix("/roll/").Subrouter()
	roll.HandleFunc("/"+dsides, RollN).Methods(http.MethodGet)

	dRoll := roll.PathPrefix("/" + dsides + "/").Subrouter()
	dRoll.HandleFunc("/"+count, DRollN).Methods(http.MethodGet)
	dRoll.Queries("count", count).HandlerFunc(RollN)

	return r
}

type rollResponse struct {
	Count  int `json:"count" bson:"count"`   // The number of dice that got rolled.
	Sides  int `json:"sides" bson:"sides"`   // The number of sides each dice had.
	Result int `json:"result" bson:"result"` // The result of the rolling.
}

type errorResponse struct {
	Error        string `json:"error" bson:"error"`
	ErrorMessage string `json:"error_message" bson:"error_message"`
}

// writeHeader writes the header of a valid response.
// func writeHeader(w http.ResponseWriter) {
// w.Header().Set("Content-Type", "application/json")
// w.WriteHeader(http.StatusOK)
// }

// jsonEncode wraps the json.Encoder.Encode and error checking.
func jsonEncode(w http.ResponseWriter, enc *json.Encoder, v interface{}) {
	if err := enc.Encode(v); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type invalidDice struct {
	error
}

// rollDice rolls the specified number of the specified dice.
func rollDice(sides, count int) (result int, err error) {
	d := chooseDice(sides)
	if d == nil {
		return 0, invalidDice{}
	}

	for i := 0; i < count; i++ {
		result += d()
	}

	return result, nil
}

// getSides gets the integer value from the sides from the passed string.
func getSides(sides string) int {
	if sides == "" {
		return 20
	}

	s, err := strconv.Atoi(sides)
	if err != nil {
		return 0
	}

	return s
}

// getCount gets the integer value from the count from the passed string.
func getCount(count string) int {
	if count == "" {
		return 1
	}

	c, err := strconv.Atoi(count)
	if err != nil {
		return 0
	}

	return c
}

// Roll is the handler for all the requested rolls of one die.
func Roll(w http.ResponseWriter, r *http.Request) {
	sides := r.FormValue("sides")
	count := r.FormValue("count")

	s := getSides(sides)
	if s == 0 {
		errResponse := errorResponse{"invalid sides", "The dice requested is not available."}
		w.WriteHeader(http.StatusNotAcceptable)
		enc := json.NewEncoder(w)
		jsonEncode(w, enc, errResponse)
		return
	}

	c := getCount(count)
	if c == 0 {
		errResponse := errorResponse{"invalid count", "The number of dice requested is invalid."}
		w.WriteHeader(http.StatusNotAcceptable)
		enc := json.NewEncoder(w)
		jsonEncode(w, enc, errResponse)
		return
	}

	response(w, s, c)
}

// chooseDice chooses the appropriate dice given a number of sides and returns
// it and the number of sides as an integer..
func chooseDice(sides int) dice.Dice {
	switch sides {
	case 4:
		return dice.D4
	case 6:
		return dice.D6
	case 8:
		return dice.D8
	case 10:
		return dice.D10
	case 12:
		return dice.D12
	case 20:
		return dice.D20
	case 100:
		return dice.D100
	default:
		return nil
	}
}

// response deals with the response part of the HTTP response, whether that is an error response or not.
func response(w http.ResponseWriter, s, c int) {
	w.Header().Set("Content-Type", "application/json")

	result, err := rollDice(s, c)
	if err != nil {
		errResponse := errorResponse{"invalid sides", "The dice requested is not available."}
		w.WriteHeader(http.StatusNotAcceptable)
		enc := json.NewEncoder(w)
		jsonEncode(w, enc, errResponse)
		return
	}

	response := rollResponse{c, s, result}
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	jsonEncode(w, enc, response)
}

// RollN is the handler for all the requested rolls of n d dice, where n is
// specified with a query.
func RollN(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sides := vars["sides"]
	count := r.FormValue("count")

	s := getSides(sides[1:])
	if s == 0 {
		errResponse := errorResponse{"invalid sides", "The dice requested is not available."}
		w.WriteHeader(http.StatusNotAcceptable)
		enc := json.NewEncoder(w)
		jsonEncode(w, enc, errResponse)
		return
	}

	c := getCount(count)
	if c == 0 {
		errResponse := errorResponse{"invalid count", "The number of dice requested is invalid."}
		w.WriteHeader(http.StatusNotAcceptable)
		enc := json.NewEncoder(w)
		jsonEncode(w, enc, errResponse)
		return
	}

	response(w, s, c)
}

// DRollN is the handlre for all the requested rolls of n d dice, where n is
// specified with a variable.
func DRollN(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sides := vars["sides"]
	count := vars["count"]

	s := getSides(sides)
	if s == 0 {
		errResponse := errorResponse{"invalid sides", "The dice requested is not available."}
		w.WriteHeader(http.StatusNotAcceptable)
		enc := json.NewEncoder(w)
		jsonEncode(w, enc, errResponse)
		return
	}

	c := getCount(count)
	if c == 0 {
		errResponse := errorResponse{"invalid count", "The number of dice requested is invalid."}
		w.WriteHeader(http.StatusNotAcceptable)
		enc := json.NewEncoder(w)
		jsonEncode(w, enc, errResponse)
	}

	response(w, s, c)
}
