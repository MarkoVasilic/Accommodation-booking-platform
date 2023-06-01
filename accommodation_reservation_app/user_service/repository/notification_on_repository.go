package repository

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
