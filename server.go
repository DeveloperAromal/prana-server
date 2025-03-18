package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DeveloperAromal/prana-server/internal/events/initializer"
	"github.com/nats-io/nats.go"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"` // Changed from Log to Lng for consistency
}

type ReqData struct {
	Location     Location `json:"location"`
	Timestamp    string   `json:"time"`
	AccidentType string   `json:"accident_type"`
}

var nc *nats.Conn

func initNATS() {
	var err error
	nc, err = nats.Connect(nats.DefaultURL) // Default URL is "nats://127.0.0.1:4222"
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	fmt.Println("Connected to NATS")
}

func handler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now() // Start time tracking

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data ReqData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	initializer.PrintData(data.Location.Lat, data.Location.Lng, data.Timestamp, data.AccidentType)

	// Publish data to NATS
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	if err := nc.Publish("accidents", jsonData); err != nil {
		http.Error(w, "Failed to publish data to NATS", http.StatusInternalServerError)
		return
	}

	elapsedTime := time.Since(startTime)

	message := fmt.Sprintf(
		"Data received and published successfully | Time Taken: %.9f s, %d ns, %d ms, %d µs",
		elapsedTime.Seconds(),
		elapsedTime.Nanoseconds(),
		elapsedTime.Milliseconds(),
		elapsedTime.Microseconds(),
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": message,
	})
}

func main() {
	initNATS()
	defer nc.Close()

	http.HandleFunc("/receive", handler)
	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
