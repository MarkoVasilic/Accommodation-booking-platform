package controllers

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/services"
	"github.com/gin-gonic/gin"
)

type RegularController struct {
	RegularService *services.RegularService
}

func (controller *RegularController) BookFlightTickets() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		controller.RegularService.BookFlightTickets(c, userID)
	}
}
