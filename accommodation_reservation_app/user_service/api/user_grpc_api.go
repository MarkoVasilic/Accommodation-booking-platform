package api

import (
	"context"
	"fmt"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/token"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service                 *service.UserService
	grade_service           *service.GradeService
	notification_service    *service.NotificationService
	notification_on_service *service.NotificationOnService
	accommodation_client    accommodation_service.AccommodationServiceClient
	reservation_client      reservation_service.ReservationServiceClient
}

func NewUserHandler(service *service.UserService, grade_service *service.GradeService, notification_service *service.NotificationService, notification_on_service *service.NotificationOnService, accommodation_client accommodation_service.AccommodationServiceClient, reservation_client reservation_service.ReservationServiceClient) *UserHandler {
	return &UserHandler{
		service:                 service,
		grade_service:           grade_service,
		notification_service:    notification_service,
		notification_on_service: notification_on_service,
		accommodation_client:    accommodation_client,
		reservation_client:      reservation_client,
	}
}

func createContextForAuthorization(ctx context.Context) context.Context {
	token, _ := grpc_auth.AuthFromMD(ctx, "Bearer")
	if len(token) > 0 {
		return metadata.NewOutgoingContext(context.Background(), metadata.Pairs("Authorization", "Bearer "+token))
	}
	return context.TODO()
}

func (handler *UserHandler) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	user, err := handler.service.GetUserById(objectId)
	if err != nil {
		return nil, err
	}
	userPb := mapUser(&user)
	response := &pb.GetUserResponse{
		User: userPb,
	}
	return response, nil
}

func (handler *UserHandler) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := models.User{Username: &request.Username, FirstName: &request.FirstName, LastName: &request.LastName, Password: &request.Password, Email: &request.Email, Address: &request.Address, Role: (*models.Role)(&request.Role)}
	mess, err, user_id := handler.service.CreateUser(user)
	if err != nil {
		err := status.Errorf(codes.Internal, mess)
		return nil, err
	}
	err = handler.notification_on_service.InitializeNotificationsOn(user)
	if err != nil {
		handler.service.DeleteUser(user_id, true, "")
		err := status.Errorf(codes.Internal, mess)
		return nil, err
	}
	response := &pb.CreateUserResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *UserHandler) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	id := request.Id
	user := models.User{Username: &request.Username, FirstName: &request.FirstName, LastName: &request.LastName, Password: &request.Password, Email: &request.Email, Address: &request.Address}
	mess, err := handler.service.UpdateUser(user, id)
	if err != nil {
		err := status.Errorf(codes.Internal, mess)
		return nil, err
	}
	response := &pb.UpdateUserResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *UserHandler) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	id := request.Id
	ClientToken, _ := grpc_auth.AuthFromMD(ctx, "Bearer")
	claims, _ := token.ValidateToken(ClientToken)
	if claims.Uid != id {
		err := status.Errorf(codes.PermissionDenied, "you are only allowed to delete yourself")
		response := &pb.DeleteUserResponse{
			Message: "you are only allowed to delete yourself",
		}
		return response, err
	}
	delete_now := false
	if claims.Role == "GUEST" {
		reservations, err := handler.reservation_client.GetFindReservationAcceptedGuest(createContextForAuthorization(ctx), &reservation_service.GetFindReservationAcceptedGuestRequest{Id: id})
		if err != nil {
			if err.Error() != "rpc error: code = InvalidArgument desc = There is no accepted reservations!" {
				err := status.Errorf(codes.Internal, "something went wrong")
				response := &pb.DeleteUserResponse{
					Message: "something went wrong",
				}
				return response, err
			}
		}
		if reservations != nil && len(reservations.FindReservation) > 0 {
			err := status.Errorf(codes.PermissionDenied, "There are existing reservations, please cancel them before you proceede")
			response := &pb.DeleteUserResponse{
				Message: "There are existing reservations, please cancel them before you proceede",
			}
			return response, err
		}
		delete_now = true
	}
	if claims.Role == "HOST" {
		reservations, err := handler.reservation_client.GetAllReservationsHost(createContextForAuthorization(ctx), &reservation_service.GetAllReservationsHostRequest{Id: id})
		if err != nil {
			if err.Error() != "rpc error: code = InvalidArgument desc = There is no accommodations!" {
				err := status.Errorf(codes.Internal, "something went wrong")
				response := &pb.DeleteUserResponse{
					Message: "something went wrong",
				}
				return response, err
			}
		}
		if reservations != nil && len(reservations.Reservation) > 0 {
			err := status.Errorf(codes.PermissionDenied, "There are existing reservations, please cancel them before you proceede")
			response := &pb.DeleteUserResponse{
				Message: "There are existing reservations, please cancel them before you proceede",
			}
			return response, err
		}
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	mess, err := handler.service.DeleteUser(objectId, delete_now, ClientToken)
	if err != nil {
		response := &pb.DeleteUserResponse{
			Message: mess,
		}
		return response, err
	}
	response := &pb.DeleteUserResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *UserHandler) GetLoggedUser(ctx context.Context, request *pb.GetLoggedUserRequest) (*pb.GetLoggedUserResponse, error) {
	ClientToken, _ := grpc_auth.AuthFromMD(ctx, "Bearer")
	if len(ClientToken) < 1 {
		return nil, fmt.Errorf("No token provided")
	}
	claims, _ := token.ValidateToken(ClientToken)
	if claims == nil {
		return nil, nil
	}
	objectId, err := primitive.ObjectIDFromHex(claims.Uid)
	user, err := handler.service.GetUserById(objectId)

	if err != nil {
		return nil, err
	}
	userPb := mapUser(&user)
	response := &pb.GetLoggedUserResponse{
		User: userPb,
	}
	return response, nil
}

