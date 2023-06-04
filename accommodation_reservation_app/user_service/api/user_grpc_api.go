package api

import (
	"context"
	"fmt"
	"strconv"
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
	err = handler.notification_on_service.InitializeNotificationsOn(user_id)
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

/*
func (handler *UserHandler) GetAllGuestGrades(ctx context.Context, request *pb.GetAllGuestGradesRequest) (*pb.GetAllGuestGradesResponse, error) {
	//TODO pomocna metoda za dobavljanje svih ocijena guesta za poslani id guesta
	//a vraca se lista dtova koji sam napravio
}
*/

/*
func (handler *UserHandler) GetAllHosts(ctx context.Context, request *pb.GetAllHostsRequest) (*pb.GetAllHostsResponse, error) {
	//TODO pomocna metoda za dobavljanje svih hostova, ne salje se nista
	//a vraca se lista dtova koji sam napravio
}
*/

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
	grade, err := strconv.Atoi(request.Grade)
	userGrade := models.UserGrade{GuestID: guestId, HostID: hostId, Grade: grade, DateOfGrade: time.Now()}
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
	mess, err := handler.grade_service.UpdateUserGrade( /*request.Grade*/ 5, request.Id) //izmeniti proto pa otkomentarisati
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
	//ovako se izvlaci id osobe koja salje zahtjev, id se nalazi u claims.Uid
	//provjeriti uslov da li moze da ga ocijeni
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

/*
func (handler *UserHandler) GetAllUserGrade(ctx context.Context, request *pb.GetAllUserGradeRequest) (*pb.GetAllUserGradeResponse, error) {
	//TODO zahtjev 1.11 dobavljanje svih ocijena koje je host dobio salje se id hosta
	//treba da se vrati lista svih ocijena tog hosta, napravio sam dto kako treba da izgleda
	// i treba da se izracuna prosijecna ocijena, vjerovatno cete morati mapper praviti neki da to vratite
}*/

func (handler *UserHandler) HostProminent(ctx context.Context, request *pb.HostProminentRequest) (*pb.HostProminentResponse, error) {
	//TODO zahtjev 1.13, salje se id hosta ili ga izvucite kao gore sto sam naveo i vraca se bool koje je true ako je istaknut i false ako nije
	//na frontu mozete staviti na ono profile kad se otvori dodatno dugme koje ovo provjerava ili samo da napravite jos jedno polje koje kaze da li je istaknut ili ne
	//samo pazite da je ovo samo za hosta, a na frontu stranica profila je ista za guesta i hosta pa morate to nekako da razdvojite
	response := &pb.HostProminentResponse{
		Prominent: true,
	}
	return response, nil
}

func (handler *UserHandler) UpdateNotificationOn(ctx context.Context, request *pb.UpdateNotificationOnRequest) (*pb.UpdateNotificationOnResponse, error) {
	//TODO zahtjev 1.15, azuriranje odgovarajuceg tipa notifikacije da li je ukljuceno ili ne
	response := &pb.UpdateNotificationOnResponse{
		Message: "Success",
	}
	return response, nil
}

func (handler *UserHandler) CreateNotification(ctx context.Context, request *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error) {
	//TODO zahtjev 1.15, kreiranje notifikacije
	response := &pb.CreateNotificationResponse{
		Message: "Success",
	}
	return response, nil
}

/*
func (handler *UserHandler) GetAllNotifications(ctx context.Context, request *pb.GetAllNotificationsRequest) (*pb.GetAllNotificationsResponse, error) {
	//TODO zahtjev 1.15, dobavljanje notifikacija, treba provjeriti koje su ukljucene i te dobaviti za korisnika za kojeg je id poslan
	//mozes dodatno sortirati po datumu pravljenja da se vide najnovije
	//postoji i polje seen kojeg mozes ukljuciti a i ne moras, neka stoji samo ako ne zelis

}*/
