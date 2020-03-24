package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SimplyVC/oasis_api_server/src/config"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
	"github.com/SimplyVC/oasis_api_server/src/responses"
)

// Pong responds with a ping if entire API is online
func Pong(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")
	lgr.Info.Println("Received request for /api/pingapi")
	json.NewEncoder(w).Encode(responses.SuccessResponsed)
}

// GetConnections retrieves all possible connections that have been loaded in configuration file
func GetConnections(w http.ResponseWriter, r *http.Request) {

	// Add header so that received knows they're receiving JSON
	w.Header().Add("Content-Type", "application/json")
	lgr.Info.Println("Received request for /api/getconnectionslist")

	// Create new empty Slice of strings where connections will be stored
	connectionsResponse := []string{}
	allSockets := config.GetSockets()

	lgr.Info.Println("Iterating through all socket connections.")
	for _, socket := range allSockets {
		lgr.Info.Println("Node : ", socket["node_name"], " has socket at : ", socket["ws_url"])
		connectionsResponse = append(connectionsResponse, socket["ws_url"])
	}
	// Encode object and send it using a predefind response
	json.NewEncoder(w).Encode(responses.ConnectionsResponse{Results: connectionsResponse})
}
