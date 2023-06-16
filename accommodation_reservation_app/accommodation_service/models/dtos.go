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

type AccommodationGradeDetails struct {
	GuestFirstName    string    `bson:"guest_first_name" json:"guest_first_name"`
	GuestLastName     string    `bson:"guest_last_name" json:"guest_last_name"`
	AccommodationName string    `bson:"accommodation_name" json:"accommodation_name"`
	Grade             int       `bson:"grade" json:"grade"`
	DateOfGrade       time.Time `bson:"date_of_grade" json:"date_of_grade"`
}

type FilterAvailability struct {
	Location      string    `bson:"location" json:"location" validate:"required,min=4,max=100"`
	GuestsNum     int32     `bson:"guests_num" json:"guests_num" validate:"required,min=1"`
	StartDate     time.Time `bson:"start_date" json:"start_date" validate:"required"`
	EndDate       time.Time `bson:"end_date" json:"end_date" validate:"required,gtfield=StartDate"`
	GradeMin      int32     `bson:"grade_min" json:"grade_min" validate:"required, min=1, max=5"`
	GradeMax      int32     `bson:"grade" json:"grade" validate:"required, min=1, max=5"`
	Wifi          bool      `bson:"wifi" json:"wifi" validate:"required"`
	Kitchen       bool      `bson:"kitchen" json:"kitchen" validate:"required"`
	AC            bool      `bson:"ac" json:"ac" validate:"required"`
	ParkingLot    bool      `bson:"parking_lot" json:"parking_lot" validate:"required"`
	ProminentHost bool      `bson:"prominent_host" json:"prominent_host" validate:"required"`
}

type AccommodationGradeDetailsDTO struct {
	AccommodationGradeDetails []*AccommodationGradeDetails
	AverageGrade              float64
}
