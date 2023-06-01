package service

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/repository"
)

type NotificationService struct {
	NotificationRepository *repository.NotificationRepository
}
