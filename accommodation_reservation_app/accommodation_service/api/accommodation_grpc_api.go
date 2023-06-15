package api

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/token"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccommodationHandler struct {
	pb.UnimplementedAccommodationServiceServer
	accommodation_service *service.AccommodationService
	availability_service  *service.AvailabilityService
	grade_service         *service.GradeService
	user_client           user_service.UserServiceClient
	reservation_client    reservation_service.ReservationServiceClient
}

func NewAccommodationHandler(accommodation_service *service.AccommodationService, availability_service *service.AvailabilityService, grade_service *service.GradeService, user_client user_service.UserServiceClient, reservation_client reservation_service.ReservationServiceClient) *AccommodationHandler {
	return &AccommodationHandler{
		accommodation_service: accommodation_service,
		availability_service:  availability_service,
		grade_service:         grade_service,
		user_client:           user_client,
		reservation_client:    reservation_client,
	}
}

func (handler *AccommodationHandler) GetAllAccommodations(ctx context.Context, request *pb.GetAllAccommodationsRequest) (*pb.GetAllAccommodationsResponse, error) {
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

func (handler *AccommodationHandler) GetAccommodationById(ctx context.Context, request *pb.GetAccommodationByIdRequest) (*pb.GetAccommodationByIdResponse, error) {
	id := request.Id
	accommodationId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid AccommodationID")
		return nil, err
	}
	accommodation, err := handler.accommodation_service.GetAccommodationById(accommodationId)
	if err != nil {
		return nil, err
	}
	accommodationPb := mapAccommodation(&accommodation)
	response := &pb.GetAccommodationByIdResponse{
		Accommodation: accommodationPb,
	}
	return response, nil
}

func (handler *AccommodationHandler) GetAccommodationByAvailability(ctx context.Context, request *pb.GetAccommodationByAvailabilityRequest) (*pb.GetAccommodationByAvailabilityResponse, error) {
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

func (handler *AccommodationHandler) GetAllAccommodationGuestGrades(ctx context.Context, request *pb.GetAllAccommodationGuestGradesRequest) (*pb.GetAllAccommodationGuestGradesResponse, error) {
	//TODO pomocna metoda za dobavljanje svih ocijena guesta za poslani id guesta
	//a vraca se lista dtova koji sam napravio
	id := request.Id
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided guestId is not a valid ObjectID")
		return nil, err
	}
	res, err := handler.grade_service.GetAllAccommodationGuestGrades(request.Id)
	if err != nil {
		return nil, err
	}
	user, err := handler.user_client.GetUser(createContextForAuthorization(ctx), &user_service.GetUserRequest{Id: request.Id})
	if err != nil {
		return nil, err
	}
	var gradeDTOs []models.AccommodationGradeDetails
	for _, grade := range res {
		accomodation, err := handler.accommodation_service.GetAccommodationById(grade.AccommodationID)
		if err != nil {
			return nil, err
		}
		gradeDTO := models.AccommodationGradeDetails{GuestFirstName: *&user.User.FirstName, GuestLastName: *&user.User.LastName, AccommodationName: accomodation.Name, Grade: grade.Grade, DateOfGrade: grade.DateOfGrade}
		gradeDTOs = append(gradeDTOs, gradeDTO)
	}
	gradesDetails := []*pb.AccommodationGradeDetails{}
	for _, r := range gradeDTOs {
		gradesPb := mapAccommodationGradeDetails(&r)
		gradesDetails = append(gradesDetails, gradesPb)
	}
	response := &pb.GetAllAccommodationGuestGradesResponse{
		AccommodationGradeDetails: gradesDetails,
	}
	return response, nil
}

