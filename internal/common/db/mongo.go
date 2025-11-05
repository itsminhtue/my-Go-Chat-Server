package db

import (
	"context"
	"fmt"
	"goChat/internal/common/config"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoClient *mongo.Client

func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, _ := mongo.Connect(options.Client().ApplyURI(config.GetEnv("MONGO_URI", "mongodb://localhost:27017")))
	if err := client.Ping(ctx, nil); err != nil {
		panic(fmt.Errorf("%v", err))
	}

	MongoClient = client
	return client.Database("go_chat_db")

}
