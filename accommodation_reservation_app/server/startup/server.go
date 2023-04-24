package startup

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/server/api"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/server/initializer"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/server/repository"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/server/service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/server/startup/config"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/server/token"
	accommodation "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
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
	user_collection := initializer.UserCollection(client)
	user_repository := &repository.UserRepository{UserCollection: user_collection}
	user_service := &service.UserService{UserRepository: user_repository}
	user_handler := api.NewUserHandler(user_service)

	server.startGrpcServer(user_handler)
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
		"GetUser",
		"CreateUser",
		"Login",
	}
	return checkRoles(fullMethod, skipMethods)
}

func checkIsRoleHost(fullMethod string, ClientToken string) bool {
	claims, _ := token.ValidateToken(ClientToken)
	if claims.Role == "HOST" {
		skipMethods := []string{
			"UpdateUser",
			"DeleteUser",
		}
		return checkRoles(fullMethod, skipMethods)
	}
	return false
}

func checkIsRoleGuest(fullMethod string, ClientToken string) bool {
	claims, _ := token.ValidateToken(ClientToken)
	if claims.Role == "GUEST" {
		skipMethods := []string{
			"UpdateUser",
			"DeleteUser",
		}
		return checkRoles(fullMethod, skipMethods)
	}
	return false
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(Authentication))
	accommodation.RegisterAccommodationServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
