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
	accommodationClient := services.NewAccommodationClient(accommodationEndpoint)
	accommodationHandler := api.NewAccommodationHandler(accommodationClient)
	accommodationHandler.Init(server.mux)
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%s", server.config.Port), server.mux))
}