func (handler *UserHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	user := models.User{Password: &request.Password, Username: &request.Username}
	founduser, err := handler.service.Login(user)
	if err != "" {
		err := status.Errorf(codes.Internal, err)
		return nil, err
	}

	userPb := mapUser(&founduser)
	response := &pb.LoginResponse{
		User: userPb,
	}
	return response, nil
}

func (handler *UserHandler) GetAllGuestGrades(ctx context.Context, request *pb.GetAllGuestGradesRequest) (*pb.GetAllGuestGradesResponse, error) {
	//TODO pomocna metoda za dobavljanje svih ocijena guesta za poslani id guesta
	//a vraca se lista dtova koji sam napravio
	id := request.Id
	guestId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided guestId is not a valid ObjectID")
		return nil, err
	}
	res, err := handler.grade_service.GetAllGuestGrades(request.Id)
	if err != nil {
		return nil, err
	}
	user, err := handler.service.GetUserById(guestId)
	if err != nil {
		return nil, err
	}
	var gradeDTOs []models.UserGradeDetails
	for _, grade := range res {
		host, err := handler.service.GetUserById(grade.HostID)
		if err != nil {
			return nil, err
		}
		gradeDTO := models.UserGradeDetails{GuestFirstName: *user.FirstName, GuestLastName: *user.LastName, HostFirstName: *host.FirstName, HostLastName: *host.LastName, Grade: grade.Grade, DateOfGrade: grade.DateOfGrade}
		gradeDTOs = append(gradeDTOs, gradeDTO)
	}
	gradesDetails := []*pb.UserGradeDetails{}
	for _, r := range gradeDTOs {
		gradesPb := mapUserGradeDetails(&r)
		gradesDetails = append(gradesDetails, gradesPb)
	}
	response := &pb.GetAllGuestGradesResponse{
		UserGradeDetails: gradesDetails,
	}
	return response, nil
}

func (handler *UserHandler) GetAllHosts(ctx context.Context, request *pb.GetAllHostsRequest) (*pb.GetAllHostsResponse, error) {
	//TODO pomocna metoda za dobavljanje svih hostova koje ulogovani user moze da oceni, ne salje se nista
	//a vraca se lista dtova koji sam napravio
	ClientToken, _ := grpc_auth.AuthFromMD(ctx, "Bearer")
	claims, _ := token.ValidateToken(ClientToken)
	_, err := primitive.ObjectIDFromHex(claims.Uid)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	reservations, err := handler.reservation_client.GetAllReservations(createContextForAuthorization(ctx), &reservation_service.GetAllReservationsRequest{Id: claims.Id})
	if reservations == nil {
		err := status.Errorf(codes.InvalidArgument, "There is no reservations!")
		return nil, err
	} else if err != nil {
		return nil, err
	}
	hostsDetails := []models.HostDetails{}
	for _, res := range reservations.Reservations {
		if res.IsCanceled || res.IsDeleted {
			continue
		}
		availabilitiy, err := handler.accommodation_client.GetAvailabilityById(createContextForAuthorization(ctx), &accommodation_service.GetAvailabilityByIdRequest{Id: res.AvailabilityID})
		if err != nil {
			err := status.Errorf(codes.InvalidArgument, "the no availability")
			return nil, err
		}
		accommodation, err := handler.accommodation_client.GetAccommodationById(createContextForAuthorization(ctx), &accommodation_service.GetAccommodationByIdRequest{Id: availabilitiy.Availability.AccommodationID})
		if err != nil {
			err := status.Errorf(codes.InvalidArgument, "the no accommodation")
			return nil, err
		}
		hostId, err := primitive.ObjectIDFromHex(accommodation.Accommodation.HostId)
		if err != nil {
			err := status.Errorf(codes.InvalidArgument, "the provided hostId is not a valid ObjectID")
			return nil, err
		}
		host, err := handler.service.GetUserById(hostId)
		hostDetails := models.HostDetails{Id: hostId, FirstName: *host.FirstName, LastName: *host.LastName}
		hostsDetails = append(hostsDetails, hostDetails)
	}
	if hostsDetails == nil {
		err := status.Errorf(codes.InvalidArgument, "There is no hosts to be graded!")
		return nil, err
	}

	hosts := []*pb.HostDetails{}
	for _, h := range hostsDetails {
		hostDetailsPb := mapHost(&h)
		hosts = append(hosts, hostDetailsPb)
	}
	response := &pb.GetAllHostsResponse{
		Hosts: hosts,
	}
	return response, nil
}

