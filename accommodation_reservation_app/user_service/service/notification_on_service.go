package service

import (
	"context"
	"log"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/models"
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NotificationOnService struct {
	NotificationOnRepository *repository.NotificationOnRepository
}

func createNotification(userID primitive.ObjectID, notificationType models.NotificationType) *models.NotificationOn {
	return &models.NotificationOn{
		ID:     primitive.NewObjectID(),
		UserID: userID,
		Type:   &notificationType,
		On:     true,
	}
}

func (service *NotificationOnService) InitializeNotificationsOn(user models.User) error {
	if *user.Role == "HOST" {
		notificationTypes := []models.NotificationType{models.CreateAcc, models.CancelAcc, models.GradedUsr, models.GradedAcc, models.Prominent}

		for _, notificationType := range notificationTypes {
			err := service.NotificationOnRepository.CreateNotificationOn(createNotification(user.Id, notificationType))
			if err != nil {
				err := status.Errorf(codes.Internal, "something went wrong")
				return err
			}
		}
	} else {
		err := service.NotificationOnRepository.CreateNotificationOn(createNotification(user.Id, models.ReservationReact))
		if err != nil {
			err := status.Errorf(codes.Internal, "something went wrong")
			return err
		}
	}

	return nil
}

func (service *NotificationOnService) UpdateNotificationOn(notification_on models.NotificationOn, user_id string) (string, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	userId, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		err := status.Errorf(codes.Internal, "Id conversion failed")
		return "Id conversion failed", err
	}

	foundnotification, founderr := service.NotificationOnRepository.GetNotificationByUserAndType(userId, notification_on.Type)
	if founderr != nil {
		err := status.Errorf(codes.Internal, "something went wrong")
		return "something went wrong", err
	}

	notification_on.ID = foundnotification.ID
	notification_on.UserID = userId

	inserterr := service.NotificationOnRepository.UpdateNotificationOn(&notification_on)
	if inserterr != nil {
		log.Panic(err)
		err := status.Errorf(codes.Internal, "something went wrong")
		return "something went wrong", err
	}
	return "Succesffully updated notification on", nil

}

func (service *NotificationOnService) GetNotificationOnByUser(userID primitive.ObjectID) ([]models.NotificationOn, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var userNotificationsOn []models.NotificationOn
	userNotificationsOn, err := service.NotificationOnRepository.GetNotificationOnByUser(userID)
	if err != nil {
		return nil, err
	}

	return userNotificationsOn, nil
}
