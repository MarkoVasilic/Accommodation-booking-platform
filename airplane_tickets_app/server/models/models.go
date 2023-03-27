package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	Regular Role = "REGULAR"
	Admin   Role = "ADMIN"
)

type User struct {
	ID            primitive.ObjectID   `bson:"_id" json:"_id"`
	First_Name    *string              `bson:"first_name" json:"first_name" validate:"required,min=2,max=30"`
	Last_Name     *string              `bson:"last_name" json:"last_name"  validate:"required,min=2,max=30"`
	Password      *string              `bson:"password" json:"password"   validate:"required,min=6"`
	Email         *string              `bson:"email" json:"email"      validate:"email,required"`
	Token         *string              `bson:"token" json:"token"`
	Refresh_Token *string              `bson:"refresh_token" json:"refresh_token"`
	Role          *Role                `bson:"role" json:"role" validate:"required"`
	User_ID       string               `bson:"user_id" json:"user_id"`
	Created_At    time.Time            `bson:"created_at" json:"created_at"`
	Updated_At    time.Time            `bson:"updated_at" json:"updated_at"`
	UserTickets   []primitive.ObjectID `bson:"user_tickets" json:"user_tickets"`
}

type Flight struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id"`
	Name              *string            `bson:"name" json:"name"`
	Taking_Off_Date   time.Time          `bson:"taking_off_date" json:"taking_off_date" validate:"required"`
	Start_Location    *string            `bson:"start_location" json:"start_location"  validate:"required"`
	End_Location      *string            `bson:"end_location" json:"end_location"   validate:"required,min=3"`
	Price             *float64           `bson:"price" json:"price" validate:"required,gte=0"`
	Number_Of_Tickets *uint64            `bson:"number_of_tickets" json:"number_of_tickets" validate:"required,min=1,gte=0"`
}

type Ticket struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id"`
	Flight primitive.ObjectID `bson:"flight" json:"flight"`
}
