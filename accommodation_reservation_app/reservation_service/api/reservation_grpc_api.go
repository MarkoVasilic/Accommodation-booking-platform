package api

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ReservationHandler struct {
	pb.UnimplementedReservationServiceServer
	reservation_service  *service.ReservationService
	accommodation_client accommodation_service.AccommodationServiceClient
	user_client          user_service.UserServiceClient
}

func NewReservationHandler(reservation_service *service.ReservationService, accommodation_client accommodation_service.AccommodationServiceClient, user_client user_service.UserServiceClient) *ReservationHandler {
	return &ReservationHandler{
		reservation_service:  reservation_service,
		accommodation_client: accommodation_client,
		user_client:          user_client,
	}
}

func createContextForAuthorization(ctx context.Context) context.Context {
	token, _ := grpc_auth.AuthFromMD(ctx, "Bearer")
	if len(token) > 0 {
		return metadata.NewOutgoingContext(context.Background(), metadata.Pairs("Authorization", "Bearer "+token))
	}
	return context.TODO()
}

// by availability
func (handler *ReservationHandler) GetAllReservations(ctx context.Context, request *pb.GetAllReservationsRequest) (*pb.GetAllReservationsResponse, error) {
	//TODO pomocna metoda za dobavljanje svih rezervacija koje mozete koristiti u drugim mikroservisima
	id := request.Id
	availabilityId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	res, err := handler.reservation_service.GetAllReservations(availabilityId)
	if err != nil {
		return nil, err
	}
	reservations := []*pb.Reservation{}
	for _, r := range res {
		reservationsPb := mapReservation(&r)
		reservations = append(reservations, reservationsPb)
	}

	response := &pb.GetAllReservationsResponse{
		Reservations: reservations,
	}
	return response, nil
}

func (handler *ReservationHandler) CreateReservation(ctx context.Context, request *pb.CreateReservationRequest) (*pb.CreateReservationResponse, error) {
	year, month, day := request.StartDate.AsTime().Date()
	yearE, monthE, dayE := request.EndDate.AsTime().Date()
	startDate := time.Date(year, month, day, int(0), int(0), int(0), int(0), time.UTC)
	endDate := time.Date(yearE, monthE, dayE, int(0), int(0), int(0), int(0), time.UTC)
	availabilityId, err := primitive.ObjectIDFromHex(request.AvailabilityID)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	guestId, err := primitive.ObjectIDFromHex(request.GuestId)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	reservation := models.Reservation{AvailabilityID: availabilityId, GuestID: guestId,
		StartDate: startDate, EndDate: endDate, NumGuests: int(request.NumGuests), IsAccepted: false, IsCanceled: false, IsDeleted: false}
	mess, err := handler.reservation_service.CreateReservation(reservation)
	if err != nil {
		err := status.Errorf(codes.Internal, mess)
		return nil, err
	}

	//automatska potvrda rezervacije
	accommodation, err1 := handler.accommodation_client.GetAccommodationByAvailability(createContextForAuthorization(ctx), &accommodation_service.GetAccommodationByAvailabilityRequest{Id: request.AvailabilityID})
	if err1 != nil {
		err1 := status.Errorf(codes.Internal, mess)
		return nil, err1
	}

	if accommodation.Accommodation.AutoAccept {
		reservationId, err := primitive.ObjectIDFromHex(mess)
		if err != nil {
			err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
			return nil, err
		}

		mes, err := handler.reservation_service.AcceptReservation(reservationId)
		if err != nil {
			err := status.Errorf(codes.Internal, mes)
			return nil, err
		}
	}
	////////////

	response := &pb.CreateReservationResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *ReservationHandler) GetFindReservationPendingGuest(ctx context.Context, request *pb.GetFindReservationPendingGuestRequest) (*pb.GetFindReservationPendingGuestResponse, error) {
	guestId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	res, err := handler.reservation_service.GetFindReservationPendingGuest(guestId)
	if err != nil {
		return nil, err
	} else if res == nil {
		err := status.Errorf(codes.InvalidArgument, "There is no pending reservations!")
		return nil, err
	}
	var filteredReservations []models.FindReservation
	for _, reservation := range res {
		availabilityId := string(reservation.AvailabilityID.Hex())
		accommodation, err := handler.accommodation_client.GetAccommodationByAvailability(createContextForAuthorization(ctx), &accommodation_service.GetAccommodationByAvailabilityRequest{Id: availabilityId})
		if err != nil {
			return nil, err
		}
		findRes := models.FindReservation{ReservationId: reservation.ID, GuestID: reservation.GuestID, Name: accommodation.Accommodation.Name, Location: accommodation.Accommodation.Location, StartDate: reservation.StartDate, EndDate: reservation.EndDate, NumOfCancelation: 0, IsAccepted: reservation.IsAccepted, IsCanceled: reservation.IsCanceled}
		filteredReservations = append(filteredReservations, findRes)
	}

	findReservation := []*pb.FindReservation{}
	for _, r := range filteredReservations {
		reservationsPb := mapFindReservation(&r)
		findReservation = append(findReservation, reservationsPb)
	}
	response := &pb.GetFindReservationPendingGuestResponse{
		FindReservation: findReservation,
	}

	return response, nil
}

