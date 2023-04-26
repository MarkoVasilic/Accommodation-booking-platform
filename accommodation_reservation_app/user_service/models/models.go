package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	Host  Role = "HOST"
	Guest Role = "GUEST"
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
