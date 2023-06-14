package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/MarkoVasilic/Accommodation-booking-platform/api_gateway/domain"
	accommodation_service "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	reservation_service "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/reservation_service"
	user_service "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
)

type GlobalHandler struct {
	accommodationService accommodation_service.AccommodationServiceClient
	userService          user_service.UserServiceClient
	reservationService   reservation_service.ReservationServiceClient
}

func NewGlobalHandler(accommodationService accommodation_service.AccommodationServiceClient, userService user_service.UserServiceClient, reservationService reservation_service.ReservationServiceClient) Handler {
	return &GlobalHandler{
		accommodationService: accommodationService,
		userService:          userService,
		reservationService:   reservationService,
	}
}

func (handler *GlobalHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/user/{userId}", handler.GetUser)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("GET", "/user/logged", handler.GetLoggedUser)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("POST", "/user", handler.CreateUser)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("POST", "/login", handler.Login)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("PUT", "/user/{userId}", handler.UpdateUser)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("DELETE", "/user/{userId}", handler.DeleteUser)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("POST", "/accommodation", handler.CreateAccommodation)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("GET", "/accommodation/all/{hostId}", handler.GetAllAccommodations)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("POST", "/availability", handler.CreateAvailability)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("GET", "/availability/all/{accommodationId}", handler.GetAllAvailabilities)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("PUT", "/availability/{availabilityId}", handler.UpdateAvailability)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("POST", "/availability/search", handler.SearchAvailability)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("POST", "/reservation", handler.CreateReservation)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("GET", "/reservation/guest/pending/{id}", handler.GetFindReservationPendingGuest)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("GET", "/reservation/guest/accepted/{id}", handler.GetFindReservationAcceptedGuest)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("GET", "/reservation/host/{id}", handler.GetFindReservationHost)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("PUT", "/accommodation/reservation/cancel/{id}", handler.CancelReservation)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("PUT", "/accommodation/reservation/ldelete/{id}", handler.DeleteLogicallyReservation)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("PUT", "/accommodation/reservation/accept/{id}", handler.AcceptReservation)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("GET", "/user/guest/grades", handler.GetAllGuestGrades)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("GET", "/user/host/all", handler.GetAllHosts)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("POST", "/user/grade", handler.CreateUserGrade)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("PUT", "/user/grade/{id}", handler.UpdateUserGrade)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("DELETE", "/user/grade/{id}", handler.DeleteUserGrade)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("GET", "/user/grade/{id}", handler.GetAllUserGrade)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("GET", "/user/prominent/{id}", handler.HostProminent)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("GET", "/accommodations", handler.GetEveryAccommodation)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("POST", "/accommodation/grade", handler.CreateAccommodationGrade)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("PUT", "/accommodation/guest/grades/{id}", handler.UpdateAccommodationGrade)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("DELETE", "/accommodation/grade/{id}", handler.DeleteAccommodationGrade)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("GET", "/accommodation/grade/{id}", handler.GetAllAccommodationGrade)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("POST", "/availability/filter", handler.FilterAvailability)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("PUT", "/user/notificationon/{id}", handler.UpdateNotificationOn)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("POST", "/user/notification", handler.CreateNotification)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/user/notification/{id}", handler.GetAllNotifications)
	if err != nil {
		panic(err)
	}

}

func createContextForAuthorization(header []string) context.Context {
	if len(header) > 0 {
		return metadata.NewOutgoingContext(context.Background(), metadata.Pairs("Authorization", "Bearer "+header[0]))
	}
	return context.TODO()
}

func (handler *GlobalHandler) GetUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["userId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := handler.userService.GetUser(createContextForAuthorization(r.Header["Authorization"]), &user_service.GetUserRequest{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *GlobalHandler) GetLoggedUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	user, err := handler.userService.GetLoggedUser(createContextForAuthorization(r.Header["Authorization"]), &user_service.GetLoggedUserRequest{})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *GlobalHandler) CreateUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse request body: %v", err)
		return
	}
	resp, err := handler.userService.CreateUser(createContextForAuthorization(r.Header["Authorization"]), &user_service.CreateUserRequest{Username: user.Username, FirstName: user.FirstName, LastName: user.LastName, Password: user.Password, Email: user.Email, Address: user.Address, Role: user.Role})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call CreateUser method: %v", err)
		return
	}

	fmt.Fprintf(w, "%s", resp)
}

func (handler *GlobalHandler) Login(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse request body: %v", err)
		return
	}
	founduser, err := handler.userService.Login(createContextForAuthorization(r.Header["Authorization"]), &user_service.LoginRequest{Password: user.Password, Username: user.Username})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Failed to call Login method: %v", err)
		return
	}
	response, err := json.Marshal(founduser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *GlobalHandler) UpdateUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["userId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse request body: %v", err)
		return
	}
	resp, err := handler.userService.UpdateUser(createContextForAuthorization(r.Header["Authorization"]), &user_service.UpdateUserRequest{Id: id, Username: user.Username, FirstName: user.FirstName, LastName: user.LastName, Password: user.Password, Email: user.Email, Address: user.Address})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call UpdateUser method: %v", err)
		return
	}
	fmt.Fprintf(w, "%s", resp)
}

