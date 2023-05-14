package api

import (
	"context"
	"fmt"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/models"
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
	id := request.Id
	hostId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	acs, err := handler.accommodation_service.GetAllAccommodations(hostId)
	if err != nil {
		return nil, err
	} else if acs == nil {
		err := status.Errorf(codes.InvalidArgument, "There is no accommodations!")
		return nil, err
	}
	accommodations := []*pb.Accommodation{}
	for _, a := range acs {
		accommodationPb := mapAccommodation(&a)
		accommodations = append(accommodations, accommodationPb)
	}
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
	hostID, err := primitive.ObjectIDFromHex(request.HostId)
	fmt.Println(err)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid HostId")
	}

	var images []*string
	for _, img := range request.Images {
		images = append(images, &img)
	}

	accommodation := models.Accommodation{
		HostID:     hostID,
		Name:       request.Name,
		Location:   request.Location,
		Wifi:       request.Wifi,
		Kitchen:    request.Kitchen,
		AC:         request.AC,
		ParkingLot: request.ParkingLot,
		MinGuests:  int(request.MinGuests),
		MaxGuests:  int(request.MaxGuests),
		Images:     images,
		AutoAccept: request.AutoAccept}
	mess, err := handler.accommodation_service.CreateAccommodation(accommodation)
	if err != nil {
		err := status.Errorf(codes.Internal, mess)
		return nil, err
	}

	response := &pb.CreateAccommodationResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *AccommodationHandler) DeleteAccommodationsByHost(ctx context.Context, request *pb.DeleteAccommodationsByHostRequest) (*pb.DeleteAccommodationsByHostResponse, error) {
	hostID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid HostId")
	}
	allAccommodations, err := handler.accommodation_service.GetAllAccommodations(hostID)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}

	var hostAccomodations []models.Accommodation
	for _, accommodation := range allAccommodations {
		if accommodation.HostID == hostID {
			hostAccomodations = append(hostAccomodations, accommodation)
		}
	}
	allAvailabilities, err := handler.availability_service.GetAllAvailabilities()
	if err != nil {
		return nil, err
	}

	var hostAvailabilities []models.Availability
	for _, availability := range allAvailabilities {
		for _, accommodation := range hostAccomodations {
			if availability.AccommodationID == accommodation.ID {
				hostAvailabilities = append(hostAvailabilities, availability)
			}
		}
	}

	_, err = handler.availability_service.DeleteAvailabilitiesHost(hostAvailabilities)
	if err != nil {
		err := status.Errorf(codes.Internal, "something went wrong")
		return nil, err
	}

	_, err = handler.accommodation_service.DeleteAccommodationsHost(hostAccomodations)
	if err != nil {
		err := status.Errorf(codes.Internal, "something went wrong")
		return nil, err
	}
	response := &pb.DeleteAccommodationsByHostResponse{
		Message: "success",
	}
	return response, nil
}
