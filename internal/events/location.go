package events

import (
	"github.com/DeveloperAromal/prana-server/internal/events/finalizer"
)

func LogLocation(lat, log float64) {
	finalizer.FinalLocation(log, lat)

}
