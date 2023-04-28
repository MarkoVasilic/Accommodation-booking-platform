package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

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

	err = mux.HandlePath("POST", "/availability", handler.CreateAvailability)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("PUT", "/availability/{availabilityId}", handler.UpdateAvailability)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("DELETE", "/availability/{availabilityId}", handler.DeleteAvailability)
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
		fmt.Println(err)
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
		fmt.Println(err)
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
		fmt.Println(err)
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
	//TODO mihaela
}

func (handler *GlobalHandler) CreateAvailability(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO nadja i aleksandra
}

func (handler *GlobalHandler) UpdateAvailability(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO nadja i aleksandra
}

func (handler *GlobalHandler) DeleteAvailability(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO marko
}

func (handler *GlobalHandler) SearchAvailability(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO nadja i aleksandra
}

func (handler *GlobalHandler) CreateReservation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO nadja i aleksandra
}

func (handler *GlobalHandler) GetFindReservationPendingGuest(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO nadja i aleksandra
}

func (handler *GlobalHandler) GetFindReservationAcceptedGuest(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO nadja i aleksandra
}

func (handler *GlobalHandler) GetFindReservationHost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO mihaela
}

func (handler *GlobalHandler) CancelReservation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO nadja i aleksandra
}

func (handler *GlobalHandler) DeleteLogicallyReservation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO nadja i aleksandra
}

func (handler *GlobalHandler) AcceptReservation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO mihaela
}
