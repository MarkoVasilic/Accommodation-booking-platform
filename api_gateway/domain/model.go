package domain

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
