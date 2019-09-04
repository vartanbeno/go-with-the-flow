package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Env is basically the handler we pass to the router.
type Env struct {
	Flow *Flow
}

// RequestToAdd holds the fields necessary to create a new Request struct to add to the Flow.
type RequestToAdd struct {
	ID           string `json:"id"`
	Message      string `json:"message"`
	Milliseconds uint   `json:"ms"`
}

func (env *Env) addRequest(w http.ResponseWriter, r *http.Request) {
	var req RequestToAdd
	json.NewDecoder(r.Body).Decode(&req)

	pollingInterval := time.Duration(time.Millisecond * time.Duration(req.Milliseconds))

	request := NewRequest(req.ID, req.Message, pollingInterval)
	env.Flow.AddRequest(&request)
}

func createRouter(flow *Flow) *mux.Router {
	env := Env{Flow: flow}

	r := mux.NewRouter()
	r.HandleFunc("/api", env.addRequest).Methods("POST")

	return r
}
