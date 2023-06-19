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

func (handler *GlobalHandler) GetAllAccommodations(ctx context.Context, request *pb.GetAllAccommodationsRequest) (*pb.GetAllAccommodationsResponse, error) {
	return handler.accommodationHandler.GetAllAccommodations(ctx, request)
}

func (handler *GlobalHandler) GetAllAvailabilities(ctx context.Context, request *pb.GetAllAvailabilitiesRequest) (*pb.GetAllAvailabilitiesResponse, error) {
	return handler.availabilityHandler.GetAllAvailabilities(ctx, request)
}

func (handler *GlobalHandler) GetAccommodationByAvailability(ctx context.Context, request *pb.GetAccommodationByAvailabilityRequest) (*pb.GetAccommodationByAvailabilityResponse, error) {
	return handler.accommodationHandler.GetAccommodationByAvailability(ctx, request)
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

func (handler *GlobalHandler) DeleteAccommodationsByHost(ctx context.Context, request *pb.DeleteAccommodationsByHostRequest) (*pb.DeleteAccommodationsByHostResponse, error) {
	return handler.accommodationHandler.DeleteAccommodationsByHost(ctx, request)
}

func (handler *GlobalHandler) GetAvailabilityById(ctx context.Context, request *pb.GetAvailabilityByIdRequest) (*pb.GetAvailabilityByIdResponse, error) {
	return handler.availabilityHandler.GetAvailabilityById(ctx, request)
}

func (handler *GlobalHandler) GetAccommodationById(ctx context.Context, request *pb.GetAccommodationByIdRequest) (*pb.GetAccommodationByIdResponse, error) {
	return handler.accommodationHandler.GetAccommodationById(ctx, request)
}

func (handler *GlobalHandler) GetEveryAccommodation(ctx context.Context, request *pb.GetEveryAccommodationRequest) (*pb.GetEveryAccommodationResponse, error) {
	return handler.accommodationHandler.GetEveryAccommodation(ctx, request)
}

func (handler *GlobalHandler) GetAllAccommodationGrade(ctx context.Context, request *pb.GetAllAccommodationGradeRequest) (*pb.GetAllAccommodationGradeResponse, error) {
	return handler.accommodationHandler.GetAllAccommodationGrade(ctx, request)
}

func (handler *GlobalHandler) GetAllAccommodationGuestGrades(ctx context.Context, request *pb.GetAllAccommodationGuestGradesRequest) (*pb.GetAllAccommodationGuestGradesResponse, error) {
	return handler.accommodationHandler.GetAllAccommodationGuestGrades(ctx, request)
}

func (handler *GlobalHandler) CreateAccommodationGrade(ctx context.Context, request *pb.CreateAccommodationGradeRequest) (*pb.CreateAccommodationGradeResponse, error) {
	return handler.accommodationHandler.CreateAccommodationGrade(ctx, request)
}

func (handler *GlobalHandler) UpdateAccommodationGrade(ctx context.Context, request *pb.UpdateAccommodationGradeRequest) (*pb.UpdateAccommodationGradeResponse, error) {
	return handler.accommodationHandler.UpdateAccommodationGrade(ctx, request)
}

func (handler *GlobalHandler) DeleteAccommodationGrade(ctx context.Context, request *pb.DeleteAccommodationGradeRequest) (*pb.DeleteAccommodationGradeResponse, error) {
	return handler.accommodationHandler.DeleteAccommodationGrade(ctx, request)
}

func (handler *GlobalHandler) FilterAvailability(ctx context.Context, request *pb.FilterAvailabilityRequest) (*pb.FilterAvailabilityResponse, error) {
	return handler.availabilityHandler.FilterAvailability(ctx, request)
}
