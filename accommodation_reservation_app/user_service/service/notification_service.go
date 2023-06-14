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

type NotificationService struct {
	NotificationRepository *repository.NotificationRepository
}

func (service *NotificationService) CreateNotification(notification models.Notification) (string, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	inserterr := service.NotificationRepository.CreateNotification(&notification)
	if inserterr != nil {
		log.Panic(inserterr)
		err := status.Errorf(codes.Internal, "something went wrong")
		return "something went wrong", err
	}

	return "Succesffully added new notification", nil

}

func (service *NotificationService) GetNotificationByUser(userID primitive.ObjectID) ([]models.Notification, error) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var userNotifications []models.Notification
	userNotifications, err := service.NotificationRepository.GetNotificationByUser(userID)
	if err != nil {
		return nil, err
	}

	return userNotifications, nil
}
