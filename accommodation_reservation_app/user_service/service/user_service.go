package service

import (
	"context"
	"log"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/repository"
	generate "github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/token"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	UserRepository *repository.UserRepository
	Orchestrator   *DeleteUserOrchestrator
}

var Validate = validator.New()

func (service *UserService) GetUserById(id primitive.ObjectID) (models.User, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	founduser, err := service.UserRepository.GetUserById(id)
	return founduser, err
}

func (service *UserService) CreateUser(user models.User) (string, error, primitive.ObjectID) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	validationErr := Validate.Struct(user)
	if validationErr != nil {
		err := status.Errorf(codes.InvalidArgument, "user fields are not valid")
		return "user fields are not valid", err, primitive.NilObjectID
	}

	count_email, err := service.UserRepository.CountByEmail(*user.Email)
	if err != nil {
		log.Panic(err)
		err := status.Errorf(codes.Internal, "something went wrong")
		return "something went wrong", err, primitive.NilObjectID
	}
	if count_email > 0 {
		err := status.Errorf(codes.NotFound, "user with that email already exists")
		return "user with that email already exists", err, primitive.NilObjectID
	}
	count_username, err := service.UserRepository.CountByUsername(*user.Username)
	if err != nil {
		log.Panic(err)
		err := status.Errorf(codes.Internal, "something went wrong")
		return "something went wrong", err, primitive.NilObjectID
	}
	if count_username > 0 {
		err := status.Errorf(codes.NotFound, "user with that username already exists")
		return "user with that username already exists", err, primitive.NilObjectID
	} else {
		inserterr, user_id := service.UserRepository.CreateUser(&user)

		if inserterr != nil {
			log.Panic(err)
			err := status.Errorf(codes.Internal, "something went wrong")
			return "something went wrong", err, user_id
		}
		return "Succesffully added new user", nil, user_id
	}
}

func (service *UserService) UpdateUser(user models.User, id string) (string, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(id)
	founduser, founderr := service.UserRepository.GetUserById(objectId)
	if founderr != nil {
		log.Panic(err)
		err := status.Errorf(codes.Internal, "something went wrong")
		return "something went wrong", err
	}
	user.Role = founduser.Role
	validationErr := Validate.Struct(user)
	if validationErr != nil {
		err := status.Errorf(codes.InvalidArgument, "user fields are not valid")
		return "user fields are not valid", err
	}
	if *founduser.Email != *user.Email {
		count_email, err := service.UserRepository.CountByEmail(*user.Email)
		if err != nil {

			log.Panic(err)
			err := status.Errorf(codes.Internal, "something went wrong")
			return "something went wrong", err
		}
		if count_email > 0 {
			err := status.Errorf(codes.NotFound, "user with that email already exists")
			return "user with that email already exists", err
		}
	}
	if *founduser.Username != *user.Username {
		count_username, err := service.UserRepository.CountByUsername(*user.Username)
		if err != nil {
			log.Panic(err)
			err := status.Errorf(codes.Internal, "something went wrong")
			return "something went wrong", err
		}
		if count_username > 0 {
			err := status.Errorf(codes.NotFound, "user with that username already exists")
			return "user with that username already exists", err
		}

	}
	user.Id = objectId
	inserterr := service.UserRepository.UpdateUser(&user)
	if inserterr != nil {
		log.Panic(err)
		err := status.Errorf(codes.Internal, "something went wrong")
		return "something went wrong", err
	}
	return "Succesffully updated user", nil
}

func (s *UserService) DeleteUser(id primitive.ObjectID, delete_now bool, ClientToken string) (string, error) {
	if delete_now {
		result, err := s.UserRepository.DeleteUser(id)
		if err != nil {
			log.Panic(err)
			err := status.Errorf(codes.Internal, "something went wrong")
			return "something went wrong", err
		}

		if result.DeletedCount == 0 {
			err := status.Errorf(codes.NotFound, "no user found with the given ID")
			return "no user found with the given ID", err
		}
		return "Succesffully deleted user", nil
	}
	err := s.Orchestrator.Start(id, ClientToken)
	if err != nil {
		return "Failed to delete user", nil
	}
	return "Succesffully deleted user", nil
}

func (service *UserService) Login(user models.User) (models.User, string) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	founduser, err := service.UserRepository.GetUserByUsername(*user.Username)
	if err != nil {
		return user, "user with that username doesn't exist"
	}
	if *user.Password != *founduser.Password {
		return user, "wrong password"
	}
	token, refreshToken, _ := generate.TokenGenerator(*founduser.Email, *founduser.FirstName, *founduser.LastName, founduser.User_ID, models.Role(*founduser.Role))
	generate.UpdateAllTokens(service.UserRepository.UserCollection, token, refreshToken, founduser.User_ID)
	return founduser, ""
}

func (service *UserService) GetAllHosts() ([]models.User, error) {
	hosts, err := service.UserRepository.GetAllHosts()
	if err != nil {
		return nil, err
	}

	return hosts, nil
}
