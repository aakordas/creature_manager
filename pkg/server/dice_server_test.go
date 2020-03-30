package server

import (
	"bytes"
	"net/http"
	"os"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/gorilla/mux"
)

var router *mux.Router

func diceRouter() {
	router = mux.NewRouter()
	router = diceRoutes(router)
	router = playerRoutes(router)
}

// TestRoll tests the Roll handler.
func TestRoll(t *testing.T) {
	type response struct {
		Code int
		Body string
	}
	tests := []struct {
		name string
		args string
		want response
	}{
		{"Default roll", "", response{http.StatusOK, ``}},
		{"Valid query for sides", "?sides=4", response{http.StatusOK, `"sides":4`}},
		{"Invalid query for sides", "?sides=5", response{http.StatusNotAcceptable, `"invalid sides"`}},
		{"Valid query for count", "?count=2", response{http.StatusOK, `"count":2`}},
		{"Invalid query for count", "?count=0", response{http.StatusNotAcceptable, `"invalid count"`}},
		{"Valid query for sides, invalid for count", "?sides=4&count=0", response{http.StatusNotAcceptable, `"invalid count"`}},
		{"Valid query for count, invalid for sides", "?count=2&sides=1", response{http.StatusNotAcceptable, `"invalid sides"`}},
		{"Valid query for sides and count", "?sides=4&count=2", response{http.StatusOK, `"count":2,"sides":4`}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gofight.New()

			r.GET("/api/v1/roll"+tt.args).
				Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
					if r.Code != tt.want.Code {
						t.Errorf("Handler returned wrong status code: got %v want %v", r.Code, tt.want.Code)
					}

					if !bytes.Contains(r.Body.Bytes(), []byte(tt.want.Body)) {
						t.Errorf("Unexpected body returned.\ngot %v\nwant %v", r.Body, tt.want.Body)
					}
				})
		})
	}
}

// TestRollN tests the RollN handler.
func TestRollN(t *testing.T) {
	type response struct {
		Code int
		Body string
	}
	tests := []struct {
		name string
		args string
		want response
	}{
		{"Valid roll", "/d4", response{http.StatusOK, `"sides":4`}},
		{"Valid roll", "/D4", response{http.StatusOK, `"sides":4`}},
		{"Invalid variable", "/d5", response{http.StatusNotAcceptable, `"error"`}},
		{"Invalid variable", "/D5", response{http.StatusNotAcceptable, `"error"`}},
		{"Valid query for count", "/d4?count=2", response{http.StatusOK, `"count":2,"sides":4`}},
		{"Valid query for count", "/D4?count=2", response{http.StatusOK, `"count":2,"sides":4`}},
		{"Invalid query for count", "/d4?count=0", response{http.StatusNotAcceptable, `"error"`}},
		{"Invalid query for count", "/D4?count=0", response{http.StatusNotAcceptable, `"error"`}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gofight.New()

			r.GET("/api/v1/roll"+tt.args).
				Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
					if r.Code != tt.want.Code {
						t.Errorf("Handler returned wrong status code: got %v want %v", r.Code, tt.want.Code)
					}

					if !bytes.Contains(r.Body.Bytes(), []byte(tt.want.Body)) {
						t.Errorf("Unexpected body returned.\ngot %v\nwant %v", r.Body, tt.want.Body)
					}
				})
		})
	}
}

// TestMain TODO: NEEDS COMMENT INFO
func TestMain(m *testing.M) {
	var (
		serverPort    = ":8080"
		serverAddress = "http://127.0.0.1" + serverPort
	)
	_ = serverAddress

	diceRouter()

	os.Exit(m.Run())
}
