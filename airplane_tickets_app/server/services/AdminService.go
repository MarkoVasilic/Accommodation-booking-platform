package services

import (
	"fmt"
	"net/http"

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
