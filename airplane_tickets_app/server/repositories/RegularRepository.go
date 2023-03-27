package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RegularRepository struct {
	UserCollection   *mongo.Collection
	FlightCollection *mongo.Collection
	TicketCollection *mongo.Collection
}

func (repo *RegularRepository) ReduceFlightTickets(flightID primitive.ObjectID, numTickets uint64) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flight := bson.M{"_id": flightID}
	update := bson.M{"$inc": bson.M{"number_of_tickets": -int64(numTickets)}}

	_, err := repo.FlightCollection.UpdateOne(ctx, flight, update)

	return err
}

func (repo *RegularRepository) CreateTicket(ticket *models.Ticket) (primitive.ObjectID, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	id := primitive.NewObjectID()
	ticket.ID = id
	_, err := repo.TicketCollection.InsertOne(ctx, ticket)

	return id, err
}

func (repo *RegularRepository) AddTicketToUser(ticketID primitive.ObjectID, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	userId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	user := bson.M{"_id": userId}
	update := bson.M{"$push": bson.M{"user_tickets": ticketID}}
	fmt.Println(user)

	_, errr := repo.UserCollection.UpdateOne(ctx, user, update)

	return errr
}

func (repo *RegularRepository) GetFlightById(id primitive.ObjectID) (models.Flight, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	var foundFlight models.Flight
	err := repo.FlightCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&foundFlight)
	defer cancel()
	return foundFlight, err
}
