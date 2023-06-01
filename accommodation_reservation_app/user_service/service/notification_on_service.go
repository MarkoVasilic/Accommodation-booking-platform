package service

import (
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

func (service *NotificationOnService) InitializeNotificationsOn(userID primitive.ObjectID) error {
	notificationTypes := []models.NotificationType{models.CreateAcc, models.CancelAcc, models.GradedUsr, models.GradedAcc, models.Prominent}

	for _, notificationType := range notificationTypes {
		err := service.NotificationOnRepository.CreateNotificationOn(createNotification(userID, notificationType))
		if err != nil {
			err := status.Errorf(codes.Internal, "something went wrong")
			return err
		}
	}
	return nil
}
