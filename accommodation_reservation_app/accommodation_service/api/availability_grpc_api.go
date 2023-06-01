package api

import (
	"context"
	"strings"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/service"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AvailabilityHandler struct {
	pb.UnimplementedAccommodationServiceServer
	accommodation_service *service.AccommodationService
	availability_service  *service.AvailabilityService
	grade_service         *service.GradeService
	user_client           user_service.UserServiceClient
	reservation_client    reservation_service.ReservationServiceClient
}

func NewAvailabilityHandler(accommodation_service *service.AccommodationService, availability_service *service.AvailabilityService, grade_service *service.GradeService, user_client user_service.UserServiceClient, reservation_client reservation_service.ReservationServiceClient) *AvailabilityHandler {
	return &AvailabilityHandler{
		accommodation_service: accommodation_service,
		availability_service:  availability_service,
		grade_service:         grade_service,
		user_client:           user_client,
		reservation_client:    reservation_client,
	}
}

func createContextForAuthorization(ctx context.Context) context.Context {
	token, _ := grpc_auth.AuthFromMD(ctx, "Bearer")
	if len(token) > 0 {
		return metadata.NewOutgoingContext(context.Background(), metadata.Pairs("Authorization", "Bearer "+token))
	}
	return context.TODO()
}

func (handler *AvailabilityHandler) GetAllAvailabilities(ctx context.Context, request *pb.GetAllAvailabilitiesRequest) (*pb.GetAllAvailabilitiesResponse, error) {
	id := request.Id
	accomodationId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}

	temp, err := primitive.ObjectIDFromHex("64580a2e9f857372a34602c2")
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}

	var as []models.Availability
	if accomodationId == temp {
		as, err = handler.availability_service.GetAllAvailabilities()
	} else {
		as, err = handler.availability_service.GetAllAvailabilitiesByAccommodationID(accomodationId)
	}

	if err != nil {
		return nil, err
	} else if as == nil {
		err := status.Errorf(codes.InvalidArgument, "There is no availabilities!")
		return nil, err
	}
	availabilities := []*pb.Availability{}
	for _, a := range as {
		availabilityPb := mapAvailability(&a)
		availabilities = append(availabilities, availabilityPb)
	}

	response := &pb.GetAllAvailabilitiesResponse{
		Availabilities: availabilities,
	}
	return response, nil
}

func (handler *AvailabilityHandler) GetAvailabilityById(ctx context.Context, request *pb.GetAvailabilityByIdRequest) (*pb.GetAvailabilityByIdResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	availability, err := handler.availability_service.GetAvailabilityById(objectId)
	if err != nil {
		return nil, err
	}
	availabilityPb := mapAvailability(&availability)
	response := &pb.GetAvailabilityByIdResponse{
		Availability: availabilityPb,
	}
	return response, nil
}

func (handler *AvailabilityHandler) CreateAvailability(ctx context.Context, request *pb.CreateAvailabilityRequest) (*pb.CreateAvailabilityResponse, error) {
	accommodationID, err := primitive.ObjectIDFromHex(request.AccommodationID)
	if err != nil {
		return nil, err
	}

	startDate := request.StartDate.AsTime()
	endDate := request.EndDate.AsTime()

	accommodation, err := handler.accommodation_service.GetAccommodationById(accommodationID)
	if err != nil {
		return nil, err
	}

	availability := models.Availability{AccommodationID: accommodation.ID, StartDate: startDate, EndDate: endDate, Price: request.Price, IsPricePerGuest: request.IsPricePerGuest}
	mess, err := handler.availability_service.CreateAvailability(availability)
	if err != nil {
		err := status.Errorf(codes.Internal, mess)
		return nil, err
	}
	response := &pb.CreateAvailabilityResponse{
		Message: mess,
	}
	return response, nil
}

