package api

import (
	"context"

	//"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/service"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AvailabilityHandler struct {
	pb.UnimplementedAccommodationServiceServer
	accommodation_service *service.AccommodationService
	availability_service  *service.AvailabilityService
	user_client           user_service.UserServiceClient
	reservation_client    reservation_service.ReservationServiceClient
}

func NewAvailabilityHandler(accommodation_service *service.AccommodationService, availability_service *service.AvailabilityService, user_client user_service.UserServiceClient, reservation_client reservation_service.ReservationServiceClient) *AvailabilityHandler {
	return &AvailabilityHandler{
		accommodation_service: accommodation_service,
		availability_service:  availability_service,
		user_client:           user_client,
		reservation_client:    reservation_client,
	}
}

func (handler *AvailabilityHandler) GetAllAvailabilities(ctx context.Context, request *pb.GetAllAvailabilitiesRequest) (*pb.GetAllAvailabilitiesResponse, error) {
	//TODO pomocna metoda za dobavljanje svih dostupnosti koje mozete koristiti u drugim mikroservisima
	id := request.Id
	accomodationId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	as, err := handler.availability_service.GetAllAvailabilitiesByAccommodationID(accomodationId)
	if err != nil {
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
	//TODO
	response := &pb.GetAvailabilityByIdResponse{
		Availability: nil,
	}
	return response, nil
}

func (handler *AvailabilityHandler) CreateAvailability(ctx context.Context, request *pb.CreateAvailabilityRequest) (*pb.CreateAvailabilityResponse, error) {
	//TODO
	accommodationID, err := primitive.ObjectIDFromHex(request.AccommodationID)
	if err != nil {
		return nil, err
	}

	availability := models.Availability{AccommodationID: accommodationID, StartDate: request.EndDate.AsTime(), EndDate: request.EndDate.AsTime(), Price: request.Price, IsPricePerGuest: request.IsPricePerGuest}
	mess, err := handler.availability_service.CreateAvailability(availability)
	if err != nil {
		err := status.Errorf(codes.Internal, mess)
		return nil, err
	}
	response := &pb.CreateAvailabilityResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *AvailabilityHandler) UpdateAvailability(ctx context.Context, request *pb.UpdateAvailabilityRequest) (*pb.UpdateAvailabilityResponse, error) {
	//TODO treba dobaviti sve rezervacije i provjeriti da li postoje neke za availiabilty koji treba da se mjenja a da je isdeleted na false
	//samo ako nema onda moze da se azurira
	//ako imamo rezervacije ne mozemo da menjamo tj. ukoliko i ima rezervacija, ali ako je isdeleted na true ili iscanceled na true onda moze da se menja (obrisane su)
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
