package storage

import (
	"context"
	"log"

	"github.com/keenfury/api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client

func init() {
	if config.StorageMongo {
		var errNew error
		MongoDB, errNew = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if errNew != nil {
			panic("Unable to connect: mongo")
		}
		errPing := MongoDB.Ping(context.TODO(), nil)
		if errPing != nil {
			log.Panicln("Unable to ping: mongo")
		}
	}
}