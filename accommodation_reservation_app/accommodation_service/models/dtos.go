package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FindAvailability struct {
	AccommodationId primitive.ObjectID `bson:"accommodation_id" json:"accommodation_id"`
	AvailabilityID  primitive.ObjectID `bson:"availability_id" json:"availability_id"`
	HostID          primitive.ObjectID `bson:"host_id" json:"host_id" validate:"required"`
	Name            string             `bson:"name" json:"name" validate:"required,min=4,max=30"`
	Location        string             `bson:"location" json:"location" validate:"required,min=4,max=100"`
	Wifi            bool               `bson:"wifi" json:"wifi" validate:"required"`
	Kitchen         bool               `bson:"kitchen" json:"kitchen" validate:"required"`
	AC              bool               `bson:"ac" json:"ac" validate:"required"`
	ParkingLot      bool               `bson:"parking_lot" json:"parking_lot" validate:"required"`
	Images          []*string          `bson:"images" json:"images"`
	StartDate       time.Time          `bson:"start_date" json:"start_date" validate:"required"`
	EndDate         time.Time          `bson:"end_date" json:"end_date" validate:"required,gtfield=StartDate"`
	TotalPrice      float64            `bson:"total_price" json:"total_price" validate:"required,min=0"`
	SinglePrice     float64            `bson:"single_price" json:"single_price" validate:"required,min=0"`
	IsPricePerGuest bool               `bson:"is_price_per_guest" json:"is_price_per_guest" validate:"required"`
}
