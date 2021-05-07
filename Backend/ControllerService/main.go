package main

import (
	"context"
	"controller/Database"
	"controller/Delivery"
	"controller/Delivery/Service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fiberApp := fiber.New()
	var ConfigDefault = cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length, Authorization, access-control-allow-origin, access-control-allow-headers, access-control-allow-methods",
	}
	fiberApp.Use(cors.New(ConfigDefault))
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
	ud := Service.NewUserData()
	Delivery.SetupRoute(fiberApp, ud)
	fiberApp.Listen(":8079")
}
