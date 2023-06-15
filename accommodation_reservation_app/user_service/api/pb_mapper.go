package api

import (
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
	userGradePb := &pb.UserGrade{
		ID:          userGrade.ID.Hex(),
		GuestID:     userGrade.GuestID.Hex(),
		HostID:      userGrade.HostID.Hex(),
		Grade:       float64(userGrade.Grade),
		DateOfGrade: dateOfGrade,
	}
	return userGradePb
}

func mapNotification(notification *models.Notification) *pb.Notification {
	dateOfNotification := timestamppb.New(notification.DateOfNotification)
	notificationPb := &pb.Notification{
		Id:                 notification.ID.Hex(),
		UserId:             notification.UserID.Hex(),
		Type:               string(*notification.Type),
		Message:            *notification.Message,
		DateOfNotification: dateOfNotification,
		Seen:               notification.Seen,
	}
	return notificationPb
}

func mapNotificationOn(notification *models.NotificationOn) *pb.NotificationOn {
	notificationPb := &pb.NotificationOn{
		Id:     notification.ID.Hex(),
		UserId: notification.UserID.Hex(),
		Type:   string(*notification.Type),
		On:     notification.On,
	}
	return notificationPb
}

func mapUserGradeDetails(gradeDetails *models.UserGradeDetails) *pb.UserGradeDetails {
	dateOfGrade := timestamppb.New(gradeDetails.DateOfGrade)
	userGradeDetailsPb := &pb.UserGradeDetails{
		GuestFirstName: gradeDetails.GuestFirstName,
		GuestLastName:  gradeDetails.GuestLastName,
		HostFirstName:  gradeDetails.HostFirstName,
		HostLastName:   gradeDetails.GuestLastName,
		Grade:          float64(gradeDetails.Grade),
		DateOfGrade:    dateOfGrade,
	}
	return userGradeDetailsPb
}

func mapHost(hostDetails *models.HostDetails) *pb.HostDetails {
	hostDetailsPb := &pb.HostDetails{
		Id:        hostDetails.Id.Hex(),
		FirstName: hostDetails.FirstName,
		LastName:  hostDetails.LastName,
	}
	return hostDetailsPb
}
