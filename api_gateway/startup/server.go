package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

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
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"PUT", "PATCH", "POST", "GET", "OPTIONS", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Origin", "X-Api-Key", "X-Requested-With", "Content-Type", "Accept", "Authorization"})
	handler := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(server.mux)
	fmt.Println("api gateway running")
	url := fmt.Sprintf("127.0.0.1:%s", server.config.Port)
	if os.Getenv("RUN_ENV") == "production" {
		url = fmt.Sprintf("0.0.0.0:%s", server.config.Port)
	} else {
		url = fmt.Sprintf("127.0.0.1:%s", server.config.Port)
	}
	log.Fatal(http.ListenAndServe(url, handler))
}
