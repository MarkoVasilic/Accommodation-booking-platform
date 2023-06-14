package repository

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationRepository struct {
	NotificationCollection *mongo.Collection
}

func (repo *NotificationRepository) CreateNotification(notification *models.Notification) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	notification.ID = primitive.NewObjectID()
	_, inserterr := repo.NotificationCollection.InsertOne(ctx, notification)
	return inserterr
}

func (repo *NotificationRepository) GetNotificationByUser(userID primitive.ObjectID) ([]models.Notification, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userID}
	cursor, err := repo.NotificationCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var userNotifications []models.Notification
	if err = cursor.All(ctx, &userNotifications); err != nil {
		return nil, err
	}

	return userNotifications, nil
}
