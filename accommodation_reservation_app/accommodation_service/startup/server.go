package startup

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/api"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/initializer"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/repository"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/startup/config"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/token"
	accommodation_service "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
	saga "github.com/MarkoVasilic/Accommodation-booking-platform/common/saga/messaging"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/saga/messaging/nats"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Server struct {
	config *config.Config
}

const (
	QueueGroup = "accommodation_service"
)

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) InitializeUserClient() user_service.UserServiceClient {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	conn, err := grpc.Dial(userEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return user_service.NewUserServiceClient(conn)
}

func (server *Server) InitializeReservationClient() reservation_service.ReservationServiceClient {
	reservationEndpoint := fmt.Sprintf("%s:%s", server.config.ReservationHost, server.config.ReservationPort)
	conn, err := grpc.Dial(reservationEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Reservation service: %v", err)
	}
	return reservation_service.NewReservationServiceClient(conn)
}

func (server *Server) Start() {
	client := initializer.ConnectToDatabase(server.config.AccommodationDBHost, server.config.AccommodationDBPort)

	accommodation_collection := initializer.AccommodationCollection(client)
	accommodation_repository := &repository.AccommodationRepository{AccommodationCollection: accommodation_collection}
	accommodation_service := &service.AccommodationService{AccommodationRepository: accommodation_repository}

	availability_collection := initializer.AvailabilityCollection(client)
	availability_repository := &repository.AvailabilityRepository{AvailabilityCollection: availability_collection}
	availability_service := &service.AvailabilityService{AvailabilityRepository: availability_repository}

	grade_collection := initializer.GradeCollection(client)
	grade_repository := &repository.GradeRepository{GradeCollection: grade_collection}
	grade_service := &service.GradeService{GradeRepository: grade_repository}

	user_client := server.InitializeUserClient()
	reservation_client := server.InitializeReservationClient()

	commandSubscriber := server.initSubscriber(server.config.DeleteUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.DeleteUserReplySubject)
	server.initDeleteUserHandler(accommodation_service, availability_service, replyPublisher, commandSubscriber)

	accommodation_handler := api.NewAccommodationHandler(accommodation_service, availability_service, grade_service, user_client, reservation_client)
	availability_handler := api.NewAvailabilityHandler(accommodation_service, availability_service, grade_service, user_client, reservation_client)

	global_handler := api.NewGlobalHandler(accommodation_handler, availability_handler)

	server.startGrpcServer(global_handler)
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initDeleteUserHandler(acc_service *service.AccommodationService, ava_service *service.AvailabilityService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewDeleteUserCommandHandler(acc_service, ava_service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
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
	fmt.Println(err)
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
		"SearchAvailability",
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
			"SearchAvailability",
			"GetAllAccommodations",
			"GetAllAvailabilities",
			"GetAccommodationByAvailability",
			"DeleteAccommodationsByHost",
		}
		return checkRoles(fullMethod, skipMethods)
	}
	return false
}

func checkIsRoleGuest(fullMethod string, ClientToken string) bool {
	claims, _ := token.ValidateToken(ClientToken)
	if claims.Role == "GUEST" {
		skipMethods := []string{
			"GetAccommodationByAvailability",
			"SearchAvailability",
			"GetAvailabilityById",
			"GetAccommodationById",
		}
		return checkRoles(fullMethod, skipMethods)
	}
	return false
}

func (server *Server) startGrpcServer(globalHandler *api.GlobalHandler) {
	url := fmt.Sprintf("accommodation_service:%s", server.config.Port)
	if os.Getenv("RUN_ENV") == "production" {
		url = fmt.Sprintf("accommodation_service:%s", server.config.Port)
	} else {
		url = fmt.Sprintf("localhost:%s", server.config.Port)
	}
	listener, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(Authentication))
	accommodation_service.RegisterAccommodationServiceServer(grpcServer, globalHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
