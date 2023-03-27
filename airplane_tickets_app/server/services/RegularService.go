package services

import (
	"context"
	"net/http"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegularService struct {
	RegularRepository *repositories.RegularRepository
	AdminRepository   *repositories.AdminRepository
}

func (service *RegularService) BookFlightTickets(c *gin.Context, userID string) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var userTicketsDto models.UserTickets
	if er := c.BindJSON(&userTicketsDto); er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	//validacija broja dostupnih karti
	flightID, err_hex := primitive.ObjectIDFromHex(*userTicketsDto.Flight)
	if err_hex != nil {
		panic(err_hex)
	}
	foundFlight, err := service.RegularRepository.GetFlightById(flightID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Flight with that id doesn't exist."})
		return
	}
	if int(*userTicketsDto.Number_Of_Tickets) > int(*foundFlight.Number_Of_Tickets) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Flight doesn't have that amount of tickets."})
		return
	}

	//za svaku kartu
	for i := 0; i < int(*userTicketsDto.Number_Of_Tickets); i++ {

		//napravi kartu
		ticket := new(models.Ticket)
		ticket.Flight = flightID
		ticketID, err1 := service.RegularRepository.CreateTicket(ticket)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ticket isn't booked."})
			return
		}

		//azuriraj usera
		err2 := service.RegularRepository.AddTicketToUser(ticketID, userID)
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ticket isn't add to user."})
			return
		}

	}

	//smanji broj dostupnih mesta
	err3 := service.RegularRepository.ReduceFlightTickets(flightID, *userTicketsDto.Number_Of_Tickets)
	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Number of tickets isn't reduced."})
		return
	}

	//sve ok
	c.JSON(http.StatusOK, gin.H{"message": "Ticket booked successfully."})

}
