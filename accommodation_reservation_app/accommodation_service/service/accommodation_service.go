package service

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/repository"
)

type AccommodationService struct {
	AccommodationRepository *repository.AccommodationRepository
}
