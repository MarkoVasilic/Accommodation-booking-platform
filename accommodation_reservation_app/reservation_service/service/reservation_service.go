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
	validationErr := Validate.Struct(reservation)
	if validationErr != nil {
		err := status.Errorf(codes.InvalidArgument, "Reservation fields are not valid")
		return "Reservation fields are not valid", err
	}
	if reservation.StartDate.After(reservation.EndDate) {
		err := status.Errorf(codes.InvalidArgument, "Start date can not be after end date")
		return "Start date can not be after end date", err
	}

	reservations, err := svc.ReservationRepository.GetAllReservationsByAvailability(reservation.AvailabilityID)
	if err != nil {
		err := status.Errorf(codes.Internal, "Failed to get reservations")
		return "Failed to get reservations", err
	}

	for _, r := range reservations {
		if r.IsAccepted && !r.IsCanceled && !r.IsDeleted {
			if reservation.StartDate.Before(r.EndDate) && r.StartDate.Before(reservation.EndDate) {
				err := status.Errorf(codes.InvalidArgument, "Reservation overlaps with an existing reservation")
				return "Reservation overlaps with an existing reservation", err
			}
		}
	}

	reservationId, insertErr := svc.ReservationRepository.CreateReservation(&reservation)
	if insertErr != nil {
		err := status.Errorf(codes.Internal, "Failed to create reservation")
		return "Failed to create reservation", err
	}

	return reservationId.String(), nil
}

func (svc *ReservationService) CancelReservation(ReservationId primitive.ObjectID) (string, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservation, error := svc.ReservationRepository.GetReservationById(string(ReservationId.Hex()))
	if error != nil {
		err := status.Errorf(codes.NotFound, "There is no reservation with that id")
		return "There is no reservation with that id", err
	}

	if !reservation.IsAccepted {
		return "You cannot delete reservation that is not accepted!", status.Errorf(codes.NotFound, "You cannot delete reservation that is not accepted!")
	}

	year, month, day := reservation.StartDate.UTC().Date()
	startDate := time.Date(year, month, day, int(0), int(0), int(0), int(0), time.UTC)
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), int(0), int(0), int(0), int(0), time.UTC)

	if startDate.Sub(today) <= 24*time.Hour {
		return "You cannot cancel reservation a day before it starts!", status.Errorf(codes.NotFound, "You cannot cancel reservation a day before it starts!")
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

func (svc *ReservationService) AcceptReservation(ReservationId primitive.ObjectID) (string, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservation, error := svc.ReservationRepository.GetReservationById(string(ReservationId.Hex()))
	if error != nil {
		err1 := status.Errorf(codes.NotFound, "There is no reservation with that id")
		return "There is no reservation with that id", err1
	}

	if reservation.IsDeleted || reservation.IsCanceled {
		err3 := status.Errorf(codes.PermissionDenied, "You can't accept reservation that is deleted or canceled!")
		return "You can't accept reservation that is deleted or canceled!", err3
	}

	if reservation.IsAccepted {
		err3 := status.Errorf(codes.AlreadyExists, "You can't accept reservation that is already accepted!")
		return "You can't accept reservation that is already accepted!", err3
	}

	year, month, day := reservation.StartDate.Date()
	startDate := time.Date(year, month, day, int(0), int(0), int(0), int(0), time.UTC)
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), int(0), int(0), int(0), int(0), time.UTC)

	if startDate.Before(today) {
		err3 := status.Errorf(codes.PermissionDenied, "You can't accept expired reservation!")
		return "You can't accept expired reservation!", err3
	}

	reservations, err2 := svc.ReservationRepository.GetAllReservations()
	if err2 != nil {
		err2 := status.Errorf(codes.Internal, "Failed to get reservations")
		return "Failed to get reservations", err2
	}

	var reservationsByAvailability []models.Reservation
	for _, r := range reservations {
		if r.AvailabilityID == reservation.AvailabilityID {
			reservationsByAvailability = append(reservationsByAvailability, r)
		}
	}

	for _, r := range reservationsByAvailability {
		if !(reservation.StartDate.After(r.EndDate) || r.StartDate.After(reservation.EndDate)) && !r.IsCanceled && !r.IsDeleted {

			if reservation.ID == r.ID {
				continue
			}
			if r.IsAccepted {

				err3 := svc.ReservationRepository.DeleteLogicallyReservation(ReservationId)
				if err3 != nil {
					err3 := status.Errorf(codes.Internal, "something went wrong")
					return "something went wrong", err3
				}
				err4 := status.Errorf(codes.InvalidArgument, "Already have accepted reservation in this timespan")
				return "Reservation overlaps with an existing accepted reservation", err4

			} else {

				err5 := svc.ReservationRepository.DeleteLogicallyReservation(r.ID)
				if err5 != nil {
					err5 := status.Errorf(codes.Internal, "something went wrong")
					return "something went wrong", err5
				}
			}

		}

	}

	err6 := svc.ReservationRepository.AcceptReservation(ReservationId)
	if err6 != nil {
		err6 := status.Errorf(codes.Internal, "something went wrong")
		return "something went wrong", err6
	}

	validationErr := Validate.Struct(reservation)
	if validationErr != nil {
		err := status.Errorf(codes.InvalidArgument, "reservation fields are not valid")
		return "reservation fields are not valid", err
	}
	return "Succesffully accepted reservation", nil
}

func (svc *ReservationService) GetAllCanceledReservationsByGuest(guestId primitive.ObjectID) ([]models.Reservation, error) {
	reservations, err := svc.ReservationRepository.GetAllCanceledReservationsByGuest(guestId)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (svc *ReservationService) GetAll() ([]models.Reservation, error) {
	reservations, err := svc.ReservationRepository.GetAllReservations()
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (svc *ReservationService) DeleteReservationsHost(reservations []models.Reservation) (string, error) {
	for _, r := range reservations {
		_, err := svc.ReservationRepository.DeleteReservationById(r.ID)
		if err != nil {
			return "something went wrong", err
		}
	}
	return "success", nil
}
