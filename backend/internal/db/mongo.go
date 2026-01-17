package db

import (
	"be-request-insident/internal/config"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConn struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func ConnectMongo() *MongoConn {
	mongoHost := config.GetEnvVariable("MONGO_HOST")
	mongoPort := config.GetEnvVariable("MONGO_PORT")
	mongoUser := config.GetEnvVariable("MONGO_INITDB_ROOT_USERNAME")
	mongoPassword := config.GetEnvVariable("MONGO_INITDB_ROOT_PASSWORD")
	dbName := config.GetEnvVariable("MONGO_INITDB_DATABASE")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := "mongodb://" + mongoUser + ":" + mongoPassword + "@" + mongoHost + ":" + mongoPort + "/?authSource=admin"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	log.Println("Connected to MongoDB successfully")
	return &MongoConn{
		Client: client,
		DB:     client.Database(dbName),
	}
}
