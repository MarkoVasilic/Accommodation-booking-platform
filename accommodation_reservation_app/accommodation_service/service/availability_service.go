package service

import (
	"fmt"
	"net/http"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AvailabilityService struct {
	AvailabilityRepository *repository.AvailabilityRepository
}

var Validate = validator.New()

func (service *AvailabilityService) CreateAvailability(c *gin.Context) {
	var availability models.Availability

	if err := c.BindJSON(&availability); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := Validate.Struct(availability)
	if validationErr != nil {
		fmt.Println(validationErr)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	if availability.EndDate.Before(availability.StartDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "End date can not be before start date!"})
		return
	}

	inserterr := service.AvailabilityRepository.CreateAvailability(&availability)

	if inserterr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not created"})
		return
	}

	c.JSON(http.StatusCreated, "Successfully created availability!!")
}

func (service *AvailabilityService) GetAllAvailabilities() ([]models.Availability, error) {
	availabilities, err := service.AvailabilityRepository.GetAllAvailabilities()
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
