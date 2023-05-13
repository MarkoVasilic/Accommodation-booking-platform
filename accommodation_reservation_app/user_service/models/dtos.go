package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FindReservation struct {
	ReservationId    primitive.ObjectID `bson:"reservation_id" json:"reservation_id"`
	GuestID          primitive.ObjectID `bson:"guest_id" json:"guest_id" validate:"required"`
	Name             string             `bson:"name" json:"name" validate:"required,min=4,max=30"`
	Location         string             `bson:"location" json:"location" validate:"required,min=4,max=100"`
	StartDate        time.Time          `bson:"start_date" json:"start_date" validate:"required"`
	EndDate          time.Time          `bson:"end_date" json:"end_date" validate:"required,gtfield=StartDate"`
	NumOfCancelation int                `bson:"num_cancelation" json:"num_cancelation" validate:"required"`
	IsAccepted       bool               `bson:"is_accepted" json:"is_accepted"`
	IsCanceled       bool               `bson:"is_canceled" json:"is_canceled"`
}
