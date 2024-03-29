package repository

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AvailabilityRepository struct {
	AvailabilityCollection *mongo.Collection
}

func (repo *AvailabilityRepository) CreateAvailability(availability *models.Availability) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	availability.ID = primitive.NewObjectID()
	_, inserterr := repo.AvailabilityCollection.InsertOne(ctx, availability)
	return inserterr
}

func (repo *AvailabilityRepository) GetAllAvailabilitiesByAccommodationID(accommodationID primitive.ObjectID) ([]models.Availability, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"accomodation_id": accommodationID}
	cursor, err := repo.AvailabilityCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var availabilities []models.Availability
	if err = cursor.All(ctx, &availabilities); err != nil {
		return nil, err
	}

	return availabilities, nil
}

func (repo *AvailabilityRepository) UpdateAvailability(avail *models.Availability) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": avail.ID}
	update := bson.M{
		"$set": bson.M{
			"price":              avail.Price,
			"is_price_per_guest": avail.IsPricePerGuest,
		},
	}
	options := options.Update().SetUpsert(false)
	_, updateErr := repo.AvailabilityCollection.UpdateOne(ctx, filter, update, options)
	return updateErr
}

func (repo *AvailabilityRepository) GetAvailabilityById(id primitive.ObjectID) (models.Availability, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	var availability models.Availability
	err := repo.AvailabilityCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&availability)
	defer cancel()
	return availability, err
}

func (repo *AvailabilityRepository) GetAllAvailabilities() ([]models.Availability, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := repo.AvailabilityCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var availabilities []models.Availability
	if err = cursor.All(ctx, &availabilities); err != nil {
		return nil, err
	}

	return availabilities, nil
}

func (repo *AvailabilityRepository) DeleteAvailability(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return repo.AvailabilityCollection.DeleteOne(ctx, bson.M{"_id": id})
}
