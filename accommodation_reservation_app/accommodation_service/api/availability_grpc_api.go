package api

import (
	"context"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/service"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
)

type AvailabilityHandler struct {
	pb.UnimplementedAccommodationServiceServer
	accommodation_service *service.AccommodationService
	availability_service  *service.AvailabilityService
}

func NewAvailabilityHandler(accommodation_service *service.AccommodationService, availability_service *service.AvailabilityService) *AvailabilityHandler {
	return &AvailabilityHandler{
		accommodation_service: accommodation_service,
		availability_service:  availability_service,
	}
}

func (handler *AvailabilityHandler) GetAllAvailabilities(ctx context.Context, request *pb.GetAllAvailabilitiesRequest) (*pb.GetAllAvailabilitiesResponse, error) {
	//TODO pomocna metoda za dobavljanje svih dostupnosti koje mozete koristiti u drugim mikroservisima
	availabilities := []*pb.Availability{}
	response := &pb.GetAllAvailabilitiesResponse{
		Availabilities: availabilities,
	}
	return response, nil
}

func (handler *AvailabilityHandler) CreateAvailability(ctx context.Context, request *pb.CreateAvailabilityRequest) (*pb.CreateAvailabilityResponse, error) {
	//TODO
	response := &pb.CreateAvailabilityResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *AvailabilityHandler) UpdateAvailability(ctx context.Context, request *pb.UpdateAvailabilityRequest) (*pb.UpdateAvailabilityResponse, error) {
	//TODO treba dobaviti sve rezervacije i provjeriti da li postoje neke za availiabilty koji treba da se mjenja a da je isdeleted na false
	//samo ako nema onda moze da se azurira
	response := &pb.UpdateAvailabilityResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *AvailabilityHandler) SearchAvailability(ctx context.Context, request *pb.SearchAvailabilityRequest) (*pb.SearchAvailabilityResponse, error) {
	//TODO, napraviti mapper, metoda je post,
	//naci prvo sve availiabilty koji zadovoljavaju filtere
	//Provjere:
	// da li postoje rezervacije za isti smjestaj u preklapajucem vremenu
	// ako da provjeriti da li su rezervisana odnosno da li je polje isaccepted na true, onda iskljuciti, a ako je i iscanceled ili isdelted na true onda ukljuciti
	//Metoda je post, ali moze da se radi i sa query parametrima
	//Treba napraviti mapper koji mapira dto u pb i pravi listu tih koji ce biti vraceni
	//
	//na frontu ce vjerovatno trebati dvije stranice jedna za guestovi i jedna za neulogovane usere, jer oni ne mogu da rezervisu
	findAvailabilities := []*pb.FindAvailability{}
	response := &pb.SearchAvailabilityResponse{
		FindAvailability: findAvailabilities,
	}
	return response, nil
}
