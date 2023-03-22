package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/repositories"
	"github.com/gin-gonic/gin"
)

type AdminService struct {
	AdminRepository *repositories.AdminRepository
}

func (service *AdminService) CreateFlight(c *gin.Context) {
	var flight models.Flight
	if err := c.BindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErr := Validate.Struct(flight)
	if validationErr != nil {
		fmt.Println(validationErr)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	inserterr := service.AdminRepository.CreateFlight(&flight)

	if inserterr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not created"})
		return
	}
	c.JSON(http.StatusCreated, "Successfully created flight!!")
}

func (service *AdminService) GetFlightById(c *gin.Context, id string) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	foundflight, err := service.AdminRepository.GetFlightById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Flight with that id doesn't exist."})
		return
	}
	c.JSON(http.StatusOK, foundflight)
}

func (service *AdminService) DeleteFlightById(c *gin.Context, id string) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err := service.AdminRepository.DeleteFlightById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete flight. "})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Flight deleted successfully. "})
}
