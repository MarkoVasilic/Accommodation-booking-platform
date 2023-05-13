package repository

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReservationRepository struct {
	ReservationCollection *mongo.Collection
}

// by availability
func (repo *ReservationRepository) GetAllReservationsByAvailability(availibiltyId primitive.ObjectID) ([]models.Reservation, error) {
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

func (repo *ReservationRepository) CreateReservation(reservation *models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservation.ID = primitive.NewObjectID()
	_, err := repo.ReservationCollection.InsertOne(ctx, reservation)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ReservationRepository) GetReservationById(id string) (models.Reservation, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	var foundReservation models.Reservation

	objID, err_hex := primitive.ObjectIDFromHex(id)
	if err_hex != nil {
		panic(err_hex)
	}

	err := repo.ReservationCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&foundReservation)
	defer cancel()
	return foundReservation, err
}

func (repo *ReservationRepository) CancelReservation(reservationID primitive.ObjectID) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": reservationID}

	update := bson.M{
		"$set": bson.M{
			"is_canceled": true,
		},
	}

	options := options.Update().SetUpsert(false)
	_, updateErr := repo.ReservationCollection.UpdateOne(ctx, filter, update, options)
	return updateErr
}

func (repo *ReservationRepository) DeleteLogicallyReservation(reservationID primitive.ObjectID) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": reservationID}

	update := bson.M{
		"$set": bson.M{
			"is_deleted": true,
		},
	}

	options := options.Update().SetUpsert(false)
	_, updateErr := repo.ReservationCollection.UpdateOne(ctx, filter, update, options)
	return updateErr
}

func (repo *ReservationRepository) AcceptReservation(reservationID primitive.ObjectID) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": reservationID}

	update := bson.M{
		"$set": bson.M{
			"is_accepted": true,
		},
	}

	options := options.Update().SetUpsert(false)
	_, updateErr := repo.ReservationCollection.UpdateOne(ctx, filter, update, options)
	return updateErr
}

func (repo *ReservationRepository) GetAllCanceledReservationsByGuest(guestId primitive.ObjectID) ([]models.Reservation, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"guest_id": guestId}
	cursor, err := repo.ReservationCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var reservations []models.Reservation
	if err = cursor.All(ctx, &reservations); err != nil {
		return nil, err
	}

	var canceledReservations []models.Reservation
	for _, r := range reservations {
		if r.IsCanceled == true {
			canceledReservations = append(canceledReservations, r)
		}
	}

	return canceledReservations, nil
}
