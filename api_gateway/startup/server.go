package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/MarkoVasilic/Accommodation-booking-platform/api_gateway/infrastructure/api"
	"github.com/MarkoVasilic/Accommodation-booking-platform/api_gateway/infrastructure/services"
	cfg "github.com/MarkoVasilic/Accommodation-booking-platform/api_gateway/startup/config"
	accommodationGw "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	reservationGw "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
	userGw "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	err := accommodationGw.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationEndpoint, opts)
	if err != nil {
		panic(err)
	}
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err = userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}
	reservationEndpoint := fmt.Sprintf("%s:%s", server.config.ReservationHost, server.config.ReservationPort)
	err = reservationGw.RegisterReservationServiceHandlerFromEndpoint(context.TODO(), server.mux, reservationEndpoint, opts)
	if err != nil {
		panic(err)
	}
	accommodationClient := services.NewAccommodationClient(accommodationEndpoint)
	userClient := services.NewUserClient(userEndpoint)
	reservationClient := services.NewReservationClient(reservationEndpoint)
	globalHandler := api.NewGlobalHandler(accommodationClient, userClient, reservationClient)
	globalHandler.Init(server.mux)
}

func (server *Server) Start() {
	// Define the allowed origins
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	// Define the allowed methods
	allowedMethods := handlers.AllowedMethods([]string{"PUT", "PATCH", "POST", "GET", "OPTIONS", "DELETE"})
	// Define the allowed headers
	allowedHeaders := handlers.AllowedHeaders([]string{"Origin", "X-Api-Key", "X-Requested-With", "Content-Type", "Accept", "Authorization"})

	// Wrap the server.mux with CORS handlers
	handler := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(server.mux)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%s", server.config.Port), handler))
}
