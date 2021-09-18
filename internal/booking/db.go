package booking

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

// config mongo
const (
	dbName   = "pttrainer"
	mongoURI = "mongodb+srv://pttrainer:NmSJWnnBmApV5sEu@maincluster.gkfe6.mongodb.net/pttrainer?authSource=admin&replicaSet=atlas-u555p7-shard-0&w=majority&readPreference=primary&appname=MongoDB%20Compass&retryWrites=true&ssl=true"
	// mongoURI = "mongodb://localhost:27017/" + dbName
)

func Connect() *MongoInstance {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Error(fmt.Sprintf("mongo.NewClient, err: %v", err))
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Error(fmt.Sprintf("client.Ping, err: %v", err))
		return nil
	}

	db := client.Database(dbName)
	if err != nil {
		log.Error("client.Database, err: %v", err)
		return nil
	}

	return &MongoInstance{
		Client: client,
		Db:     db,
	}
}
