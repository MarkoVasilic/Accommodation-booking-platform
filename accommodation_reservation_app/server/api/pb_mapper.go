package api

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/server/models"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/accommodation_service"
)

func mapUser(user *models.User) *pb.User {
	userPb := &pb.User{
		Id:        user.Id.Hex(),
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		Password:  *user.Password,
		Email:     *user.Email,
		Token:     *user.Token,
		Role:      string(*user.Role),
	}
	return userPb
}
