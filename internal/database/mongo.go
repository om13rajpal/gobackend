package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/om13rajpal/gobackend/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	UserCollection *mongo.Collection
)

func ConnectMongo() {
	clientOptions := options.Client().ApplyURI(config.MONGO_URI)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal("Could not connect to MONGO DB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not ping MONGO DB")
	}

	fmt.Println("Connected to MONGO DB")

	UserCollection = client.Database("golang").Collection("users")
}
