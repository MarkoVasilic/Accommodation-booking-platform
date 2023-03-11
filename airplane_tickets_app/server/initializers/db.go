package initializers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDatabase() *mongo.Client {
	dsn := os.Getenv("DB_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect to mongodb")
	} else {
		fmt.Println("Successfully Connected to the mongodb")
	}

	return client
}

func UserCollection(client *mongo.Client) *mongo.Collection {
	var collection *mongo.Collection = client.Database("airplane_tickets_db").Collection("Users")
	return collection
}

func FlightCollection(client *mongo.Client) *mongo.Collection {
	var collection *mongo.Collection = client.Database("airplane_tickets_db").Collection("Flights")
	return collection
}

func TicketCollection(client *mongo.Client) *mongo.Collection {
	var collection *mongo.Collection = client.Database("airplane_tickets_db").Collection("Tickets")
	return collection
}
