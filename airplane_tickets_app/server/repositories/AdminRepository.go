package repositories

import (
	"context"
	"time"

	"github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRepository struct {
	UserCollection   *mongo.Collection
	FlightCollection *mongo.Collection
	TicketCollection *mongo.Collection
}

func (repo *AdminRepository) CreateFlight(flight *models.Flight) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	flight.ID = primitive.NewObjectID()
	_, inserterr := repo.FlightCollection.InsertOne(ctx, flight)
	return inserterr
}

func (repo *AdminRepository) GetFlightById(id string) (models.Flight, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var foundflight models.Flight
	objID, err_hex := primitive.ObjectIDFromHex(id)
	if err_hex != nil {
		panic(err_hex)
	}
	err := repo.FlightCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&foundflight)
	defer cancel()
	return foundflight, err
}

func (repo *AdminRepository) DeleteFlightById(id string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	_, err = repo.FlightCollection.DeleteOne(ctx, bson.M{"_id": objID})
	defer cancel()
	return err
}

func (repo *AdminRepository) GetTicketsByFlightId(id string) ([]models.Ticket, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var tickets []models.Ticket
	cursor, err := repo.TicketCollection.Find(ctx, bson.M{"flight": objID})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &tickets); err != nil {
		return nil, err
	}

	return tickets, nil
}

func (repo *AdminRepository) DeleteTicketsByFlightId(flightId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightObjectId, err := primitive.ObjectIDFromHex(flightId)
	if err != nil {
		return err
	}

	_, err = repo.TicketCollection.DeleteMany(
		ctx,
		bson.M{"flight": flightObjectId},
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo *AdminRepository) GetAllUsers() ([]models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	cursor, err := repo.UserCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *AdminRepository) UpdateUserTickets(userID string, tickets []primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = repo.UserCollection.UpdateOne(
		ctx,
		bson.M{"_id": userId},
		bson.M{"$set": bson.M{"user_tickets": tickets}},
	)
	if err != nil {
		return err
	}

	return nil
}
