package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id"`
	AvailabilityID primitive.ObjectID `bson:"availability_id" json:"availability_id" validate:"required"`
	GuestID        primitive.ObjectID `bson:"guest_id" json:"guest_id" validate:"required"`
	StartDate      time.Time          `bson:"start_date" json:"start_date" validate:"required"`
	EndDate        time.Time          `bson:"end_date" json:"end_date" validate:"required,gtfield=StartDate"`
	NumGuests      int                `bson:"num_guests" json:"num_guests" validate:"required,min=1"`
	IsAccepted     bool               `bson:"is_accepted" json:"is_accepted"`
	IsCanceled     bool               `bson:"is_canceled" json:"is_canceled"`
	IsDeleted      bool               `bson:"is_deleted" json:"is_deleted"`
}
