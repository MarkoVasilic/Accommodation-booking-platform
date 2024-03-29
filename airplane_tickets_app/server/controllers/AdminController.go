package controllers

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/services"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	AdminService *services.AdminService
}

func (controller *AdminController) CreateFlight() gin.HandlerFunc {
	return func(c *gin.Context) {
		controller.AdminService.CreateFlight(c)
	}
}

func (controller *AdminController) GetFlightById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		controller.AdminService.GetFlightById(c, id)
	}
}

func (controller *AdminController) DeleteFlight() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		controller.AdminService.DeleteFlightById(c, id)
	}
}
