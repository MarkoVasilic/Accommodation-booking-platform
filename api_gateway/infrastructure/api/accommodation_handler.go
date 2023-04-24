package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MarkoVasilic/Accommodation-booking-platform/api_gateway/domain"
	accommodation "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
)

type AccommodationHandler struct {
	accommodationService accommodation.AccommodationServiceClient
}

func NewAccommodationHandler(accommodationService accommodation.AccommodationServiceClient) Handler {
	return &AccommodationHandler{
		accommodationService: accommodationService,
	}
}

func (handler *AccommodationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/user/{userId}", handler.GetUser)
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
}

func createContextForAuthorization(header []string) context.Context {
	if len(header) > 0 {
		return metadata.NewOutgoingContext(context.Background(), metadata.Pairs("Authorization", "Bearer "+header[0]))
	}
	return context.TODO()
}

func (handler *AccommodationHandler) GetUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["userId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := handler.accommodationService.GetUser(createContextForAuthorization(r.Header["Authorization"]), &accommodation.GetUserRequest{Id: id})
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

func (handler *AccommodationHandler) CreateUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse request body: %v", err)
		return
	}
	resp, err := handler.accommodationService.CreateUser(createContextForAuthorization(r.Header["Authorization"]), &accommodation.CreateUserRequest{Username: user.Username, FirstName: user.FirstName, LastName: user.LastName, Password: user.Password, Email: user.Email, Address: user.Address, Role: user.Role})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call CreateUser method: %v", err)
		return
	}

	fmt.Fprintf(w, "%s", resp)
}

func (handler *AccommodationHandler) Login(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse request body: %v", err)
		return
	}
	founduser, err := handler.accommodationService.Login(createContextForAuthorization(r.Header["Authorization"]), &accommodation.LoginRequest{Password: user.Password, Username: user.Username})
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

func (handler *AccommodationHandler) UpdateUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
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
	resp, err := handler.accommodationService.UpdateUser(createContextForAuthorization(r.Header["Authorization"]), &accommodation.UpdateUserRequest{Id: id, Username: user.Username, FirstName: user.FirstName, LastName: user.LastName, Password: user.Password, Email: user.Email, Address: user.Address})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call UpdateUser method: %v", err)
		return
	}
	fmt.Fprintf(w, "%s", resp)
}

func (handler *AccommodationHandler) DeleteUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["userId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := handler.accommodationService.DeleteUser(createContextForAuthorization(r.Header["Authorization"]), &accommodation.DeleteUserRequest{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to call DeleteUser method: %v", err)
		return
	}
	fmt.Fprintf(w, "%s", resp)
}
