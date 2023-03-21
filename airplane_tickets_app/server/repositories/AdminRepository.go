package repositories

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/models"
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
