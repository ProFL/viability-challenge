package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection() *mongo.Client {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()

	url, found := os.LookupEnv("MONGO_URL")

	if found == false {
		log.Panicf("MONGO_URL was not found")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))

	if err != nil {
		log.Panic(err)
	}

	return client
}
