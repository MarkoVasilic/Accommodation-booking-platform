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

func (repo *AccommodationRepository) CreateAccommodation(accommodation *models.Accommodation) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	accommodation.ID = primitive.NewObjectID()
	_, inserterr := repo.AccommodationCollection.InsertOne(ctx, accommodation)
	return inserterr
}

func (repo *AccommodationRepository) CountByName(name string) (int64, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return repo.AccommodationCollection.CountDocuments(ctx, bson.M{"name": name})
}

func (repo *AccommodationRepository) GetAllAccommodationsByLocation(location string) ([]models.Accommodation, error) {
	var accommodations []models.Accommodation
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"location": location}
	cursor, err := repo.AccommodationCollection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &accommodations); err != nil {
		return nil, err
	}

	return accommodations, err
}
func (repo *AccommodationRepository) DeleteAccommodation(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return repo.AccommodationCollection.DeleteOne(ctx, bson.M{"_id": id})
}

func (repo *AccommodationRepository) GetAllAccommodations(hostId primitive.ObjectID) ([]models.Accommodation, error) {
	var accommodations []models.Accommodation
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"host_id": hostId}
	cursor, err := repo.AccommodationCollection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &accommodations); err != nil {
		return nil, err
	}

	return accommodations, err
}

func (repo *AccommodationRepository) GetAll() ([]models.Accommodation, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := repo.AccommodationCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var accommodations []models.Accommodation
	if err = cursor.All(ctx, &accommodations); err != nil {
		return nil, err
	}

	return accommodations, nil
}
