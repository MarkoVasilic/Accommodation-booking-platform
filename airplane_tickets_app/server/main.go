package main

import (
	"log"
	"os"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/controllers"
	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/initializers"
	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/middleware"
	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/repositories"
	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/routes"
	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("RUN_ENV") != "production" {
		initializers.LoadEnvVariables()
	}
	client := initializers.ConnectToDatabase()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	user_collection := initializers.UserCollection(client)
	flight_collection := initializers.FlightCollection(client)
	ticket_collection := initializers.TicketCollection(client)

	public_repository := &repositories.PublicRepository{UserCollection: user_collection, FlightCollection: flight_collection}
	public_service := &services.PublicService{PublicRepository: public_repository}
	public_controller := &controllers.PublicController{PublicService: public_service}

	admin_repository := &repositories.AdminRepository{UserCollection: user_collection, FlightCollection: flight_collection}
	admin_service := &services.AdminService{AdminRepository: admin_repository}
	admin_controller := &controllers.AdminController{AdminService: admin_service}

	regular_repository := &repositories.RegularRepository{UserCollection: user_collection, FlightCollection: flight_collection, TicketCollection: ticket_collection}
	regular_service := &services.RegularService{RegularRepository: regular_repository}
	regular_controller := &controllers.RegularController{RegularService: regular_service}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "X-Api-Key", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Origin", "X-Api-Key", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	publicRoutes := router.Group("/")
	publicRoutes.Use(middleware.CORSMiddleware())
	routes.PublicRoutes(publicRoutes, public_controller)

	regularRoutes := router.Group("/")
	regularRoutes.Use(middleware.Authentication())
	regularRoutes.Use(middleware.CheckIsRoleRegular())
	routes.RegularRoutes(regularRoutes, regular_controller)

	adminRoutes := router.Group("/")
	adminRoutes.Use(middleware.Authentication())
	adminRoutes.Use(middleware.CheckIsRoleAdmin())
	routes.AdminRoutes(adminRoutes, admin_controller)

	log.Fatal(router.Run("127.0.0.1:" + port))
}
