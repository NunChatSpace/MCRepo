package main

import (
	"context"
	"moonservice/Database"
	"moonservice/Delivery"
	"moonservice/Delivery/Service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fiberApp := fiber.New()
	fiberApp.Use(logger.New())
	var ConfigDefault = cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length, Authorization, access-control-allow-origin, access-control-allow-headers, access-control-allow-methods",
	}
	fiberApp.Use(cors.New(ConfigDefault))
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo/"))
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

	creFile := "./mooncoinrtdb-firebase-adminsdk-yioeo-36db39fffa.json"
	projectID := "mooncoinrtdb"
	dbUrl := "https://mooncoinrtdb-default-rtdb.firebaseio.com"

	rtdbClient, _ := Service.SetupFirebaseService(creFile, projectID, dbUrl)
	ms := Service.NewMoonService(rtdbClient)

	Delivery.SetupRoute(fiberApp, ms)

	fiberApp.Listen(":8080")
}
