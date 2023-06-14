package repository

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/accommodation_service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GradeRepository struct {
	GradeCollection *mongo.Collection
}

func (repo *GradeRepository) CreateAccommodationGrade(accommodationGrade *models.AccommodationGrade) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	accommodationGrade.ID = primitive.NewObjectID()
	_, inserterr := repo.GradeCollection.InsertOne(ctx, accommodationGrade)
	return inserterr
}

func (repo *GradeRepository) DeleteAccommodationGrade(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return repo.GradeCollection.DeleteOne(ctx, bson.M{"_id": id})
}

func (repo *GradeRepository) GetGradeById(id primitive.ObjectID) (models.AccommodationGrade, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	var accommodationGrade models.AccommodationGrade
	err := repo.GradeCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&accommodationGrade)
	defer cancel()
	return accommodationGrade, err
}

func (repo *GradeRepository) GetAllAccommodationGuestGrades(guestID primitive.ObjectID) ([]models.AccommodationGrade, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"guest_id": guestID}
	cursor, err := repo.GradeCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var grades []models.AccommodationGrade
	if err := cursor.All(ctx, &grades); err != nil {
		return nil, err
	}

	return grades, nil
}

func (repo *GradeRepository) GetAllAccommodationGrade(accommodationID primitive.ObjectID) ([]models.AccommodationGrade, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"accommodation_id": accommodationID}

	cursor, err := repo.GradeCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var grades []models.AccommodationGrade
	if err := cursor.All(ctx, &grades); err != nil {
		return nil, err
	}

	return grades, nil
}

func (repo *GradeRepository) UpdateAccommodationGrade(accommodationGrade *models.AccommodationGrade) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": accommodationGrade.ID}
	update := bson.M{
		"$set": bson.M{
			"grade":         accommodationGrade.Grade,
			"date_of_grade": time.Now(),
		},
	}

	_, updateErr := repo.GradeCollection.UpdateOne(ctx, filter, update)
	return updateErr
}
