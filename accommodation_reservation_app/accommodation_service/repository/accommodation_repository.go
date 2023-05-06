package repository

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccommodationRepository struct {
	AccommodationCollection *mongo.Collection
}

func (repo *AccommodationRepository) GetAccommodationById(id primitive.ObjectID) (models.Accommodation, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	var accommodation models.Accommodation
	err := repo.AccommodationCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&accommodation)
	defer cancel()
	return accommodation, err
}
