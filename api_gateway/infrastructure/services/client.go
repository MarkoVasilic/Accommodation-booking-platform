package services

import (
	"log"

	accommodation "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAccommodationClient(address string) accommodation.AccommodationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return accommodation.NewAccommodationServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