func (handler *ReservationHandler) GetFindReservationAcceptedGuest(ctx context.Context, request *pb.GetFindReservationAcceptedGuestRequest) (*pb.GetFindReservationAcceptedGuestResponse, error) {
	guestId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	res, err := handler.reservation_service.GetFindReservationAcceptedGuest(guestId)
	if err != nil {
		return nil, err
	} else if res == nil {
		err := status.Errorf(codes.InvalidArgument, "There is no accepted reservations!")
		return nil, err
	}
	var filteredReservations []models.FindReservation
	for _, reservation := range res {
		availabilityId := string(reservation.AvailabilityID.Hex())
		accommodation, err := handler.accommodation_client.GetAccommodationByAvailability(createContextForAuthorization(ctx), &accommodation_service.GetAccommodationByAvailabilityRequest{Id: availabilityId})
		if err != nil {
			return nil, err
		}
		findRes := models.FindReservation{ReservationId: reservation.ID, GuestID: reservation.GuestID, Name: accommodation.Accommodation.Name, Location: accommodation.Accommodation.Location, StartDate: reservation.StartDate, EndDate: reservation.EndDate, NumOfCancelation: 0, IsAccepted: reservation.IsAccepted, IsCanceled: reservation.IsCanceled}
		filteredReservations = append(filteredReservations, findRes)
	}

	findReservation := []*pb.FindReservation{}
	for _, r := range filteredReservations {
		reservationsPb := mapFindReservation(&r)
		findReservation = append(findReservation, reservationsPb)
	}
	response := &pb.GetFindReservationAcceptedGuestResponse{
		FindReservation: findReservation,
	}
	return response, nil
}

