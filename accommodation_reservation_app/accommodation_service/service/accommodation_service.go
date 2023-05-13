package service

import (
	"context"
	"log"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (service *AccommodationService) CreateAccommodation(accommodation models.Accommodation) (string, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count_name, err := service.AccommodationRepository.CountByName(*&accommodation.Name)
	if err != nil {
		log.Panic(err)
		err := status.Errorf(codes.Internal, "something went wrong")
		return "something went wrong", err
	}
	if count_name > 0 {
		err := status.Errorf(codes.NotFound, "Accommodation with that name already exists")
		return "Accommodation with that name already exists", err
	} else {
		inserterr := service.AccommodationRepository.CreateAccommodation(&accommodation)

		if inserterr != nil {
			log.Panic(err)
			err := status.Errorf(codes.Internal, "something went wrong")
			return "something went wrong", err
		}
		return "Succesffully added new accomodation", nil
	}
}

func (service *AccommodationService) GetAllAccommodationsByLocation(location string) ([]models.Accommodation, error) {
	accommodations, err := service.AccommodationRepository.GetAllAccommodationsByLocation(location)
	if err != nil {
		return nil, err
	}
	return accommodations, nil
}

func (service *AccommodationService) GetAllAccommodations() ([]models.Accommodation, error) {
	accommodations, err := service.AccommodationRepository.GetAllAccommodations()
	if err != nil {
		return nil, err
	}
	return accommodations, nil
}