func (handler *AccommodationHandler) GetEveryAccommodation(ctx context.Context, request *pb.GetEveryAccommodationRequest) (*pb.GetEveryAccommodationResponse, error) {
	//TODO pomocna metoda za dobavljanje svih smjestaja, ne salje se nista
	//a vraca se lista smjestaja
	ClientToken, _ := grpc_auth.AuthFromMD(ctx, "Bearer")
	claims, _ := token.ValidateToken(ClientToken)
	reservations, err := handler.reservation_client.GetAllReservations(createContextForAuthorization(ctx), &reservation_service.GetAllReservationsRequest{Id: claims.Uid})
	if err != nil {
		return nil, err
	}
	accommodations := []models.Accommodation{}
	for _, res := range reservations.Reservations {
		if res.IsCanceled || res.IsDeleted {
			continue
		}
		availabilityId, err := primitive.ObjectIDFromHex(res.AvailabilityID)
		if err != nil {
			err := status.Errorf(codes.InvalidArgument, "the provided availabilityId is not a valid ObjectID")
			return nil, err
		}
		availabilitiy, err := handler.availability_service.GetAvailabilityById(availabilityId)
		accommodation, err := handler.accommodation_service.GetAccommodationById(availabilitiy.AccommodationID)
		accommodations = append(accommodations, accommodation)
	}

	if accommodations == nil {
		err := status.Errorf(codes.InvalidArgument, "There is no accommodations to be graded!")
		return nil, err
	}
	accommodationsPb := []*pb.Accommodation{}
	for _, a := range accommodations {
		accommodationPb := mapAccommodation(&a)
		accommodationsPb = append(accommodationsPb, accommodationPb)
	}
	response := &pb.GetEveryAccommodationResponse{
		Accommodations: accommodationsPb,
	}
	return response, nil
}

func (handler *AccommodationHandler) CreateAccommodationGrade(ctx context.Context, request *pb.CreateAccommodationGradeRequest) (*pb.CreateAccommodationGradeResponse, error) {
	//TODO zahtjev 1.12 kreiranje ocijene
	//ClientToken, _ := grpc_auth.AuthFromMD(ctx, "Bearer")
	//claims, _ := token.ValidateToken(ClientToken)
	//ovako se izvlaci id osobe koja salje zahtjev, id se nalazi u claims.Uid
	//provjeriti uslov da li moze da ga ocijeni
	guestId, err := primitive.ObjectIDFromHex(request.AccommodationGrade.GuestID)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided guestId is not a valid ObjectID")
		return nil, err
	}
	accommodationId, err := primitive.ObjectIDFromHex(request.AccommodationGrade.AccommodationID)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided accommodationId is not a valid ObjectID")
		return nil, err
	}
	reservations, err := handler.reservation_client.GetAllReservations(createContextForAuthorization(ctx), &reservation_service.GetAllReservationsRequest{Id: request.AccommodationGrade.GuestID})
	if err != nil {
		return nil, err
	}
	if len(reservations.Reservations) == 0 {
		err := status.Errorf(codes.InvalidArgument, "You cannot grade accommodation if you have no reservations!")
		return nil, err
	}
	numOfGuestReservations := 0
	for _, res := range reservations.Reservations {
		if res.IsCanceled || res.IsDeleted {
			continue //
		}
		availabilityId, err := primitive.ObjectIDFromHex(res.AvailabilityID)
		if err != nil {
			err := status.Errorf(codes.InvalidArgument, "the provided availabilityId is not a valid ObjectID")
			return nil, err
		}
		availabilitiy, err := handler.availability_service.GetAvailabilityById(availabilityId)
		if availabilitiy.AccommodationID == accommodationId {
			numOfGuestReservations = numOfGuestReservations + 1
		}
		if numOfGuestReservations >= 1 {
			break
		}
	}
	if numOfGuestReservations < 1 {
		err := status.Errorf(codes.InvalidArgument, "You cannot grade accommodation if you never stayed there!")
		return nil, err
	}

	grade := int(request.AccommodationGrade.Grade)
	accommodationGrade := models.AccommodationGrade{GuestID: guestId, AccommodationID: accommodationId, Grade: grade, DateOfGrade: time.Now()}
	mess, err := handler.grade_service.CreateAccommodationGrade(accommodationGrade)
	if err != nil {
		err := status.Errorf(codes.Internal, mess)
		return nil, err
	}

	response := &pb.CreateAccommodationGradeResponse{
		Message: mess,
	}
	return response, nil
}

