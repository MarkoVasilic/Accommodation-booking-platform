package startup

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/api"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/initializer"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/repository"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/startup/config"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/token"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	reservation "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
	saga "github.com/MarkoVasilic/Accommodation-booking-platform/common/saga/messaging"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/saga/messaging/nats"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.mongodb.org/mongo-driver/mongo"
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
	QueueGroup = "reservation_service"
)

func (server *Server) InitializeAccommodationClient() accommodation_service.AccommodationServiceClient {
	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	conn, err := grpc.Dial(accommodationEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return accommodation_service.NewAccommodationServiceClient(conn)
}

func (server *Server) InitializeUserClient() user_service.UserServiceClient {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	conn, err := grpc.Dial(userEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return user_service.NewUserServiceClient(conn)
}

func (server *Server) Start() {
	client := initializer.ConnectToDatabase(server.config.ReservationDBHost, server.config.ReservationDBPort)

	reservation_collection := initializer.ReservationCollection(client)
	reservation_repository := &repository.ReservationRepository{ReservationCollection: reservation_collection}
	reservation_service := &service.ReservationService{ReservationRepository: reservation_repository}
	accommodation_client := server.InitializeAccommodationClient()
	user_client := server.InitializeUserClient()

	commandSubscriber := server.initSubscriber(server.config.DeleteUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.DeleteUserReplySubject)
	server.initDeleteUserHandler(reservation_service, replyPublisher, commandSubscriber, accommodation_client, client)

	reservation_handler := api.NewReservationHandler(reservation_service, accommodation_client, user_client)

	server.startGrpcServer(reservation_handler)
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

func (server *Server) initDeleteUserHandler(service *service.ReservationService, publisher saga.Publisher, subscriber saga.Subscriber, accommodation_client accommodation_service.AccommodationServiceClient, mongo_client *mongo.Client) {
	_, err := api.NewDeleteUserCommandHandler(service, publisher, subscriber, accommodation_client, mongo_client)
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
		"GetAllReservations",
	}
	return checkRoles(fullMethod, skipMethods)
}

func checkIsRoleHost(fullMethod string, ClientToken string) bool {
	claims, _ := token.ValidateToken(ClientToken)
	if claims.Role == "HOST" {
		skipMethods := []string{
			"GetFindReservationHost",
			"DeleteLogicallyReservation",
			"AcceptReservation",
			"GetAllReservationsHost",
			"DeleteReservationsHost",
		}
		return checkRoles(fullMethod, skipMethods)
	}
	return false
}

func checkIsRoleGuest(fullMethod string, ClientToken string) bool {
	claims, _ := token.ValidateToken(ClientToken)
	if claims.Role == "GUEST" {
		skipMethods := []string{
			"CreateReservation",
			"GetFindReservationPendingGuest",
			"GetFindReservationAcceptedGuest",
			"CancelReservation",
			"DeleteLogicallyReservation",
			"GetAllReservationsByGuestId",
		}
		return checkRoles(fullMethod, skipMethods)
	}
	return false
}

func (server *Server) startGrpcServer(reservationHandler *api.ReservationHandler) {
	url := fmt.Sprintf("reservation_service:%s", server.config.Port)
	if os.Getenv("RUN_ENV") == "production" {
		url = fmt.Sprintf("reservation_service:%s", server.config.Port)
	} else {
		url = fmt.Sprintf("localhost:%s", server.config.Port)
	}
	listener, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(Authentication))
	reservation.RegisterReservationServiceServer(grpcServer, reservationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
