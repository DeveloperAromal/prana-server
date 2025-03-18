package events

import (
	"github.com/DeveloperAromal/prana-server/internal/events/finalizer"
)

func Logtime(timestamp string) {
	finalizer.FinalTime(timestamp)
}
