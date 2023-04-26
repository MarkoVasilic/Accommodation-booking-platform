package api

import (
	"context"

	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
)

type GlobalHandler struct {
	pb.UnimplementedAccommodationServiceServer
	accommodationHandler *AccommodationHandler
	availabilityHandler  *AvailabilityHandler
}

func NewGlobalHandler(accommodationHandler *AccommodationHandler, availabilityHandler *AvailabilityHandler) *GlobalHandler {
	return &GlobalHandler{
		accommodationHandler: accommodationHandler,
		availabilityHandler:  availabilityHandler,
	}
}

func (handler *GlobalHandler) CreateAccommodation(ctx context.Context, request *pb.CreateAccommodationRequest) (*pb.CreateAccommodationResponse, error) {
	return handler.accommodationHandler.CreateAccommodation(ctx, request)
}

func (handler *GlobalHandler) CreateAvailability(ctx context.Context, request *pb.CreateAvailabilityRequest) (*pb.CreateAvailabilityResponse, error) {
	return handler.availabilityHandler.CreateAvailability(ctx, request)
}

func (handler *GlobalHandler) UpdateAvailability(ctx context.Context, request *pb.UpdateAvailabilityRequest) (*pb.UpdateAvailabilityResponse, error) {
	return handler.availabilityHandler.UpdateAvailability(ctx, request)
}

func (handler *GlobalHandler) SearchAvailability(ctx context.Context, request *pb.SearchAvailabilityRequest) (*pb.SearchAvailabilityResponse, error) {
	return handler.availabilityHandler.SearchAvailability(ctx, request)
}
