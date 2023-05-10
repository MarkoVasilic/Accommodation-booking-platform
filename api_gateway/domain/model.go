package domain

import "time"

type User struct {
	Username  string
	FirstName string
	LastName  string
	Password  string
	Email     string
	Address   string
	Role      string
}

type Accommodation struct {
	Id         string
	HostId     string
	Name       string
	Location   string
	Wifi       bool
	Kitchen    bool
	AC         bool
	ParkingLot bool
	MinGuests  int32
	MaxGuests  int32
	Images     []string
	AutoAccept bool
}

type Availability struct {
	Id              string
	AccommodationId string
	StartDate       time.Time
	EndDate         time.Time
	Price           float64
	IsPricePerGuest bool
}

type Reservation struct {
	Id             string
	AvailabilityID string
	GuestId        string
	StartDate      time.Time
	EndDate        time.Time
	NumGuests      int
	IsAccepted     bool
	IsCanceled     bool
	IsDeleted      bool
}
