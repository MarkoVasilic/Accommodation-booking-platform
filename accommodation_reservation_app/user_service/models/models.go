package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	Host  Role = "HOST"
	Guest Role = "GUEST"
)

type NotificationType string

const (
	CreateAcc NotificationType = "CREATE_ACC"
	CancelAcc NotificationType = "CANCEL_ACC"
	GradedUsr NotificationType = "GRADED_USR"
	GradedAcc NotificationType = "GRADED_ACC"
	Prominent NotificationType = "PROMINENT"
)

type User struct {
	Id            primitive.ObjectID `bson:"_id" json:"_id"`
	Username      *string            `bson:"username" json:"username" validate:"required,min=4,max=30"`
	FirstName     *string            `bson:"first_name" json:"first_name" validate:"required,min=2,max=30"`
	LastName      *string            `bson:"last_name" json:"last_name"  validate:"required,min=2,max=30"`
	Password      *string            `bson:"password" json:"password"   validate:"required,min=6"`
	Email         *string            `bson:"email" json:"email"      validate:"email,required"`
	Address       *string            `bson:"address" json:"address"      validate:"min=5,max=100"`
	Token         *string            `bson:"token" json:"token"`
	Refresh_Token *string            `bson:"refresh_token" json:"refresh_token"`
	Role          *Role              `bson:"role" json:"role" validate:"required"`
	User_ID       string             `bson:"user_id" json:"user_id"`
}

type UserGrade struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	GuestID     primitive.ObjectID `bson:"guest_id" json:"guest_id" validate:"required"`
	HostID      primitive.ObjectID `bson:"host_id" json:"host_id" validate:"required"`
	Grade       int                `bson:"grade" json:"grade" validate:"required,min=1, max=5"`
	DateOfGrade time.Time          `bson:"date_of_grade" json:"date_of_grade" validate:"required"`
}

type Notification struct {
	ID                 primitive.ObjectID `bson:"_id" json:"_id"`
	UserID             primitive.ObjectID `bson:"user_id" json:"user_id" validate:"required"`
	Type               *NotificationType  `bson:"type" json:"type" validate:"required"`
	Message            *string            `bson:"message" json:"message"      validate:"min=1,max=500"`
	DateOfNotification time.Time          `bson:"date_of_grade" json:"date_of_grade" validate:"required"`
	Seen               bool               `bson:"seen" json:"seen" validate:"required"`
}

type NotificationOn struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id"`
	UserID primitive.ObjectID `bson:"user_id" json:"user_id" validate:"required"`
	Type   *NotificationType  `bson:"type" json:"type" validate:"required"`
	On     bool               `bson:"on" json:"on" validate:"required"`
}
