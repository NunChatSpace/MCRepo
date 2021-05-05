package Database

import (
	"moonservice/Model"

	"go.mongodb.org/mongo-driver/mongo"
)

var mongoDB Model.MongoDBStruct

func GetMongoDBStruct() Model.MongoDBStruct {
	return mongoDB
}

func SetupMongoDB(client *mongo.Client) {
	dbTemp := client.Database("MoonCoin")
	mongoDB = Model.MongoDBStruct{
		MongoDB:           dbTemp,
		UserLogCollection: dbTemp.Collection("UserLog"),
		MongoDBClient:     client,
	}
}
