package service

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/repository"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AvailabilityService struct {
	AvailabilityRepository *repository.AvailabilityRepository
}

var Validate = validator.New()

func (service *AvailabilityService) CreateAvailability(availability models.Availability) (string, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	validationErr := Validate.Struct(availability)
	if validationErr != nil {
		err := status.Errorf(codes.InvalidArgument, "user fields are not valid")
		return "Availability fields are not valid", err
	}

	if availability.EndDate.Before(availability.StartDate) {
		err := status.Errorf(codes.InvalidArgument, "End date can not be before start date!")
		return "End date can not be before start date!", err
	}

	inserterr := service.AvailabilityRepository.CreateAvailability(&availability)

	if inserterr != nil {
		err := status.Errorf(codes.Internal, "not created")
		return "not created", err
	}

	return "Successfully created availability!", nil
}

func (service *AvailabilityService) GetAllAvailabilitiesByAccommodationID(accommodationID primitive.ObjectID) ([]models.Availability, error) {
	availabilities, err := service.AvailabilityRepository.GetAllAvailabilitiesByAccommodationID(accommodationID)
	if err != nil {
		return nil, err
	}
	return availabilities, nil
}

// availability_grpc_api -> opis kada treba? videti kako da dobavimo sve rez
func (s *AvailabilityService) UpdateAvailability(availID primitive.ObjectID, price float64) error {
	avail, err := s.AvailabilityRepository.GetAvailabilityById(availID.Hex())
	if err != nil {
		return err
	}

	avail.Price = price

	return s.AvailabilityRepository.UpdateAvailability(&avail)
}
