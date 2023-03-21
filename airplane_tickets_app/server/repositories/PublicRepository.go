package repositories

import (
	"context"
	"log"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/models"
	generate "github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/tokens"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PublicRepository struct {
	UserCollection   *mongo.Collection
	FlightCollection *mongo.Collection
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func (repo *PublicRepository) CreateUser(user *models.User) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	password := HashPassword(*user.Password)
	user.Password = &password
	user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_ID = user.ID.Hex()
	token, refreshtoken, _ := generate.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, user.User_ID, models.Role(*user.Role))
	user.Token = &token
	user.Refresh_Token = &refreshtoken
	user.UserTickets = make([]primitive.ObjectID, 0)
	_, inserterr := repo.UserCollection.InsertOne(ctx, user)
	return inserterr
}

func (repo *PublicRepository) CountByEmail(email string) (int64, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return repo.UserCollection.CountDocuments(ctx, bson.M{"email": email})
}

func (repo *PublicRepository) GetUserByEmail(email string) (models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	var founduser models.User
	err := repo.UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&founduser)
	defer cancel()
	return founduser, err
}

func (repo *PublicRepository) GetUserById(id string) (models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	var founduser models.User
	err := repo.UserCollection.FindOne(ctx, bson.M{"user_id": id}).Decode(&founduser)
	defer cancel()
	return founduser, err
}
