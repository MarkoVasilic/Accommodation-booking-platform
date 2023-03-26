package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OnlyEmail struct {
	Email *string `json:"email"      validate:"email,required"`
}

type SearchedFlights struct {
	ID                primitive.ObjectID `bson:"_id" form:"_id"`
	Name              *string            `bson:"name" form:"name"`
	Taking_Off_Date   time.Time          `bson:"taking_off_date" form:"taking_off_date" validate:"required"`
	Start_Location    *string            `bson:"start_location" form:"start_location"  validate:"required"`
	End_Location      *string            `bson:"end_location" form:"end_location"   validate:"required,min=3"`
	Price             *float64           `bson:"price" form:"price" validate:"required"`
	Number_Of_Tickets *uint64            `bson:"number_of_tickets" form:"number_of_tickets" validate:"required, min=1"`
	Total_Price       *float64           `bson:"total_price" form:"total_price"`
}
