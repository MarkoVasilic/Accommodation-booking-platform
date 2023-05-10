package domain

import (
	"time"
)

type FindAvailability struct {
	Location  string
	GuestsNum int
	StartDate time.Time
	EndDate   time.Time
}
