package repository

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NotificationOnRepository struct {
	NotificationOnCollection *mongo.Collection
}

func (repo *NotificationOnRepository) CreateNotificationOn(notification_on *models.NotificationOn) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	notification_on.ID = primitive.NewObjectID()
	_, inserterr := repo.NotificationOnCollection.InsertOne(ctx, notification_on)
	return inserterr
}

func (repo *NotificationOnRepository) GetNotificationByUserAndType(userId primitive.ObjectID, notType *models.NotificationType) (models.NotificationOn, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var foundnotification models.NotificationOn
	err := repo.NotificationOnCollection.FindOne(ctx, bson.M{"user_id": userId, "type": notType}).Decode(&foundnotification)
	return foundnotification, err
}

func (repo *NotificationOnRepository) UpdateNotificationOn(notification_on *models.NotificationOn) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": notification_on.ID}
	update := bson.M{
		"$set": bson.M{
			"on": notification_on.On,
		},
	}
	options := options.Update().SetUpsert(false)
	_, inserterr := repo.NotificationOnCollection.UpdateOne(ctx, filter, update, options)
	return inserterr
}
