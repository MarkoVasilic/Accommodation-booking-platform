package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ReservationRepository struct {
	ReservationCollection *mongo.Collection
}