func (handler *GlobalHandler) DeleteUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["userId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := handler.userService.DeleteUser(createContextForAuthorization(r.Header["Authorization"]), &user_service.DeleteUserRequest{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call DeleteUser method: %v", err)
		return
	}
	fmt.Fprintf(w, "%s", resp)
}

func (handler *GlobalHandler) CreateAccommodation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var accommodation domain.Accommodation
	err := json.NewDecoder(r.Body).Decode(&accommodation)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse request body: %v", err)
		return
	}

	resp, err := handler.accommodationService.CreateAccommodation(createContextForAuthorization(r.Header["Authorization"]),
		&accommodation_service.CreateAccommodationRequest{
			Id:         "ttt",
			HostId:     accommodation.HostId,
			Name:       accommodation.Name,
			Location:   accommodation.Location,
			Wifi:       accommodation.Wifi,
			Kitchen:    accommodation.Kitchen,
			AC:         accommodation.AC,
			ParkingLot: accommodation.ParkingLot,
			MinGuests:  accommodation.MinGuests,
			MaxGuests:  accommodation.MaxGuests,
			Images:     accommodation.Images,
			AutoAccept: accommodation.AutoAccept})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call CreateAccommodation method: %v", err)
		return
	}

	fmt.Fprintf(w, "%s", resp)
}

func (handler *GlobalHandler) GetAllAccommodations(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	id := pathParams["hostId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := handler.accommodationService.GetAllAccommodations(createContextForAuthorization(r.Header["Authorization"]), &accommodation_service.GetAllAccommodationsRequest{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call GetAllAccommodations method: %v", err)
		return
	}
	//fmt.Fprintf(w, "%s", resp)
	response, err := json.Marshal(resp.Accommodations)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *GlobalHandler) GetAllAvailabilities(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["accommodationId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := handler.accommodationService.GetAllAvailabilities(createContextForAuthorization(r.Header["Authorization"]), &accommodation_service.GetAllAvailabilitiesRequest{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call GetAllAccommodations method: %v", err)
		return
	}
	//fmt.Fprintf(w, "%s", resp)
	response, err := json.Marshal(resp.Availabilities)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *GlobalHandler) CreateAvailability(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var availability domain.Availability
	err := json.NewDecoder(r.Body).Decode(&availability)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse request body: %v", err)
		return
	}
	startDate := timestamppb.New(availability.StartDate)
	endDate := timestamppb.New(availability.EndDate)
	resp, err := handler.accommodationService.CreateAvailability(createContextForAuthorization(r.Header["Authorization"]), &accommodation_service.CreateAvailabilityRequest{Id: availability.ID, AccommodationID: availability.AccommodationID, StartDate: startDate, EndDate: endDate, Price: availability.Price, IsPricePerGuest: availability.IsPricePerGuest})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call createAvailability method: %v", err)
		return
	}

	fmt.Fprintf(w, "%s", resp)
}

func (handler *GlobalHandler) UpdateAvailability(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["availabilityId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var availability domain.Availability
	err := json.NewDecoder(r.Body).Decode(&availability)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse request body: %v", err)
		return
	}

	startDate := timestamppb.New(availability.StartDate)
	endDate := timestamppb.New(availability.EndDate)
	resp, err := handler.accommodationService.UpdateAvailability(createContextForAuthorization(r.Header["Authorization"]), &accommodation_service.UpdateAvailabilityRequest{Id: id, StartDate: startDate, EndDate: endDate, Price: availability.Price, IsPricePerGuest: availability.IsPricePerGuest})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call UpdateAvailability method: %v", err)
		return
	}
	fmt.Fprintf(w, "%s", resp)
}

func (handler *GlobalHandler) SearchAvailability(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var findAvailability domain.FindAvailability
	err := json.NewDecoder(r.Body).Decode(&findAvailability)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse request body: %v", err)
		return
	}
	startDate := timestamppb.New(findAvailability.StartDate)
	endDate := timestamppb.New(findAvailability.EndDate)
	resp, err := handler.accommodationService.SearchAvailability(createContextForAuthorization(r.Header["Authorization"]), &accommodation_service.SearchAvailabilityRequest{Location: findAvailability.Location, GuestsNum: int32(findAvailability.GuestsNum), StartDate: startDate, EndDate: endDate})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call searchAvailability method: %v", err)
		return
	}
	response, err := json.Marshal(resp.FindAvailability)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (handler *GlobalHandler) CreateReservation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var reservation domain.Reservation
	err := json.NewDecoder(r.Body).Decode(&reservation)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse request body: %v", err)
		return
	}
	startDate := timestamppb.New(reservation.StartDate)
	endDate := timestamppb.New(reservation.EndDate)
	resp, err := handler.reservationService.CreateReservation(createContextForAuthorization(r.Header["Authorization"]), &reservation_service.CreateReservationRequest{Id: reservation.Id, AvailabilityID: reservation.AvailabilityID, GuestId: reservation.GuestId, StartDate: startDate, EndDate: endDate, NumGuests: int32(reservation.NumGuests)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call createReservation method: %v", err)
		return
	}

	fmt.Fprintf(w, "%s", resp)

}

