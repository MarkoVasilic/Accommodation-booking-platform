package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	if flight.Taking_Off_Date.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't create flight in the past!"})
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

// nadje mi index u slicu
func findIndex(slice []primitive.ObjectID, val primitive.ObjectID) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}

// uklanja element iz slice-a
func removeIndex(slice []primitive.ObjectID, index int) []primitive.ObjectID {
	return append(slice[:index], slice[index+1:]...)
}

func (service *AdminService) DeleteUserTickets(c *gin.Context, flightId string) {

	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	tickets, err := service.AdminRepository.GetTicketsByFlightId(flightId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There are no tickets for that flight."})
		return
	}

	users, err := service.AdminRepository.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There are no users."})
		return
	}

	//prolazim kroz sve usere
	for _, user := range users {

		userTickets := user.UserTickets

		for _, ticket := range tickets {

			index := findIndex(userTickets, ticket.ID)

			if index != -1 {
				//uklanjam id karte iz user-ovih karata
				userTickets = removeIndex(userTickets, index)
			}

		}

		//update u bazi
		err = service.AdminRepository.UpdateUserTickets(user.ID.Hex(), userTickets)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user tickets."})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "User tickets for the given flight have been deleted."})

}

func (service *AdminService) DeleteFlightById(c *gin.Context, id string) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	//brisem sve karte kod user-a
	service.DeleteUserTickets(c, id)

	//brisem sve karte
	err1 := service.AdminRepository.DeleteTicketsByFlightId(id)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete tickets for the given flight."})
		return
	}

	//na kraju brisem let
	err2 := service.AdminRepository.DeleteFlightById(id)

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete flight. "})
		return
	}

	//poruka ako je sve ok proslo
	c.JSON(http.StatusOK, gin.H{"message": "Flight and tickets deleted successfully."})

}
