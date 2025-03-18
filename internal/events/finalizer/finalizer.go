package finalizer

import (
	"fmt"
)

func FinalLocation(lat, log float64) {
	fmt.Printf("Location: Latitude = %f, Longitude = %f\n", lat, log)
}

func FinalTime(timestamp string) {
	fmt.Printf("Timestamp: %s\n", timestamp)
}

func FinalAccidentType(accidentType string) {
	fmt.Printf("Accident Type: %s\n", accidentType)
}
