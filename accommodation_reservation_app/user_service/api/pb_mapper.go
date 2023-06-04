package api

import (
	"strconv"

	"github.com/MarkoVasilic/Accommodation-booking-platform/accomodation_reservation_app/user_service/models"
	pb "github.com/MarkoVasilic/Accommodation-booking-platform/common/proto/user_service"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func mapGrade(userGrade *models.UserGrade) *pb.UserGrade {
	dateOfGrade := timestamppb.New(userGrade.DateOfGrade)
	grade := strconv.Itoa(userGrade.Grade)
	userGradePb := &pb.UserGrade{
		ID:          userGrade.ID.Hex(),
		GuestID:     userGrade.GuestID.Hex(),
		HostID:      userGrade.HostID.Hex(),
		Grade:       grade,
		DateOfGrade: dateOfGrade,
	}
	return userGradePb
}