func (handler *AvailabilityHandler) UpdateAvailability(ctx context.Context, request *pb.UpdateAvailabilityRequest) (*pb.UpdateAvailabilityResponse, error) {
	Id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	res, err := handler.reservation_client.GetAllReservations(createContextForAuthorization(ctx), &reservation_service.GetAllReservationsRequest{Id: request.Id}) //&request.Id?
	if err != nil {
		return nil, err
	} else if res != nil {
		var acceptedReservations []*reservation_service.Reservation
		for _, reservation := range res.Reservations {
			if reservation.IsAccepted && !reservation.IsCanceled && !reservation.IsDeleted {
				acceptedReservations = append(acceptedReservations, reservation)
			}
		}
		if acceptedReservations != nil {
			err := status.Errorf(codes.InvalidArgument, "Cannot update availability price because there are reservations in that period!")
			return nil, err
		}
		avail, err := handler.availability_service.GetAvailabilityById(Id)
		if err != nil {
			return nil, err
		}
		availability := models.Availability{ID: Id, AccommodationID: avail.AccommodationID, StartDate: request.StartDate.AsTime(), EndDate: request.EndDate.AsTime(), Price: request.Price, IsPricePerGuest: request.IsPricePerGuest}
		mess, err := handler.availability_service.UpdateAvailability(availability, request.Id)
		if err != nil {
			err := status.Errorf(codes.Internal, mess)
			return nil, err
		}
	} else if res == nil {
		availability := models.Availability{ID: Id, StartDate: request.StartDate.AsTime(), EndDate: request.EndDate.AsTime(), Price: request.Price, IsPricePerGuest: request.IsPricePerGuest}
		mess, err := handler.availability_service.UpdateAvailability(availability, request.Id)
		if err != nil {
			err := status.Errorf(codes.Internal, mess)
			return nil, err
		}
	}

	response := &pb.UpdateAvailabilityResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *AvailabilityHandler) SearchAvailability(ctx context.Context, request *pb.SearchAvailabilityRequest) (*pb.SearchAvailabilityResponse, error) {
	year, month, day := request.StartDate.AsTime().Date()
	yearE, monthE, dayE := request.EndDate.AsTime().Date()
	startDate := time.Date(year, month, day, int(0), int(0), int(0), int(0), time.UTC)
	endDate := time.Date(yearE, monthE, dayE, int(0), int(0), int(0), int(0), time.UTC)
	availabilities, err := handler.availability_service.GetAllAvailabilitiesByDates(startDate, endDate)
	if err != nil {
		return nil, err
	}
	favailabilities := []models.FindAvailability{}
	for _, avail := range availabilities {
		accommodation, err := handler.accommodation_service.GetAccommodationById(avail.AccommodationID)
		if err != nil {
			return nil, err
		}
		reservations, err := handler.reservation_client.GetAllReservations(createContextForAuthorization(ctx), &reservation_service.GetAllReservationsRequest{Id: string(avail.ID.Hex())})
		if reservations.Reservations == nil {
			if accommodation.Location == request.Location && accommodation.MinGuests <= int(request.GuestsNum) && accommodation.MaxGuests >= int(request.GuestsNum) {
				duration := endDate.Sub(startDate)
				nights := int(duration.Hours() / 24)
				totalPrice := avail.Price * float64(nights)
				findAvailability := models.FindAvailability{AccommodationId: accommodation.ID, AvailabilityID: avail.ID, HostID: accommodation.HostID, Name: accommodation.Name,
					Location: accommodation.Location, Wifi: accommodation.Wifi, Kitchen: accommodation.Kitchen, AC: accommodation.AC, ParkingLot: accommodation.ParkingLot, Images: accommodation.Images,
					StartDate: avail.StartDate, EndDate: avail.EndDate, TotalPrice: totalPrice, SinglePrice: avail.Price, IsPricePerGuest: avail.IsPricePerGuest}
				favailabilities = append(favailabilities, findAvailability)
			}
		} else {
			i := 0
			for _, res := range reservations.Reservations {
				yearRes, monthRes, dayRes := res.StartDate.AsTime().Date()
				yearERes, monthERes, dayERes := res.EndDate.AsTime().Date()
				startDateRes := time.Date(yearRes, monthRes, dayRes, int(0), int(0), int(0), int(0), time.UTC)
				endDateRes := time.Date(yearERes, monthERes, dayERes, int(0), int(0), int(0), int(0), time.UTC)
				if res.IsAccepted && !res.IsCanceled && !res.IsDeleted && ((!startDate.Before(startDateRes) && !startDate.After(endDateRes)) || (!endDate.After(endDateRes) && !endDate.Before(startDateRes)) || (!startDate.After(startDateRes) && !endDate.Before(endDateRes))) {
					i++
				}
			}
			if i == 0 {
				if strings.Title(strings.ToLower(accommodation.Location)) == strings.Title(strings.ToLower(request.Location)) && int(request.GuestsNum) >= accommodation.MinGuests && int(request.GuestsNum) <= accommodation.MaxGuests {
					duration := endDate.Sub(startDate)
					nights := int(duration.Hours() / 24)
					totalPrice := avail.Price * float64(nights)
					findAvailability := models.FindAvailability{AccommodationId: accommodation.ID, AvailabilityID: avail.ID, HostID: accommodation.HostID, Name: accommodation.Name,
						Location: accommodation.Location, Wifi: accommodation.Wifi, Kitchen: accommodation.Kitchen, AC: accommodation.AC, ParkingLot: accommodation.ParkingLot, Images: accommodation.Images,
						StartDate: avail.StartDate, EndDate: avail.EndDate, TotalPrice: totalPrice, SinglePrice: avail.Price, IsPricePerGuest: avail.IsPricePerGuest}
					favailabilities = append(favailabilities, findAvailability)
				} else {
					return nil, status.Errorf(codes.InvalidArgument, "1All accommodations are occupied at chosen time!")
				}
			}
		}
	}
	if len(favailabilities) == 0 {
		return nil, status.Errorf(codes.NotFound, "No accommodations!")
	}
	findAvailabilities := []*pb.FindAvailability{}
	for _, a := range favailabilities {
		findAvailabilitiyPb := mapFindAvailability(&a)
		findAvailabilities = append(findAvailabilities, findAvailabilitiyPb)
	}
	response := &pb.SearchAvailabilityResponse{
		FindAvailability: findAvailabilities,
	}
	return response, nil
}

/*
func (handler *AvailabilityHandler) FilterAvailability(ctx context.Context, request *pb.FilterAvailabilityRequest) (*pb.FilterAvailabilityResponse, error) {
	//filtriranje smjestaja zahtjev 1.14 mozete iskoristit search metodu ali sa dodanim poljima, prosiren je request mozete vidjeti u proto
	//na frontu mozete zamjeniti da se ne koristi vise SearchAvailabilty nego ovo ili da imate dvije odvojene metode kako vam je lakse
	response := &pb.FilterAvailabilityResponse{
		FindAvailability: findAvailabilities,
	}
	return response, nil
}*/
