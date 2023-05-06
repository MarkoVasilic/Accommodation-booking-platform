package service

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationService struct {
	AccommodationRepository *repository.AccommodationRepository
}

func (service *AccommodationService) GetAccommodationById(id primitive.ObjectID) (models.Accommodation, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	foundAccommodation, err := service.AccommodationRepository.GetAccommodationById(id)
	return foundAccommodation, err
}
