package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	Host  Role = "HOST"
	Guest Role = "GUEST"
)

type User struct {
	Id            primitive.ObjectID `bson:"_id" json:"_id"`
	Username      *string            `bson:"username" json:"username" validate:"required,min=4,max=30"`
	FirstName     *string            `bson:"first_name" json:"first_name" validate:"required,min=2,max=30"`
	LastName      *string            `bson:"last_name" json:"last_name"  validate:"required,min=2,max=30"`
	Password      *string            `bson:"password" json:"password"   validate:"required,min=6"`
	Email         *string            `bson:"email" json:"email"      validate:"email,required"`
	Address       *string            `bson:"address" json:"address"      validate:"min=5,max=100"`
	Token         *string            `bson:"token" json:"token"`
	Refresh_Token *string            `bson:"refresh_token" json:"refresh_token"`
	Role          *Role              `bson:"role" json:"role" validate:"required"`
	User_ID       string             `bson:"user_id" json:"user_id"`
}

type Accommodation struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	HostID     primitive.ObjectID `bson:"host_id" json:"host_id" validate:"required"`
	Name       string             `bson:"name" json:"name" validate:"required,min=4,max=30"`
	Location   string             `bson:"location" json:"location" validate:"required,min=4,max=100"`
	Wifi       bool               `bson:"wifi" json:"wifi" validate:"required"`
	Kitchen    bool               `bson:"kitchen" json:"kitchen" validate:"required"`
	AC         bool               `bson:"ac" json:"ac" validate:"required"`
	ParkingLot bool               `bson:"parking_lot" json:"parking_lot" validate:"required"`
	MinGuests  int                `bson:"min_guests" json:"min_guests" validate:"required,min=1"`
	MaxGuests  int                `bson:"max_guests" json:"max_guests" validate:"required,min=1"`
	Images     []*string          `bson:"images" json:"images"`
	AutoAccept bool               `bson:"auto_accept" json:"auto_accept" validate:"required"`
}

type Availability struct {
	ID              primitive.ObjectID `bson:"_id" json:"_id"`
	AccommodationID primitive.ObjectID `bson:"accomodation_id" json:"accomodation_id" validate:"required"`
	StartDate       time.Time          `bson:"start_date" json:"start_date" validate:"required"`
	EndDate         time.Time          `bson:"end_date" json:"end_date" validate:"required,gtfield=StartDate"`
	Price           float64            `bson:"price" json:"price" validate:"required,min=0"`
	IsPricePerGuest bool               `bson:"is_price_per_guest" json:"is_price_per_guest" validate:"required"`
}

type Reservation struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id"`
	AvailabilityID primitive.ObjectID `bson:"availability_id" json:"availability_id" validate:"required"`
	GuestID        primitive.ObjectID `bson:"guest_id" json:"guest_id" validate:"required"`
	StartDate      time.Time          `bson:"start_date" json:"start_date" validate:"required"`
	EndDate        time.Time          `bson:"end_date" json:"end_date" validate:"required,gtfield=StartDate"`
	NumGuests      int                `bson:"num_guests" json:"num_guests" validate:"required,min=1"`
	IsAccepted     bool               `bson:"is_accepted" json:"is_accepted"`
}
