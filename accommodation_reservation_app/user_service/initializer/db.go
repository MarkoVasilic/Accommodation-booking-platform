package initializer

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE = "user_db"
)

func ConnectToDatabase(host, port string) *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s/", host, port)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(options)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
	var collection *mongo.Collection = client.Database(DATABASE).Collection("Users")
	return collection
}

func GradeCollection(client *mongo.Client) *mongo.Collection {
	var collection *mongo.Collection = client.Database(DATABASE).Collection("UserGrades")
	return collection
}

func NotificationCollection(client *mongo.Client) *mongo.Collection {
	var collection *mongo.Collection = client.Database(DATABASE).Collection("Notifications")
	return collection
}

func NotificationOnCollection(client *mongo.Client) *mongo.Collection {
	var collection *mongo.Collection = client.Database(DATABASE).Collection("NotificationsOn")
	return collection
}
