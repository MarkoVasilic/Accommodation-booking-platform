package api

import (
	"context"

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
	//TODO paziti da li je automatsko prihvatanje
	//to jeste da li postoji rezervacija u preklapajucem intervalu sa isaccepted na true a da su isdeleted/iscanceled na false
	//ako jeste odbiti to jest ne praviti rezervaciju
	Id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	availabilityId, err := primitive.ObjectIDFromHex(request.GuestId)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	guestId, err := primitive.ObjectIDFromHex(request.AvailabilityID)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	reservation := models.Reservation{ID: Id, AvailabilityID: availabilityId, GuestID: guestId,
		StartDate: request.StartDate.AsTime(), EndDate: request.EndDate.AsTime(), NumGuests: int(request.NumGuests), IsAccepted: false, IsCanceled: false, IsDeleted: false}
	mess, err := handler.reservation_service.CreateReservation(reservation)
	if err != nil {
		err := status.Errorf(codes.Internal, mess)
		return nil, err
	}
	response := &pb.CreateReservationResponse{
		Message: "Success",
	}
	return response, nil
}

// dodati location
func (handler *ReservationHandler) GetFindReservationPendingGuest(ctx context.Context, request *pb.GetFindReservationPendingGuestRequest) (*pb.GetFindReservationPendingGuestResponse, error) {
	//TODO get zahtjev, dobija se id guesta, i naci sve njegove rezervacije za koje je isdeleted,isaccepted i iscanceled na false
	//treba napraviti mapper koji mapira na pb u obliku dto koji sam napravio i to treba da je lista za svaku rezervaciju
	//polje u dto NumOfCancellation samo postaviti na 0
	guestId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	res, err := handler.reservation_service.GetFindReservationPendingGuest(guestId)
	if err != nil {
		return nil, err
	}
	var filteredReservations []models.FindReservation
	for _, reservation := range res {
		user, err := handler.user_client.GetUser(createContextForAuthorization(ctx), &user_service.GetUserRequest{Id: request.Id})
		if err != nil {
			return nil, err
		}
		//accomodation, err := handler.accommodation_client.GetAccomodationByAvailiabilityId(reservation.availabilityId)
		findRes := models.FindReservation{ReservationId: reservation.ID, GuestID: reservation.GuestID, Name: user.User.FirstName + " " + user.User.LastName, Location: "", StartDate: reservation.StartDate, EndDate: reservation.EndDate, NumOfCancelation: 0, IsAccepted: reservation.IsAccepted, IsCanceled: reservation.IsCanceled}
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
	//TODO get zahtjev, dobija se id guesta, i naci sve njegove rezervacije za koje je isdeleted i iscanceled na false, a isaccepted na true
	//treba napraviti mapper koji mapira na pb u obliku dto koji sam napravio i to treba da je lista za svaku rezervaciju
	//polje u dto NumOfCancellation samo postaviti na 0
	guestId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	res, err := handler.reservation_service.GetFindReservationAcceptedGuest(guestId)
	if err != nil {
		return nil, err
	}
	var filteredReservations []models.FindReservation
	for _, reservation := range res {
		user, err := handler.user_client.GetUser(createContextForAuthorization(ctx), &user_service.GetUserRequest{Id: string(guestId.Hex())})
		if err != nil {
			return nil, err
		}
		//accomodation, err := handler.accommodation_client.GetAccomodationByAvailiabilityId(reservation.availabilityId)
		findRes := models.FindReservation{ReservationId: reservation.ID, GuestID: reservation.GuestID, Name: user.User.FirstName + " " + user.User.LastName, Location: "", StartDate: reservation.StartDate, EndDate: reservation.EndDate, NumOfCancelation: 0, IsAccepted: reservation.IsAccepted, IsCanceled: reservation.IsCanceled}
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
	//TODO get zahtjev, dobija se id hosta, i naci sve rezervacije koje su vezane za njegov smjestaj(dobaviti sve rezervacije, izvuci availibilty id iz svake pa dobaviti accommodation za svaku i uzeti samo one gdje je id poslan = id hosta u accommodation)
	//dodatno izvuci samo one gdje je isdeleted na false(dodatno i one gdje je iscanceled na false i isaccepted na false, ali to bi mogli provjeriti sa asistentom)
	//treba napraviti mapper koji mapira na pb u obliku dto koji sam napravio i to treba da je lista za svaku rezervaciju
	//za svaku konacnu rezervaciju izvuci listu guestova i za svakog pronaci njihove rezervacije i provjeriti koliko puta ima iscancelled i to sacuvati u polje NumOfCancellation u svakom dto
	findReservation := []*pb.FindReservation{}
	response := &pb.GetFindReservationHostResponse{
		FindReservation: findReservation,
	}
	return response, nil
}

func (handler *ReservationHandler) CancelReservation(ctx context.Context, request *pb.CancelReservationRequest) (*pb.CancelReservationResponse, error) {
	//TODO dobija se id rezervacije provjeriti da li je datum zahtjeva u dozvoljenom vremenu otkazivanja ako jeste
	//promjeniti samo iscanceled na true
	Id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	mess, err := handler.reservation_service.CancelReservation(Id)
	if err != nil {
		err := status.Errorf(codes.Internal, mess)
		return nil, err
	}
	response := &pb.CancelReservationResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *ReservationHandler) DeleteLogicallyReservation(ctx context.Context, request *pb.DeleteLogicallyReservationRequest) (*pb.DeleteLogicallyReservationResponse, error) {
	//TODO postaviti samo isdelted na true za rezervaciju za koju je dobijen id, ovo se poziva kad guest zeli da obrise rezervaciju
	//i kad host zeli da odbije rezervaciju
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
	//TODO postaviti isaccepted na true za rezervaciju za koju je id poslan
	//dodatno treba dobaviti sve rezervacije i naci one sa preklopljenim vremenima za isti smjestaj i njima staviti isdeleted na true
	response := &pb.AcceptReservationResponse{
		Message: "Success",
	}
	return response, nil
}