func (handler *UserHandler) CreateUserGrade(ctx context.Context, request *pb.CreateUserGradeRequest) (*pb.CreateUserGradeResponse, error) {
	//TODO zahtjev 1.11 kreiranje ocijene
	ClientToken, _ := grpc_auth.AuthFromMD(ctx, "Bearer")
	claims, _ := token.ValidateToken(ClientToken)
	//ovako se izvlaci id osobe koja salje zahtjev, id se nalazi u claims.Uid
	//provjeriti uslov da li moze da ga ocijeni
	guestId, err := primitive.ObjectIDFromHex(claims.Uid)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	hostId, err := primitive.ObjectIDFromHex(request.HostID)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided id is not a valid ObjectID")
		return nil, err
	}
	//provera da li sme da oceni
	reservations, err := handler.reservation_client.GetAllReservations(createContextForAuthorization(ctx), &reservation_service.GetAllReservationsRequest{Id: claims.Id})
	if err != nil {
		return nil, err
	}
	if len(reservations.Reservations) == 0 {
		err := status.Errorf(codes.InvalidArgument, "You cannot grade host if you have no reservations!")
		return nil, err
	}
	numOfGuestReservations := 0
	for _, res := range reservations.Reservations {
		if res.IsCanceled || res.IsDeleted {
			continue
		}
		_, err := primitive.ObjectIDFromHex(res.AvailabilityID)
		if err != nil {
			err := status.Errorf(codes.InvalidArgument, "the provided availabilityID is not a valid ObjectID")
			return nil, err
		}
		availabilitiy, err := handler.accommodation_client.GetAvailabilityById(createContextForAuthorization(ctx), &accommodation_service.GetAvailabilityByIdRequest{Id: res.AvailabilityID})
		accommodation, err := handler.accommodation_client.GetAccommodationById(createContextForAuthorization(ctx), &accommodation_service.GetAccommodationByIdRequest{Id: availabilitiy.Availability.AccommodationID})
		if accommodation.Accommodation.HostId == request.HostID {
			numOfGuestReservations = numOfGuestReservations + 1
		}
		if numOfGuestReservations >= 1 {
			break
		}
	}
	if numOfGuestReservations < 1 {
		err := status.Errorf(codes.InvalidArgument, "You cannot grade host if you never stayed in theirs accommodation!")
		return nil, err
	}

	userGrade := models.UserGrade{GuestID: guestId, HostID: hostId, Grade: int(request.Grade), DateOfGrade: time.Now()}
	mess, err := handler.grade_service.CreateUserGrade(userGrade)
	if err != nil {
		err := status.Errorf(codes.Internal, mess)
		return nil, err
	}
	response := &pb.CreateUserGradeResponse{
		Message: mess,
	}
	return response, nil
}

func (handler *UserHandler) UpdateUserGrade(ctx context.Context, request *pb.UpdateUserGradeRequest) (*pb.UpdateUserGradeResponse, error) {
	//TODO zahtjev 1.11 azuriranje ocijene, provjeriti da li je njegova ocijena da li smije da je promijeni
	ClientToken, _ := grpc_auth.AuthFromMD(ctx, "Bearer")
	claims, _ := token.ValidateToken(ClientToken)
	_, err := primitive.ObjectIDFromHex(claims.Uid)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided logged user id is not a valid ObjectID")
		return nil, err
	}
	mess, err := handler.grade_service.UpdateUserGrade(int(request.Grade), request.Id, claims.Id) //izmeniti proto pa otkomentarisati
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, mess)
		return nil, err
	}
	response := &pb.UpdateUserGradeResponse{
		Message: mess,
	}
	return response, nil
}

