package api

import (
	"context"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/service"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
)

type AccommodationHandler struct {
	pb.UnimplementedAccommodationServiceServer
	accommodation_service *service.AccommodationService
}

func NewAccommodationHandler(accommodation_service *service.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		accommodation_service: accommodation_service,
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

func (handler *AccommodationHandler) CreateAccommodation(ctx context.Context, request *pb.CreateAccommodationRequest) (*pb.CreateAccommodationResponse, error) {
	//TODO
	response := &pb.CreateAccommodationResponse{
		Message: "Success",
	}
	return response, nil
}
