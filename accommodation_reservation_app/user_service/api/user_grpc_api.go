package api

import (
	"context"
	"fmt"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/service"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/token"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
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
	mess, err := handler.service.CreateUser(user)
	if err != nil {
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
	objectId, err := primitive.ObjectIDFromHex(id)
	mess, err := handler.service.DeleteUser(objectId)
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