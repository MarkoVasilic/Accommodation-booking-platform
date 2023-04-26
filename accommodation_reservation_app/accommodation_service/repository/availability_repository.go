package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AvailabilityRepository struct {
	AvailabilityCollection *mongo.Collection
}
