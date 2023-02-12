package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	Host   string
	Port   string
	DBName string
	AppEnv string
}

type Repository struct {
	Client       *mongo.Client
	DbConnection *mongo.Database
}

func (config *Config) Connect() (Repository, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(config.Host))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")

	if err != nil {
		return Repository{Client: client, DbConnection: client.Database(config.DBName)}, err
	}

	return Repository{Client: client, DbConnection: client.Database(config.DBName)}, nil
}