func (handler *UserHandler) DeleteUserGrade(ctx context.Context, request *pb.DeleteUserGradeRequest) (*pb.DeleteUserGradeResponse, error) {
	//TODO zahtjev 1.11 brisanje ocijene, provjeriti da li je njegova ocijena da li smije da je obrise
	ClientToken, _ := grpc_auth.AuthFromMD(ctx, "Bearer")
	claims, _ := token.ValidateToken(ClientToken)
	_, err := primitive.ObjectIDFromHex(claims.Uid)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided logged user id is not a valid ObjectID")
		return nil, err
	}
	mess, err := handler.grade_service.DeleteUserGrade(request.Id, claims.Uid)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, mess)
		return nil, err
	}
	response := &pb.DeleteUserGradeResponse{
		Message: mess,
	}
	return response, nil
}

func (handler *UserHandler) GetAllUserGrade(ctx context.Context, request *pb.GetAllUserGradeRequest) (*pb.GetAllUserGradeResponse, error) {
	//TODO zahtjev 1.11 dobavljanje svih ocijena koje je host dobio salje se id hosta
	//treba da se vrati lista svih ocijena tog hosta, napravio sam dto kako treba da izgleda
	// i treba da se izracuna prosijecna ocijena, vjerovatno cete morati mapper praviti neki da to vratite
	id := request.Id
	guestId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided guestId is not a valid ObjectID")
		return nil, err
	}
	hostGrades, err := handler.grade_service.GetAllUserGrade(request.Id)
	if err != nil {
		return nil, err
	}
	user, err := handler.service.GetUserById(guestId)
	if err != nil {
		return nil, err
	}
	var sum int
	var gradeDTOs []models.UserGradeDetails
	for _, grade := range hostGrades {
		host, err := handler.service.GetUserById(grade.HostID)
		if err != nil {
			return nil, err
		}
		gradeDTO := models.UserGradeDetails{GuestFirstName: *user.FirstName, GuestLastName: *user.LastName, HostFirstName: *host.FirstName, HostLastName: *host.LastName, Grade: grade.Grade, DateOfGrade: grade.DateOfGrade}
		gradeDTOs = append(gradeDTOs, gradeDTO)
		sum = sum + grade.Grade
	}
	avergeGrade := float64(sum / len(hostGrades))
	gradesDetails := []*pb.UserGradeDetails{}
	for _, r := range gradeDTOs {
		gradesPb := mapUserGradeDetails(&r)
		gradesDetails = append(gradesDetails, gradesPb)
	}
	finalResp := mapUserGradeDetailsDTO(gradesDetails, avergeGrade)
	response := &pb.GetAllUserGradeResponse{
		UserGradeDetailsDTO: finalResp,
	}
	return response, nil
}

func (handler *UserHandler) HostProminent(ctx context.Context, request *pb.HostProminentRequest) (*pb.HostProminentResponse, error) {
	//TODO zahtjev 1.13, salje se id hosta ili ga izvucite kao gore sto sam naveo i vraca se bool koje je true ako je istaknut i false ako nije
	//na frontu mozete staviti na ono profile kad se otvori dodatno dugme koje ovo provjerava ili samo da napravite jos jedno polje koje kaze da li je istaknut ili ne
	//samo pazite da je ovo samo za hosta, a na frontu stranica profila je ista za guesta i hosta pa morate to nekako da razdvojite
	id := request.Id
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "the provided hostId is not a valid ObjectID")
		return nil, err
	}

	hostGrades, err := handler.grade_service.GetAllUserGrade(request.Id)
	if err != nil {
		return nil, err
	}
	sum := 0
	for _, grade := range hostGrades {
		sum = sum + grade.Grade
	}
	numberOfCancelation := 0
	var sumReservationDurations float64
	averageGrade := float64(sum / len(hostGrades))
	reservations, err := handler.reservation_client.GetAllReservationsHost(createContextForAuthorization(ctx), &reservation_service.GetAllReservationsHostRequest{Id: request.Id})
	for _, res := range reservations.Reservation {
		if res.IsCanceled == true {
			numberOfCancelation = numberOfCancelation + 1
		}
		year, month, day := res.StartDate.AsTime().Date()
		yearE, monthE, dayE := res.EndDate.AsTime().Date()
		startDate := time.Date(year, month, day, int(0), int(0), int(0), int(0), time.UTC)
		endDate := time.Date(yearE, monthE, dayE, int(0), int(0), int(0), int(0), time.UTC)
		duration := (endDate.Sub(startDate)).Hours()
		sumReservationDurations = sumReservationDurations + duration
	}
	sumReservationDurations = sumReservationDurations / 24
	cancelationPercent := (numberOfCancelation / len(reservations.Reservation)) * 100

	var prominent bool = false
	if averageGrade > 4.7 && cancelationPercent < 5 && len(reservations.Reservation) >= 5 && sumReservationDurations > 50 {
		prominent = true
	}

	response := &pb.HostProminentResponse{
		Prominent: prominent,
	}
	return response, nil
}

