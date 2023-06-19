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

type UserGradeDetails struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id"`
	GuestFirstName string             `bson:"guest_first_name" json:"guest_first_name"`
	GuestLastName  string             `bson:"guest_last_name" json:"guest_last_name"`
	HostFirstName  string             `bson:"host_first_name" json:"host_first_name"`
	HostLastName   string             `bson:"host_last_name" json:"host_last_name"`
	Grade          int                `bson:"grade" json:"grade"`
	DateOfGrade    time.Time          `bson:"date_of_grade" json:"date_of_grade"`
}

type HostDetails struct {
	Id        primitive.ObjectID `bson:"id" json:"id"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
}

type UserGradeDetailsDTO struct {
	UserGradeDetails []*UserGradeDetails
	AverageGrade     float64
}
