package Model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBStruct struct {
	MongoDB           *mongo.Database
	UserLogCollection *mongo.Collection
	MongoDBClient     *mongo.Client
}