func (handler *AccommodationHandler) UpdateAccommodationGrade(ctx context.Context, request *pb.UpdateAccommodationGradeRequest) (*pb.UpdateAccommodationGradeResponse, error) {
	//TODO zahtjev 1.12 azuriranje ocijene, provjeriti da li je njegova ocijena da li smije da je promijeni
	ClientToken, _ := grpc_auth.AuthFromMD(ctx, "Bearer")
	claims, _ := token.ValidateToken(ClientToken)
	_, err := primitive.ObjectIDFromHex(claims.Uid)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided logged user id is not a valid ObjectID")
		return nil, err
	}
	mess, err := handler.grade_service.UpdateAccommodationGrade(int(request.Grade), request.Id, claims.Uid)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, mess)
		return nil, err
	}
	response := &pb.UpdateAccommodationGradeResponse{
		Message: mess,
	}
	return response, nil
}

func (handler *AccommodationHandler) DeleteAccommodationGrade(ctx context.Context, request *pb.DeleteAccommodationGradeRequest) (*pb.DeleteAccommodationGradeResponse, error) {
	//TODO zahtjev 1.12 brisanje ocijene, provjeriti da li je njegova ocijena da li smije da je obrise
	ClientToken, _ := grpc_auth.AuthFromMD(ctx, "Bearer")
	claims, _ := token.ValidateToken(ClientToken)
	_, err := primitive.ObjectIDFromHex(claims.Uid)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided logged user id is not a valid ObjectID")
		return nil, err
	}
	mess, err := handler.grade_service.DeleteAccommodationGrade(claims.Uid, request.Id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, mess)
		return nil, err
	}
	response := &pb.DeleteAccommodationGradeResponse{
		Message: mess,
	}
	return response, nil
}

func (handler *AccommodationHandler) GetAllAccommodationGrade(ctx context.Context, request *pb.GetAllAccommodationGradeRequest) (*pb.GetAllAccommodationGradeResponse, error) {
	//TODO zahtjev 1.12 dobavljanje svih ocijena koje je smjestaj dobio salje se id smjestaja
	//treba da se vrati lista svih ocijena tog smjestaja, napravio sam dto kako treba da izgleda
	// i treba da se izracuna prosijecna ocijena, vjerovatno cete morati mapper praviti neki da to vratite
	id := request.Id
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided accommodationId is not a valid ObjectID")
		return nil, err
	}
	accommodationGrades, err := handler.grade_service.GetAllAccommodationGrade(request.Id)
	if err != nil {
		return nil, err
	}
	user, err := handler.user_client.GetUser(createContextForAuthorization(ctx), &user_service.GetUserRequest{Id: "UserId"})
	if err != nil {
		return nil, err
	}
	var sum int
	var gradeDTOs []models.AccommodationGradeDetails
	for _, grade := range accommodationGrades {
		accomodation, err := handler.accommodation_service.GetAccommodationById(grade.AccommodationID)
		if err != nil {
			return nil, err
		}
		gradeDTO := models.AccommodationGradeDetails{GuestFirstName: *&user.User.FirstName, GuestLastName: *&user.User.LastName, AccommodationName: accomodation.Name, Grade: grade.Grade, DateOfGrade: grade.DateOfGrade}
		gradeDTOs = append(gradeDTOs, gradeDTO)
		sum = sum + grade.Grade
	}
	avergeGrade := float64(sum / len(accommodationGrades))
	gradesDetails := []*pb.AccommodationGradeDetails{}
	for _, r := range gradeDTOs {
		gradesPb := mapAccommodationGradeDetails(&r)
		gradesDetails = append(gradesDetails, gradesPb)
	}
	response := &pb.GetAllAccommodationGradeResponse{
		AccommodationGradeDetails: gradesDetails,
		AverageGrade:              avergeGrade,
	}
	return response, nil
}
