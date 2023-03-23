package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/services"
	token "github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/tokens"
)

type PublicController struct {
	PublicService *services.PublicService
}

func (controller *PublicController) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		controller.PublicService.CreateUser(c)
	}
}

func (controller *PublicController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		controller.PublicService.VerifyUser(c)
	}
}

func (controller *PublicController) GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("Authorization")
		if ClientToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No Authorization Header Provided"})
			c.Abort()
			return
		}
		claims, err := token.ValidateToken(ClientToken)
		if err != "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Nobody is logged"})
			c.Abort()
			return
		}
		controller.PublicService.GetUserById(c, claims.Uid)
	}
}

func (controller *PublicController) SearchedFlights() gin.HandlerFunc {
	return func(c *gin.Context) {
		controller.PublicService.SearchedFlights(c)
	}
}