func (handler *UserHandler) UpdateNotificationOn(ctx context.Context, request *pb.UpdateNotificationOnRequest) (*pb.UpdateNotificationOnResponse, error) {
	//TODO zahtjev 1.15, azuriranje odgovarajuceg tipa notifikacije da li je ukljuceno ili ne
	user_id := request.Id
	notificationOn := models.NotificationOn{Type: (*models.NotificationType)(&request.Type), On: request.On}

	mess, err := handler.notification_on_service.UpdateNotificationOn(notificationOn, user_id)
	if err != nil {
		err := status.Errorf(codes.Internal, mess)
		return nil, err
	}

	response := &pb.UpdateNotificationOnResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *UserHandler) CreateNotification(ctx context.Context, request *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error) {
	//TODO zahtjev 1.15, kreiranje notifikacije
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		err := status.Errorf(codes.Internal, "Id conversion failed")
		response1 := &pb.CreateNotificationResponse{
			Message: "Id conversion failed",
		}
		return response1, err
	}

	notification := models.Notification{UserID: userId, Type: (*models.NotificationType)(&request.Type), Message: &request.Message, DateOfNotification: time.Now(), Seen: false}
	mess, err := handler.notification_service.CreateNotification(notification)
	if err != nil {
		err := status.Errorf(codes.Internal, mess)
		response2 := &pb.CreateNotificationResponse{
			Message: "Error with creating notification",
		}
		return response2, err
	}

	response := &pb.CreateNotificationResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *UserHandler) GetAllNotifications(ctx context.Context, request *pb.GetAllNotificationsRequest) (*pb.GetAllNotificationsResponse, error) {
	//TODO zahtjev 1.15, dobavljanje notifikacija, treba provjeriti koje su ukljucene i te dobaviti za korisnika za kojeg je id poslan
	//mozes dodatno sortirati po datumu pravljenja da se vide najnovije
	//postoji i polje seen kojeg mozes ukljuciti a i ne moras, neka stoji samo ako ne zelis
	userId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		err := status.Errorf(codes.Internal, "Id conversion failed")
		return nil, err
	}

	notifications, err := handler.notification_service.GetNotificationByUser(userId)
	if err != nil {
		return nil, err
	} else if notifications == nil {
		err := status.Errorf(codes.InvalidArgument, "There is no notifications!")
		return nil, err
	}

	notificationsOn, err := handler.notification_on_service.GetNotificationOnByUser(userId)
	if err != nil {
		return nil, err
	} else if notifications == nil {
		err := status.Errorf(codes.InvalidArgument, "There is no notifications!")
		return nil, err
	}

	user_notifications := []*pb.Notification{}
	for _, n := range notifications {
		for _, notificationOn := range notificationsOn {
			if *notificationOn.Type == *n.Type && notificationOn.On {
				notificationPb := mapNotification(&n)
				user_notifications = append(user_notifications, notificationPb)
			}
		}

	}

	//sortiraj

	response := &pb.GetAllNotificationsResponse{
		Notifications: user_notifications,
	}

	return response, nil
}

func (handler *UserHandler) GetUserNotificationsOn(ctx context.Context, request *pb.GetUserNotificationsOnRequest) (*pb.GetUserNotificationsOnResponse, error) {
	userId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		err := status.Errorf(codes.Internal, "Id conversion failed")
		return nil, err
	}

	notifications, err := handler.notification_on_service.GetNotificationOnByUser(userId)
	if err != nil {
		return nil, err
	} else if notifications == nil {
		err := status.Errorf(codes.InvalidArgument, "There is no notifications!")
		return nil, err
	}

	user_notifications := []*pb.NotificationOn{}
	for _, n := range notifications {
		notificationPb := mapNotificationOn(&n)
		user_notifications = append(user_notifications, notificationPb)
	}

	response := &pb.GetUserNotificationsOnResponse{
		Notifications: user_notifications,
	}

	return response, nil
}
