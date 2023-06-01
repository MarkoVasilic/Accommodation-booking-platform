package repository

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/models"
	generate "github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/token"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	UserCollection *mongo.Collection
}

func (repo *UserRepository) GetUserById(id primitive.ObjectID) (models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	var founduser models.User
	err := repo.UserCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&founduser)
	defer cancel()
	return founduser, err
}

func (repo *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	var founduser models.User
	err := repo.UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&founduser)
	defer cancel()
	return founduser, err
}

func (repo *UserRepository) GetUserByUsername(username string) (models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	var founduser models.User
	err := repo.UserCollection.FindOne(ctx, bson.M{"username": username}).Decode(&founduser)
	defer cancel()
	return founduser, err
}

func (repo *UserRepository) CreateUser(user *models.User) (error, primitive.ObjectID) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user.Id = primitive.NewObjectID()
	user.User_ID = user.Id.Hex()
	token, refreshtoken, _ := generate.TokenGenerator(*user.Email, *user.FirstName, *user.LastName, user.User_ID, models.Role(*user.Role))
	user.Token = &token
	user.Refresh_Token = &refreshtoken
	_, inserterr := repo.UserCollection.InsertOne(ctx, user)
	return inserterr, user.Id
}

func (repo *UserRepository) UpdateUser(user *models.User) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": user.Id}
	update := bson.M{
		"$set": bson.M{
			"username":   *user.Username,
			"first_name": *user.FirstName,
			"last_name":  *user.LastName,
			"password":   *user.Password,
			"email":      *user.Email,
			"address":    *user.Address,
		},
	}
	options := options.Update().SetUpsert(false)
	_, inserterr := repo.UserCollection.UpdateOne(ctx, filter, update, options)
	return inserterr
}

func (repo *UserRepository) DeleteUser(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return repo.UserCollection.DeleteOne(ctx, bson.M{"_id": id})
}

func (repo *UserRepository) CountByEmail(email string) (int64, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return repo.UserCollection.CountDocuments(ctx, bson.M{"email": email})
}

func (repo *UserRepository) CountByUsername(username string) (int64, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return repo.UserCollection.CountDocuments(ctx, bson.M{"username": username})
}
