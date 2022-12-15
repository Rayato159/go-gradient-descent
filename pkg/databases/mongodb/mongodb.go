package databases

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"www.github.com/Rayato159/go-gradient-descent/configs"
)

func NewMongoDbConnect(cfg configs.MongoDb) (*mongo.Client, error) {
	// Set time out for connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set uri
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)
	// Connect to mongoDb
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil || client == nil {
		return nil, fmt.Errorf("error, can't connect to mongodb with an error: %v", err)
	}
	// Ping to test connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, fmt.Errorf("error, can't ping to mongodb with an error: %v", err)
	}
	defer log.Println("🍃mongodb has been connected, have a nice day!")
	return client, nil
}

func MongoDbDisconnect(client *mongo.Client) {
	ctx := context.Background()
	log.Println("🍃mongodb has been disconnected, good bye")

	// Disconnect from database
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
