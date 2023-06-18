package startup

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/api"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/initializer"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/repository"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/startup/config"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/token"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
	user "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
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

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "user_service"
)

func (server *Server) InitializeAccommodationClient() accommodation_service.AccommodationServiceClient {
	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	conn, err := grpc.Dial(accommodationEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return accommodation_service.NewAccommodationServiceClient(conn)
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
	client := initializer.ConnectToDatabase(server.config.UserDBHost, server.config.UserDBPort)
	user_collection := initializer.UserCollection(client)
	user_repository := &repository.UserRepository{UserCollection: user_collection}

	grade_collection := initializer.GradeCollection(client)
	grade_repository := &repository.GradeRepository{GradeCollection: grade_collection}
	grade_service := &service.GradeService{GradeRepository: grade_repository}

	notification_collection := initializer.NotificationCollection(client)
	notification_repository := &repository.NotificationRepository{NotificationCollection: notification_collection}
	notification_service := &service.NotificationService{NotificationRepository: notification_repository}

	notification_on_collection := initializer.NotificationOnCollection(client)
	notification_on_repository := &repository.NotificationOnRepository{NotificationOnCollection: notification_on_collection}
	notification_on_service := &service.NotificationOnService{NotificationOnRepository: notification_on_repository}

	commandPublisher := server.initPublisher(server.config.DeleteUserCommandSubject)
	replySubscriber := server.initSubscriber(server.config.DeleteUserReplySubject, QueueGroup)
	deleteUserOrchestrator := server.initDeleteUserOrchestrator(commandPublisher, replySubscriber)

	user_service := &service.UserService{UserRepository: user_repository, Orchestrator: deleteUserOrchestrator}

	accommodation_client := server.InitializeAccommodationClient()
	reservation_client := server.InitializeReservationClient()
	user_handler := api.NewUserHandler(user_service, grade_service, notification_service, notification_on_service, accommodation_client, reservation_client)

	commandSubscriber := server.initSubscriber(server.config.DeleteUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.DeleteUserReplySubject)
	server.initDeleteUserHandler(user_service, replyPublisher, commandSubscriber, accommodation_client, reservation_client)

	server.startGrpcServer(user_handler)
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

func (server *Server) initDeleteUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *service.DeleteUserOrchestrator {
	orchestrator, err := service.NewDeleteUserOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initDeleteUserHandler(user_service *service.UserService, publisher saga.Publisher, subscriber saga.Subscriber, accommodation_client accommodation_service.AccommodationServiceClient, reservation_client reservation_service.ReservationServiceClient) {
	_, err := api.NewDeleteUserCommandHandler(user_service, publisher, subscriber, accommodation_client, reservation_client)
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
		"GetLoggedUser",
		"GetUser",
		"CreateUser",
		"Login",
		"SearchAvailability",
		"UpdateNotificationOn",
		"CreateNotification",
		"GetAllNotifications",
		"GetUserNotificationsOn",
		"HostProminent",
	}
	return checkRoles(fullMethod, skipMethods)
}

func checkIsRoleHost(fullMethod string, ClientToken string) bool {
	claims, _ := token.ValidateToken(ClientToken)
	if claims.Role == "HOST" {
		skipMethods := []string{
			"UpdateUser",
			"DeleteUser",
			"CreateAccommodation",
			"CreateAvailability",
			"UpdateAvailability",
			"GetFindReservationHost",
			"DeleteLogicallyReservation",
			"AcceptReservation",
			"HostProminent",
			"GetAllUserGrade",
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
			"CreateReservation",
			"GetFindReservationPendingGuest",
			"GetFindReservationAcceptedGuest",
			"CancelReservation",
			"DeleteLogicallyReservation",
			"CreateUserGrade",
			"UpdateUserGrade",
			"DeleteUserGrade",
			"GetAllUserGrade",
			"GetAllHosts",
			"GetAllGuestGrades",
			"HostProminent",
		}
		return checkRoles(fullMethod, skipMethods)
	}
	return false
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {
	url := fmt.Sprintf("user_service:%s", server.config.Port)
	if os.Getenv("RUN_ENV") == "production" {
		url = fmt.Sprintf("user_service:%s", server.config.Port)
	} else {
		url = fmt.Sprintf("localhost:%s", server.config.Port)
	}
	listener, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(Authentication))
	user.RegisterUserServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
