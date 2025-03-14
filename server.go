package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DeveloperAromal/prana-server/internal/events"
)

// Location represents latitude and longitude
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// ReqData represents the structure of the received JSON request
type ReqData struct {
	Location     Location `json:"location"`
	Timestamp    string   `json:"time"`
	AccidentType string   `json:"accident_type"`
}

// handler processes incoming HTTP POST requests with accident data
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data ReqData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Dispatch data to events functions
	events.PrintLocation(data.Location.Lat, data.Location.Lng)
	events.PrintTimestamp(data.Timestamp)
	events.PrintAccidentType(data.AccidentType)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data received successfully"))
}

func main() {
	http.HandleFunc("/receive", handler)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
