package service

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/repository"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/models"
)

type ReservationService struct {
	ReservationRepository *repository.ReservationRepository
}

func (svc *ReservationService) GetAllReservations() ([]models.Reservation, error) {
	reservations, err := svc.ReservationRepository.GetAllReservations()
	if err != nil {
		return nil, err
	}
	return reservations, nil
}
