package service

import (
	"context"
	"log"
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

	allAvailabilities, err := service.AvailabilityRepository.GetAllAvailabilitiesByAccommodationID(availability.AccommodationID)
	if err != nil {
		err := status.Errorf(codes.Internal, "Failed to retrieve availabilities")
		return "Failed to retrieve availabilities", err
	}

	for _, existingAvailability := range allAvailabilities {
		if availability.StartDate.Before(existingAvailability.EndDate) && existingAvailability.StartDate.Before(availability.EndDate) {
			err := status.Errorf(codes.AlreadyExists, "Overlap with existing availability")
			return "Overlap with existing availability", err
		}
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

func (service *AvailabilityService) GetAvailabilityById(id primitive.ObjectID) (models.Availability, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	foundAvailability, err := service.AvailabilityRepository.GetAvailabilityById(id)
	return foundAvailability, err
}

func (service AvailabilityService) UpdateAvailability(availability models.Availability, id string) (string, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(id)

	_, error := service.AvailabilityRepository.GetAvailabilityById(objectId)
	if error != nil {
		err := status.Errorf(codes.NotFound, "There is no reservation with that id")
		return "There is no reservation with that id", err
	}

	validationErr := Validate.Struct(availability)
	if validationErr != nil {
		err := status.Errorf(codes.InvalidArgument, "availability fields are not valid")
		return "availability fields are not valid", err
	}

	availability.ID = objectId
	inserterr := service.AvailabilityRepository.UpdateAvailability(&availability)
	if inserterr != nil {
		log.Panic(err)
		err := status.Errorf(codes.Internal, "something went wrong")
		return "something went wrong", err
	}
	return "Succesffully updated availability", nil
}

func (service *AvailabilityService) GetAllAvailabilitiesByDates(startDate time.Time, endDate time.Time) ([]models.Availability, error) {
	availabilities, err := service.AvailabilityRepository.GetAllAvailabilities()
	var filteredAvailabilities []models.Availability
	for _, avail := range availabilities {
		if avail.StartDate.Before(startDate) && avail.EndDate.After(endDate) {
			filteredAvailabilities = append(filteredAvailabilities, avail)
		}
	}
	if filteredAvailabilities == nil {
		er := status.Errorf(codes.InvalidArgument, "There is no available accommodation for chosen dates!")
		return nil, er
	}
	if err != nil {
		return nil, err
	}
	return filteredAvailabilities, nil
}

func (service *AvailabilityService) GetAllAvailabilities() ([]models.Availability, error) {
	availabilities, err := service.AvailabilityRepository.GetAllAvailabilities()
	if err != nil {
		return nil, err
	}
	return availabilities, nil
}

func (service *AvailabilityService) DeleteAvailabilitiesHost(availabilities []models.Availability) (string, error) {
	for _, r := range availabilities {
		_, err := service.AvailabilityRepository.DeleteAvailability(r.ID)
		if err != nil {
			return "something went wrong", err
		}
	}
	return "success", nil
}
