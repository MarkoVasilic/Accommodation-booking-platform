package api

import (
	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/models"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
)

func mapUser(user *models.User) *pb.User {
	userPb := &pb.User{
		Id:        user.Id.Hex(),
		Username:  *user.Username,
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		Password:  *user.Password,
		Email:     *user.Email,
		Address:   *user.Address,
		Token:     *user.Token,
		Role:      string(*user.Role),
	}
	return userPb
}
