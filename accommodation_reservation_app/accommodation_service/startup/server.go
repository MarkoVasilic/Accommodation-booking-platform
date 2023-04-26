package startup

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/api"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/initializer"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/repository"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/startup/config"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/token"
	accommodation_service "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	client := initializer.ConnectToDatabase(server.config.AccommodationDBHost, server.config.AccommodationDBPort)

	accommodation_collection := initializer.AccommodationCollection(client)
	accommodation_repository := &repository.AccommodationRepository{AccommodationCollection: accommodation_collection}
	accommodation_service := &service.AccommodationService{AccommodationRepository: accommodation_repository}
	accommodation_handler := api.NewAccommodationHandler(accommodation_service)

	availability_collection := initializer.AvailabilityCollection(client)
	availability_repository := &repository.AvailabilityRepository{AvailabilityCollection: availability_collection}
	availability_service := &service.AvailabilityService{AvailabilityRepository: availability_repository}

	availability_handler := api.NewAvailabilityHandler(accommodation_service, availability_service)

	global_handler := api.NewGlobalHandler(accommodation_handler, availability_handler)

	server.startGrpcServer(global_handler)
}

func Authentication(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	methodName := info.FullMethod
	if shouldSkipInterceptor(methodName) {
		return handler(ctx, req)
	}
	ClientToken, _ := grpc_auth.AuthFromMD(ctx, "Bearer")
	if ClientToken == "" {
		err := status.Errorf(codes.InvalidArgument, "No Authorization Header Provided")
		return nil, err
	}
	_, err := token.ValidateToken(ClientToken)
	if err != "" {
		err := status.Errorf(codes.InvalidArgument, "Bad Authorization Token")
		return nil, err
	}
	if checkIsRoleHost(methodName, ClientToken) || checkIsRoleGuest(methodName, ClientToken) {
		return handler(ctx, req)
	} else {
		err := status.Errorf(codes.InvalidArgument, "Bad Authorization Token")
		return nil, err
	}
}

func checkRoles(fullMethod string, skipMethods []string) bool {
	parts := strings.Split(fullMethod, "/")
	methodName := parts[len(parts)-1]
	for _, method := range skipMethods {
		if method == methodName {
			return true
		}
	}
	return false
}

func shouldSkipInterceptor(fullMethod string) bool {
	skipMethods := []string{
		"GetAllAccommodations",
		"GetAllAvailabilities",
	}
	return checkRoles(fullMethod, skipMethods)
}

func checkIsRoleHost(fullMethod string, ClientToken string) bool {
	claims, _ := token.ValidateToken(ClientToken)
	if claims.Role == "HOST" {
		skipMethods := []string{
			"CreateAccommodation",
			"CreateAvailability",
			"UpdateAvailability",
		}
		return checkRoles(fullMethod, skipMethods)
	}
	return false
}

func checkIsRoleGuest(fullMethod string, ClientToken string) bool {
	claims, _ := token.ValidateToken(ClientToken)
	if claims.Role == "GUEST" {
		skipMethods := []string{}
		return checkRoles(fullMethod, skipMethods)
	}
	return false
}

func (server *Server) startGrpcServer(globalHandler *api.GlobalHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(Authentication))
	accommodation_service.RegisterAccommodationServiceServer(grpcServer, globalHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
