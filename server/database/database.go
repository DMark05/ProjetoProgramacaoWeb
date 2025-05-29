package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	mongo_options "go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGOCONNSTRING = os.Getenv("MONGOCONNSTRING")
	MongoClient     *mongo.Client
	err             error
)

type Collections string

const ( //Collections
	Users  Collections = "users"
	Events Collections = "events"
	UC Collections = "uc"
)

func GetCollectionFromMongo(collection Collections) *mongo.Collection {
	switch collection {
	case Users:
		return MongoClient.Database("ProjetoInternet").Collection(string(collection))
	case Events:
		return MongoClient.Database("ProjetoInternet").Collection(string(collection))
	case UC:
		return MongoClient.Database("ProjetoInternet").Collection(string(collection))
	default:
		log.Fatalf("no collection named %s\n", string(collection))
	}
	return nil
}

func Init() error {
	ctx := context.Background()
	MongoClient, err = mongo.Connect(ctx, mongo_options.Client().ApplyURI(MONGOCONNSTRING))
	return err
}
