package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type GradeRepository struct {
	GradeCollection *mongo.Collection
}
