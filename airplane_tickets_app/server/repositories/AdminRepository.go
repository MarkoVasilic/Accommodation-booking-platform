package repositories

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRepository struct {
	UserCollection   *mongo.Collection
	FlightCollection *mongo.Collection
}

func (repo *AdminRepository) CreateFlight(flight *models.Flight) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	flight.ID = primitive.NewObjectID()
	_, inserterr := repo.FlightCollection.InsertOne(ctx, flight)
	return inserterr
}

func (repo *AdminRepository) GetFlightById(id string) (models.Flight, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var foundflight models.Flight
	objID, err_hex := primitive.ObjectIDFromHex(id)
	if err_hex != nil {
		panic(err_hex)
	}
	err := repo.FlightCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&foundflight)
	defer cancel()
	return foundflight, err
}

func (repo *AdminRepository) DeleteFlightById(id string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	_, err = repo.FlightCollection.DeleteOne(ctx, bson.M{"_id": objID})
	defer cancel()
	return err
}
