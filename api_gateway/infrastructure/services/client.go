package services

import (
	"log"

	accommodation_service "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	reservation_service "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
	user_service "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAccommodationClient(address string) accommodation_service.AccommodationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return accommodation_service.NewAccommodationServiceClient(conn)
}

func NewUserClient(address string) user_service.UserServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return user_service.NewUserServiceClient(conn)
}

func NewReservationClient(address string) reservation_service.ReservationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Reservation service: %v", err)
	}
	return reservation_service.NewReservationServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