func (handler *ReservationHandler) GetFindReservationHost(ctx context.Context, request *pb.GetFindReservationHostRequest) (*pb.GetFindReservationHostResponse, error) {
	allAccommodations, err := handler.accommodation_client.GetAllAccommodations(createContextForAuthorization(ctx), &accommodation_service.GetAllAccommodationsRequest{Id: "64580a2e9f857372a34602c2"})
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}

	var hostAccomodationsIds []string
	for _, accommodation := range allAccommodations.Accommodations {
		if accommodation.HostId == request.Id {
			hostAccomodationsIds = append(hostAccomodationsIds, accommodation.Id)
		}
	}

	allAvailabilities, err := handler.accommodation_client.GetAllAvailabilities(createContextForAuthorization(ctx), &accommodation_service.GetAllAvailabilitiesRequest{Id: "64580a2e9f857372a34602c2"})
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}

	var hostAvailabilitiesIds []primitive.ObjectID
	for _, availability := range allAvailabilities.Availabilities {
		for _, accommodationId := range hostAccomodationsIds {
			if availability.AccommodationID == accommodationId {
				id, err := primitive.ObjectIDFromHex(availability.Id)
				if err != nil {
					err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
					return nil, err
				}
				hostAvailabilitiesIds = append(hostAvailabilitiesIds, id)
			}
		}
	}

	var hostPendingReservations []models.Reservation
	for _, availabilityId := range hostAvailabilitiesIds {
		res, err := handler.reservation_service.GetAll()
		if err != nil {
			return nil, err
		}

		for _, r := range res {
			if r.IsCanceled == false && r.IsDeleted == false && r.IsAccepted == false && r.AvailabilityID == availabilityId {
				hostPendingReservations = append(hostPendingReservations, r)
			}
		}
	}

	var filteredReservations []models.FindReservation
	for _, reservation := range hostPendingReservations {
		user, err := handler.user_client.GetUser(createContextForAuthorization(ctx), &user_service.GetUserRequest{Id: request.Id})
		if err != nil {
			return nil, err
		}

		availabilityId := string(reservation.AvailabilityID.Hex())
		accommodation, err := handler.accommodation_client.GetAccommodationByAvailability(createContextForAuthorization(ctx), &accommodation_service.GetAccommodationByAvailabilityRequest{Id: availabilityId})
		if err != nil {
			return nil, err
		}
		canceledReservations, err := handler.reservation_service.GetAllCanceledReservationsByGuest(reservation.GuestID)

		findRes := models.FindReservation{
			ReservationId:    reservation.ID,
			GuestID:          reservation.GuestID,
			Name:             user.User.FirstName + " " + user.User.LastName,
			Location:         accommodation.Accommodation.Location,
			StartDate:        reservation.StartDate,
			EndDate:          reservation.EndDate,
			NumOfCancelation: len(canceledReservations),
			IsAccepted:       accommodation.Accommodation.AutoAccept,
			IsCanceled:       reservation.IsCanceled}
		filteredReservations = append(filteredReservations, findRes)
	}

	findReservation := []*pb.FindReservation{}
	for _, r := range filteredReservations {
		reservationsPb := mapFindReservation(&r)
		findReservation = append(findReservation, reservationsPb)
	}
	response := &pb.GetFindReservationHostResponse{
		FindReservation: findReservation,
	}

	return response, nil
}