func (handler *GlobalHandler) GetFindReservationPendingGuest(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := handler.reservationService.GetFindReservationPendingGuest(createContextForAuthorization(r.Header["Authorization"]), &reservation_service.GetFindReservationPendingGuestRequest{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call GetFindReservationPendingGuest method: %v", err)
		return
	}
	//fmt.Fprintf(w, "%s", resp)
	response, err := json.Marshal(resp.FindReservation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *GlobalHandler) GetFindReservationAcceptedGuest(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := handler.reservationService.GetFindReservationAcceptedGuest(createContextForAuthorization(r.Header["Authorization"]), &reservation_service.GetFindReservationAcceptedGuestRequest{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call GetFindReservationAcceptedGuest method: %v", err)
		return
	}

	response, err := json.Marshal(resp.FindReservation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *GlobalHandler) GetFindReservationHost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := handler.reservationService.GetFindReservationHost(createContextForAuthorization(r.Header["Authorization"]), &reservation_service.GetFindReservationHostRequest{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call GetFindReservationAcceptedGuest method: %v", err)
		return
	}

	response, err := json.Marshal(resp.FindReservation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *GlobalHandler) CancelReservation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := handler.reservationService.CancelReservation(createContextForAuthorization(r.Header["Authorization"]), &reservation_service.CancelReservationRequest{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call CancelReservation method: %v", err)
		return
	}
	fmt.Fprintf(w, "%s", resp)

}

func (handler *GlobalHandler) DeleteLogicallyReservation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := handler.reservationService.DeleteLogicallyReservation(createContextForAuthorization(r.Header["Authorization"]), &reservation_service.DeleteLogicallyReservationRequest{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call DeleteLogicallyReservation method: %v", err)
		return
	}
	fmt.Fprintf(w, "%s", resp)
}

func (handler *GlobalHandler) AcceptReservation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := handler.reservationService.AcceptReservation(createContextForAuthorization(r.Header["Authorization"]), &reservation_service.AcceptReservationRequest{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call AcceptReservation method: %v", err)
		return
	}
	fmt.Fprintf(w, "%s", resp)
}

func (handler *GlobalHandler) GetAllGuestGrades(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO
}

func (handler *GlobalHandler) GetAllHosts(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO
}

func (handler *GlobalHandler) CreateUserGrade(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO
}

func (handler *GlobalHandler) UpdateUserGrade(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO
}

func (handler *GlobalHandler) DeleteUserGrade(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO
}

func (handler *GlobalHandler) GetAllUserGrade(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO
}

func (handler *GlobalHandler) HostProminent(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO
}

func (handler *GlobalHandler) GetEveryAccommodation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO
}

func (handler *GlobalHandler) CreateAccommodationGrade(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO
}

func (handler *GlobalHandler) UpdateAccommodationGrade(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO
}

func (handler *GlobalHandler) DeleteAccommodationGrade(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO
}

func (handler *GlobalHandler) GetAllAccommodationGrade(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO
}

func (handler *GlobalHandler) FilterAvailability(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO
}

func (handler *GlobalHandler) UpdateNotificationOn(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	user_id := pathParams["id"]
	if user_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var notificationOn domain.NotificationOn
	err := json.NewDecoder(r.Body).Decode(&notificationOn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse request body: %v", err)
		return
	}
	resp, err := handler.userService.UpdateNotificationOn(createContextForAuthorization(r.Header["Authorization"]), &user_service.UpdateNotificationOnRequest{Id: user_id, Type: notificationOn.Type, On: notificationOn.On})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call UpdateNotificationOn method: %v", err)
		return
	}
	fmt.Fprintf(w, "%s", resp)

}

func (handler *GlobalHandler) CreateNotification(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var notification domain.Notification
	err := json.NewDecoder(r.Body).Decode(&notification)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse request body: %v", err)
		return
	}
	resp, err := handler.userService.CreateNotification(createContextForAuthorization(r.Header["Authorization"]), &user_service.CreateNotificationRequest{UserId: notification.UserID, Type: notification.Type, Message: *notification.Message, DateOfNotification: nil})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call CreateUser method: %v", err)
		return
	}

	fmt.Fprintf(w, "%s", resp)
}

func (handler *GlobalHandler) GetAllNotifications(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO
	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := handler.userService.GetAllNotifications(createContextForAuthorization(r.Header["Authorization"]), &user_service.GetAllNotificationsRequest{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call GetAllNotifications method: %v", err)
		return
	}
	//fmt.Fprintf(w, "%s", resp)
	response, err := json.Marshal(resp.Notifications)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
