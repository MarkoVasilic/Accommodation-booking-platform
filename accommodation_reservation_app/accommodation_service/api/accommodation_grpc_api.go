package api

import (
	"context"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/service"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
)

type AccommodationHandler struct {
	pb.UnimplementedAccommodationServiceServer
	accommodation_service *service.AccommodationService
	availability_service  *service.AvailabilityService
	user_client           user_service.UserServiceClient
	reservation_client    reservation_service.ReservationServiceClient
}

func NewAccommodationHandler(accommodation_service *service.AccommodationService, availability_service *service.AvailabilityService, user_client user_service.UserServiceClient, reservation_client reservation_service.ReservationServiceClient) *AccommodationHandler {
	return &AccommodationHandler{
		accommodation_service: accommodation_service,
		availability_service:  availability_service,
		user_client:           user_client,
		reservation_client:    reservation_client,
	}
}

func (handler *AccommodationHandler) GetAllAccommodations(ctx context.Context, request *pb.GetAllAccommodationsRequest) (*pb.GetAllAccommodationsResponse, error) {
	//TODO pomocna metoda za dobavljanje svih smjestaja koje mozete koristiti u drugim mikroservisima
	accommodations := []*pb.Accommodation{}
	response := &pb.GetAllAccommodationsResponse{
		Accommodations: accommodations,
	}
	return response, nil
}

func (handler *AccommodationHandler) GetAccommodationByAvailability(ctx context.Context, request *pb.GetAccommodationByAvailabilityRequest) (*pb.GetAccommodationByAvailabilityResponse, error) {
	//TODO
	response := &pb.GetAccommodationByAvailabilityResponse{
		Accommodation: nil,
	}
	return response, nil
}

func (handler *AccommodationHandler) CreateAccommodation(ctx context.Context, request *pb.CreateAccommodationRequest) (*pb.CreateAccommodationResponse, error) {
	//TODO
	response := &pb.CreateAccommodationResponse{
		Message: "Success",
	}
	return response, nil
}
