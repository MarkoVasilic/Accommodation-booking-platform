package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationRepository struct {
	NotificationCollection *mongo.Collection
}
