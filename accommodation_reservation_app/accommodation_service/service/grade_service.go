package service

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GradeService struct {
	GradeRepository *repository.GradeRepository
}

func (service *GradeService) CreateAccommodationGrade(accommodationGrade models.AccommodationGrade) (string, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	validationErr := Validate.Struct(accommodationGrade)
	if validationErr != nil {
		err := status.Errorf(codes.InvalidArgument, "accommodation grade fields are not valid")
		return "Accommodation grade fields are not valid", err
	}

	inserterr := service.GradeRepository.CreateAccommodationGrade(&accommodationGrade)
	if inserterr != nil {
		err := status.Errorf(codes.Internal, "failed to create accommodation grade")
		return "Failed to create accommodation grade", err
	}

	return "Successfully created accommodation grade", nil
}

func (service *GradeService) DeleteAccommodationGrade(id string) (string, error) {
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
	_, err = service.GradeRepository.DeleteAccommodationGrade(objectID)
	if err != nil {
		err := status.Errorf(codes.Internal, "Failed to delete grade")
		return "Failed to delete grade", err
	}

	return "Successfully deleted grade", nil
}

func (service *GradeService) GetAllAccommodationGuestGrades(guestID string) ([]models.AccommodationGrade, error) {
	objectID, err := primitive.ObjectIDFromHex(guestID)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "Invalid ID format")
		return nil, err
	}

	grades, err := service.GradeRepository.GetAllAccommodationGuestGrades(objectID)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "Failed to retrieve guests grades")
		return nil, err
	}

	return grades, nil
}

func (service *GradeService) GetAllAccommodationGrade(accommodationID string) ([]models.AccommodationGrade, error) {
	objectID, err := primitive.ObjectIDFromHex(accommodationID)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "Invalid ID format")
		return nil, err
	}

	grades, err := service.GradeRepository.GetAllAccommodationGrade(objectID)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "Failed to retrieve accommodation grades")
		return nil, err
	}

	return grades, nil
}

func (service *GradeService) UpdateAccommodationGrade(grade int, id string) (string, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err := status.Errorf(codes.InvalidArgument, "Invalid ID format")
		return "Invalid ID format", err
	}

	//provera da li postoji ocena
	accommodationGrade, err := service.GradeRepository.GetGradeById(objectID)
	if err != nil {
		err := status.Errorf(codes.NotFound, "There is no grade with that ID")
		return "There is no grade with that ID", err
	}

	accommodationGrade.Grade = grade
	err = service.GradeRepository.UpdateAccommodationGrade(&accommodationGrade)
	if err != nil {
		err := status.Errorf(codes.Internal, "Failed to update grade")
		return "Failed to update grade", err
	}

	return "Successfully  updated grade", nil
}
