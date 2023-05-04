package service

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationService struct {
	ReservationRepository *repository.ReservationRepository
}

// by availability
func (svc *ReservationService) GetAllReservations(availibiltyId primitive.ObjectID) ([]models.Reservation, error) {
	reservations, err := svc.ReservationRepository.GetAllReservations(availibiltyId)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}
