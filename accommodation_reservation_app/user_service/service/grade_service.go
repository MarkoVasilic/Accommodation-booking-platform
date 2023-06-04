package service

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GradeService struct {
	GradeRepository *repository.GradeRepository
}

func (service *GradeService) CreateUserGrade(userGrade models.UserGrade) (string, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	validationErr := Validate.Struct(userGrade)
	if validationErr != nil {
		err := status.Errorf(codes.InvalidArgument, "user grade fields are not valid")
		return "User grade fields are not valid", err
	}

	inserterr := service.GradeRepository.CreateUserGrade(&userGrade)
	if inserterr != nil {
		err := status.Errorf(codes.Internal, "failed to create user grade")
		return "Failed to create user grade", err
	}

	return "Successfully created user grade", nil
}

func (service *GradeService) UpdateUserGrade(id string, grade int) (string, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "Invalid ID format")
		return "Invalid ID format", err
	}

	//provera da li postoji ocena
	_, err = service.GradeRepository.GetGradeById(objectID)
	if err != nil {
		err := status.Errorf(codes.NotFound, "There is no grade with that ID")
		return "There is no grade with that ID", err
	}

	//izmeni ocenu
	userGrade := &models.UserGrade{
		ID:    objectID,
		Grade: grade,
	}

	err = service.GradeRepository.UpdateUserGrade(userGrade)
	if err != nil {
		err := status.Errorf(codes.Internal, "Failed to update grade")
		return "Failed to update grade", err
	}

	return "Successfully updated grade", nil
}

func (service *GradeService) DeleteUserGrade(id string) (string, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "Invalid ID format")
		return "Invalid ID format", err
	}

	// da li postoji ocena
	_, err = service.GradeRepository.GetGradeById(objectID)
	if err != nil {
		err := status.Errorf(codes.NotFound, "There is no grade with that ID")
		return "There is no grade with that ID", err
	}

	// brisanje ocene
	_, err = service.GradeRepository.DeleteUserGrade(objectID)
	if err != nil {
		err := status.Errorf(codes.Internal, "Failed to delete grade")
		return "Failed to delete grade", err
	}

	return "Successfully deleted grade", nil
}
