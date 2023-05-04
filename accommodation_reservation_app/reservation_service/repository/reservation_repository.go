package repository

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReservationRepository struct {
	ReservationCollection *mongo.Collection
}

// by availability
func (repo *ReservationRepository) GetAllReservations(availibiltyId primitive.ObjectID) ([]models.Reservation, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"availability_id": availibiltyId}
	cursor, err := repo.ReservationCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var reservations []models.Reservation
	if err = cursor.All(ctx, &reservations); err != nil {
		return nil, err
	}

	return reservations, nil
}
