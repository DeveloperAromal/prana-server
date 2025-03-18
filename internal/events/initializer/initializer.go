package initializer

import "github.com/DeveloperAromal/prana-server/internal/events"

func PrintData(lat float64, log float64, Timestamp string, AccidentType string) {
	events.LogLocation(log, lat)
	events.Logtime(Timestamp)
	events.LogAccidentType(AccidentType)
}
