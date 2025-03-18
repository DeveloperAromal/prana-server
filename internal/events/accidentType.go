package events

import (
	"github.com/DeveloperAromal/prana-server/internal/events/finalizer"
)

func LogAccidentType(accidentType string) {
	finalizer.FinalAccidentType(accidentType)
}