func (handler *ReservationHandler) CancelReservation(ctx context.Context, request *pb.CancelReservationRequest) (*pb.CancelReservationResponse, error) {
	println("METHODD")
	Id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	mess, err := handler.reservation_service.CancelReservation(Id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, mess)
		return nil, err
	}
	response := &pb.CancelReservationResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *ReservationHandler) DeleteLogicallyReservation(ctx context.Context, request *pb.DeleteLogicallyReservationRequest) (*pb.DeleteLogicallyReservationResponse, error) {
	Id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	mess, err := handler.reservation_service.DeleteLogicallyReservation(Id)
	if err != nil {
		err := status.Errorf(codes.Internal, mess)
		return nil, err
	}
	response := &pb.DeleteLogicallyReservationResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *ReservationHandler) AcceptReservation(ctx context.Context, request *pb.AcceptReservationRequest) (*pb.AcceptReservationResponse, error) {
	Id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	mess, err := handler.reservation_service.AcceptReservation(Id)
	if err != nil {
		err := status.Errorf(codes.Internal, mess)
		return nil, err
	}

	response := &pb.AcceptReservationResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *ReservationHandler) GetAllReservationsHost(ctx context.Context, request *pb.GetAllReservationsHostRequest) (*pb.GetAllReservationsHostResponse, error) {

	allAccommodations, err := handler.accommodation_client.GetAllAccommodations(createContextForAuthorization(ctx), &accommodation_service.GetAllAccommodationsRequest{Id: "64580a2e9f857372a34602c2"})
	if err != nil {
		return nil, err
	}

	var hostAccomodationsIds []string
	for _, accommodation := range allAccommodations.Accommodations {
		if accommodation.HostId == request.Id {
			hostAccomodationsIds = append(hostAccomodationsIds, accommodation.Id)
		}
	}

	allAvailabilities, err := handler.accommodation_client.GetAllAvailabilities(createContextForAuthorization(ctx), &accommodation_service.GetAllAvailabilitiesRequest{Id: "64580a2e9f857372a34602c2"})
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}

	var hostAvailabilitiesIds []primitive.ObjectID
	for _, availability := range allAvailabilities.Availabilities {
		for _, accommodationId := range hostAccomodationsIds {
			if availability.AccommodationID == accommodationId {
				id, err := primitive.ObjectIDFromHex(availability.Id)
				if err != nil {
					err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
					return nil, err
				}
				hostAvailabilitiesIds = append(hostAvailabilitiesIds, id)
			}
		}
	}

	var hostReservations []models.Reservation
	for _, availabilityId := range hostAvailabilitiesIds {
		res, err := handler.reservation_service.GetAll()
		if err != nil {
			return nil, err
		}

		for _, r := range res {
			if r.IsCanceled == false && r.IsDeleted == false && r.IsAccepted == true && r.AvailabilityID == availabilityId && r.StartDate.After(time.Now()) {
				hostReservations = append(hostReservations, r)
			}
		}
	}

	hostReservationsMap := []*pb.Reservation{}
	for _, r := range hostReservations {
		reservationsPb := mapReservation(&r)
		hostReservationsMap = append(hostReservationsMap, reservationsPb)
	}
	response := &pb.GetAllReservationsHostResponse{
		Reservation: hostReservationsMap,
	}
	return response, nil
}

func (handler *ReservationHandler) DeleteReservationsHost(ctx context.Context, request *pb.DeleteReservationsHostRequest) (*pb.DeleteReservationsHostResponse, error) {
	allAccommodations, err := handler.accommodation_client.GetAllAccommodations(createContextForAuthorization(ctx), &accommodation_service.GetAllAccommodationsRequest{Id: request.Id})
	if err != nil {
		return nil, err
	}

	var hostAccomodationsIds []string
	for _, accommodation := range allAccommodations.Accommodations {
		if accommodation.HostId == request.Id {
			hostAccomodationsIds = append(hostAccomodationsIds, accommodation.Id)
		}
	}

	allAvailabilities, err := handler.accommodation_client.GetAllAvailabilities(createContextForAuthorization(ctx), &accommodation_service.GetAllAvailabilitiesRequest{Id: "64580a2e9f857372a34602c2"})
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}

	var hostAvailabilitiesIds []primitive.ObjectID
	for _, availability := range allAvailabilities.Availabilities {
		for _, accommodationId := range hostAccomodationsIds {
			if availability.AccommodationID == accommodationId {
				id, err := primitive.ObjectIDFromHex(availability.Id)
				if err != nil {
					err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
					return nil, err
				}
				hostAvailabilitiesIds = append(hostAvailabilitiesIds, id)
			}
		}
	}

	var hostReservations []models.Reservation
	for _, availabilityId := range hostAvailabilitiesIds {
		res, err := handler.reservation_service.GetAll()
		if err != nil {
			return nil, err
		}

		for _, r := range res {
			if r.AvailabilityID == availabilityId {
				hostReservations = append(hostReservations, r)
			}
		}
	}
	resp, err := handler.reservation_service.DeleteReservationsHost(hostReservations)
	if err != nil {
		err := status.Errorf(codes.Internal, "something went wrong")
		return nil, err
	}
	response := &pb.DeleteReservationsHostResponse{
		Message: resp,
	}
	return response, nil
}
