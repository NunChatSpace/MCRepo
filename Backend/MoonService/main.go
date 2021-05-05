package main

import (
	"context"
	"moonservice/Database"
	"moonservice/Delivery"
	"moonservice/Delivery/Service"
	"time"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fiberApp := fiber.New()
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)
	Database.SetupMongoDB(client)

	rtdbClient := Service.SetupFirebaseService()
	ms := Service.NewMoonService(rtdbClient)

	Delivery.SetupRoute(fiberApp, ms)

	fiberApp.Listen(":8080")
}
