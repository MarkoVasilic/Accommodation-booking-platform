package api

import (
	"context"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/service"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	id := request.Id
	availabilityId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	availability, err := handler.availability_service.GetAvailabilityById(availabilityId)
	if err != nil {
		return nil, err
	}
	accommodation, err := handler.accommodation_service.GetAccommodationById(availability.AccommodationID)
	if err != nil {
		return nil, err
	}
	accommodationPb := mapAccommodation(&accommodation)
	response := &pb.GetAccommodationByAvailabilityResponse{
		Accommodation: accommodationPb,
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
