package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GradeRepository struct {
	GradeCollection *mongo.Collection
}

func (repo *GradeRepository) CreateUserGrade(userGrade *models.UserGrade) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userGrade.ID = primitive.NewObjectID()
	_, inserterr := repo.GradeCollection.InsertOne(ctx, userGrade)
	return inserterr
}

func (repo *GradeRepository) UpdateUserGrade(userGrade *models.UserGrade) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": userGrade.ID}
	update := bson.M{
		"$set": bson.M{
			"grade": userGrade.Grade,
		},
	}

	_, updateErr := repo.GradeCollection.UpdateOne(ctx, filter, update)
	return updateErr
}

func (repo *GradeRepository) GetGradeById(id primitive.ObjectID) (models.UserGrade, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	var userGrade models.UserGrade
	err := repo.GradeCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&userGrade)
	defer cancel()
	return userGrade, err
}

func (repo *GradeRepository) DeleteUserGrade(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return repo.GradeCollection.DeleteOne(ctx, bson.M{"_id": id})
}

func (repo *GradeRepository) GetAllGuestGrades(guestID primitive.ObjectID) ([]models.UserGrade, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"guest_id": guestID}
	cursor, err := repo.GradeCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var grades []models.UserGrade
	if err := cursor.All(ctx, &grades); err != nil {
		return nil, err
	}

	return grades, nil
}

func (repo *GradeRepository) GetAllUserGrade(hostID primitive.ObjectID) ([]models.UserGrade, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"host_id": hostID}

	cursor, err := repo.GradeCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var grades []models.UserGrade
	if err := cursor.All(ctx, &grades); err != nil {
		return nil, err
	}

	return grades, nil
}
