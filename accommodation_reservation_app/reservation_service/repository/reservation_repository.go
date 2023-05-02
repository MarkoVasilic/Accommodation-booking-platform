package repository

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/models"
)

type ReservationRepository struct {
	ReservationCollection *mongo.Collection
}


func (repo *ReservationRepository) GetAllReservations() ([]models.Reservation, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := repo.ReservationCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var reservations []models.Reservation
	if err = cursor.All(ctx, &reservations); err != nil {
		return nil, err
	}

	return reservations, nil
}
