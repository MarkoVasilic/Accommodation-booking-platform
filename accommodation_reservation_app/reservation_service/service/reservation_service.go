package service

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/repository"
)

type ReservationService struct {
	ReservationRepository *repository.ReservationRepository
}
