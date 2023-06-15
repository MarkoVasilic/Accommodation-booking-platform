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

type FilterAvailability struct {
	Location      string
	GuestsNum     int
	StartDate     time.Time
	EndDate       time.Time
	GradeMin      int
	GradeMax      int
	Wifi          bool
	Kitchen       bool
	AC            bool
	ParkingLot    bool
	ProminentHost bool
}
