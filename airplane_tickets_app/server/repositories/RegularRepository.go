package repositories

import "go.mongodb.org/mongo-driver/mongo"

type RegularRepository struct {
	UserCollection   *mongo.Collection
	FlightCollection *mongo.Collection
	TicketCollection *mongo.Collection
}
