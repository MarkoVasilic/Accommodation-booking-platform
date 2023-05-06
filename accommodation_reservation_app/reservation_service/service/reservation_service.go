package service

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/repository"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReservationService struct {
	ReservationRepository *repository.ReservationRepository
}

var Validate = validator.New()

// by availability
func (svc *ReservationService) GetAllReservations(availibiltyId primitive.ObjectID) ([]models.Reservation, error) {
	reservations, err := svc.ReservationRepository.GetAllReservationsByAvailability(availibiltyId)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (svc *ReservationService) GetFindReservationPendingGuest(guestId primitive.ObjectID) ([]models.Reservation, error) {
	reservations, err := svc.ReservationRepository.GetAllReservations()
	if err != nil {
		return nil, err
	}

	var filteredReservations []models.Reservation
	for _, reservation := range reservations {
		if reservation.GuestID == guestId && !reservation.IsAccepted && !reservation.IsCanceled && !reservation.IsDeleted {
			filteredReservations = append(filteredReservations, reservation)
		}
	}
	return filteredReservations, nil
}

func (svc *ReservationService) GetFindReservationAcceptedGuest(guestId primitive.ObjectID) ([]models.Reservation, error) {
	reservations, err := svc.ReservationRepository.GetAllReservations()
	if err != nil {
		return nil, err
	}

	var filteredReservations []models.Reservation
	for _, reservation := range reservations {
		if reservation.GuestID == guestId && reservation.IsAccepted && !reservation.IsCanceled && !reservation.IsDeleted {
			filteredReservations = append(filteredReservations, reservation)
		}
	}
	return filteredReservations, nil
}

func (svc *ReservationService) CreateReservation(reservation models.Reservation) (string, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//validacija
	validationErr := Validate.Struct(reservation)
	if validationErr != nil {
		err := status.Errorf(codes.InvalidArgument, "Reservation fields are not valid")
		return "Reservation fields are not valid", err
	}

	//provera datuma
	if reservation.StartDate.After(reservation.EndDate) {
		err := status.Errorf(codes.InvalidArgument, "Start date can not be after end date")
		return "Start date can not be after end date", err
	}

	//dobavljam sve rezervacije kako bih rpoverila da li imamo preklapajucih rezervacija tj. da li je automatsko prihvatanje rezervacije
	reservations, err := svc.ReservationRepository.GetAllReservationsByAvailability(reservation.AvailabilityID)
	if err != nil {
		err := status.Errorf(codes.Internal, "Failed to get reservations")
		return "Failed to get reservations", err
	}

	//provera da li postoji automatsko prihvatanje u tom periodu -> ako da vracam error
	for _, r := range reservations {
		if r.IsAccepted && !r.IsCanceled && !r.IsDeleted {
			if reservation.StartDate.Before(r.EndDate) && r.StartDate.Before(reservation.EndDate) {
				err := status.Errorf(codes.InvalidArgument, "Reservation overlaps with an existing reservation")
				return "Reservation overlaps with an existing reservation", err
			}
		}
	}

	//ako nema preklapajucih i sve okej -> kreira se
	insertErr := svc.ReservationRepository.CreateReservation(&reservation)
	if insertErr != nil {
		err := status.Errorf(codes.Internal, "Failed to create reservation")
		return "Failed to create reservation", err
	}

	return "Successfully created reservation", nil
}

// proveriti
func (svc *ReservationService) CancelReservation(ReservationId primitive.ObjectID) (string, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservation, error := svc.ReservationRepository.GetReservationById(string(ReservationId.Hex()))
	if error != nil {
		err := status.Errorf(codes.NotFound, "There is no reservation with that id")
		return "There is no reservation with that id", err
	}

	year, month, day := reservation.StartDate.Date()
	startDate := time.Date(year, month, day, int(0), int(0), int(0), int(0), time.UTC)
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), int(0), int(0), int(0), int(0), time.UTC)

	if startDate.After(today.Add(24 * time.Hour)) {
		return "You cannot cancel reservation after day before!", nil
	}

	err := svc.ReservationRepository.CancelReservation(ReservationId)
	if err != nil {
		err := status.Errorf(codes.Internal, "something went wrong")
		return "something went wrong", err
	}

	validationErr := Validate.Struct(reservation)
	if validationErr != nil {
		err := status.Errorf(codes.InvalidArgument, "reservation fields are not valid")
		return "reservation fields are not valid", err
	}
	return "Succesffully canceled reservation", nil
}

// proveriti
func (svc *ReservationService) DeleteLogicallyReservation(ReservationId primitive.ObjectID) (string, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservation, error := svc.ReservationRepository.GetReservationById(string(ReservationId.Hex()))
	if error != nil {
		err := status.Errorf(codes.NotFound, "There is no reservation with that id")
		return "There is no reservation with that id", err
	}

	if reservation.IsAccepted {
		return "You cannot delete reservation that is accepted!", nil
	}

	err := svc.ReservationRepository.DeleteLogicallyReservation(ReservationId)
	if err != nil {
		err := status.Errorf(codes.Internal, "something went wrong")
		return "something went wrong", err
	}

	validationErr := Validate.Struct(reservation)
	if validationErr != nil {
		err := status.Errorf(codes.InvalidArgument, "reservation fields are not valid")
		return "reservation fields are not valid", err
	}
	return "Succesffully deleted reservation", nil
}
