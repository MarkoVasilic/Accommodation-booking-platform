package api

import (
	"context"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/reservation_service/service"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
)

type ReservationHandler struct {
	pb.UnimplementedReservationServiceServer
	reservation_service *service.ReservationService
}

func NewReservationHandler(reservation_service *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		reservation_service: reservation_service,
	}
}

func (handler *ReservationHandler) GetAllReservations(ctx context.Context, request *pb.GetAllReservationsRequest) (*pb.GetAllReservationsResponse, error) {
	//TODO pomocna metoda za dobavljanje svih rezervacija koje mozete koristiti u drugim mikroservisima
	reservations := []*pb.Reservation{}
	response := &pb.GetAllReservationsResponse{
		Reservations: reservations,
	}
	return response, nil
}

func (handler *ReservationHandler) CreateReservation(ctx context.Context, request *pb.CreateReservationRequest) (*pb.CreateReservationResponse, error) {
	//TODO paziti da li je automatsko prihvatanje
	//to jeste da li postoji rezervacija u preklapajucem intervalu sa isaccepted na true a da su isdeleted/iscanceled na false
	//ako jeste odbiti to jest ne praviti rezervaciju
	response := &pb.CreateReservationResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *ReservationHandler) GetFindReservationPendingGuest(ctx context.Context, request *pb.GetFindReservationPendingGuestRequest) (*pb.GetFindReservationPendingGuestResponse, error) {
	//TODO get zahtjev, dobija se id guesta, i naci sve njegove rezervacije za koje je isdeleted,isaccepted i iscanceled na false
	//treba napraviti mapper koji mapira na pb u obliku dto koji sam napravio i to treba da je lista za svaku rezervaciju
	//polje u dto NumOfCancellation samo postaviti na 0
	findReservation := []*pb.FindReservation{}
	response := &pb.GetFindReservationPendingGuestResponse{
		FindReservation: findReservation,
	}
	return response, nil
}

func (handler *ReservationHandler) GetFindReservationAcceptedGuest(ctx context.Context, request *pb.GetFindReservationAcceptedGuestRequest) (*pb.GetFindReservationAcceptedGuestResponse, error) {
	//TODO get zahtjev, dobija se id guesta, i naci sve njegove rezervacije za koje je isdeleted i iscanceled na false, a isaccepted na true
	//treba napraviti mapper koji mapira na pb u obliku dto koji sam napravio i to treba da je lista za svaku rezervaciju
	//polje u dto NumOfCancellation samo postaviti na 0
	findReservation := []*pb.FindReservation{}
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
	response := &pb.CancelReservationResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *ReservationHandler) DeleteLogicallyReservation(ctx context.Context, request *pb.DeleteLogicallyReservationRequest) (*pb.DeleteLogicallyReservationResponse, error) {
	//TODO postaviti samo isdelted na true za rezervaciju za koju je dobijen id, ovo se poziva kad guest zeli da obrise rezervaciju
	//i kad host zeli da odbije rezervaciju
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
