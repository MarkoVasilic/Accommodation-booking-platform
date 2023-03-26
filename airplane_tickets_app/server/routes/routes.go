package routes

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/controllers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(incomingRoutes *gin.RouterGroup, PublicController *controllers.PublicController) {
	incomingRoutes.POST("/users/signup/", PublicController.SignUp())
	incomingRoutes.POST("/users/login/", PublicController.Login())
	incomingRoutes.GET("/users/logged/", PublicController.GetUserById())
	//incomingRoutes.GET("/flights/all/", PublicController.GetAllFlights()) sa filtriranjem i sortiranjem, koristiti dto zbog cijena
}

func AdminRoutes(incomingRoutes *gin.RouterGroup, AdminController *controllers.AdminController) {
	incomingRoutes.POST("/flights/create/", AdminController.CreateFlight())
	incomingRoutes.GET("/flights/info/:id", AdminController.GetFlightById())
	//incomingRoutes.GET("/flights/tickets_left/:id", AdminController.TicketsLeft()) moze i na frontu samo
	incomingRoutes.DELETE("/flights/delete/:id", AdminController.DeleteFlight())
}

func RegularRoutes(incomingRoutes *gin.RouterGroup, RegularController *controllers.RegularController) {
	//incomingRoutes.PUT("/tickets/buy/:id", RegularController.UpdateFlight())  kupovina karata, trebalo bi da se posalje id leta, usera i broj kupljenih karata
	//zatim napraviti nove objekte karata i uvezati ih sa letom i sa userom, znaci treba biti ukljucen update usera posto sadrzi listu kupljenih karata
	//azurirati let tako da se smanji number_of_tickets ako ima dostupnih (ne mora ovako url da izgleda i zahtjev moze da bude post)
	//incomingRoutes.GET("/ticket/all/:id", RegularController.GetAllTickets()) na osnovu id usera dobaviti sve karte, i informacije izvuci iz leta za koju je karta vezana
}
