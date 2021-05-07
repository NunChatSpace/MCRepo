package Service

import (
	"context"
	"controller/Database"
	"controller/Interface"
	"controller/Model"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type UserData struct {
}

func NewUserData() Interface.IUserData {
	return &UserData{}
}

func (ud *UserData) GetHistory(ctx *fiber.Ctx) (resp Model.ResponseModel) {
	dbStruct := Database.GetMongoDBStruct()
	context := context.Background()

	cursor, err := dbStruct.UserLogCollection.Find(context, bson.D{})
	if err != nil {
		resp = Model.ResponseModel{
			Status:  http.StatusInternalServerError,
			Message: "Get failed",
		}
	}

	var dataContent []interface{}

	for cursor.Next(context) {
		var content Model.UserLog
		err := cursor.Decode(&content)
		if err != nil {
			resp = Model.ResponseModel{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
			ctx.JSON(resp)
		}
		dataContent = append(dataContent, content)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	resp = Model.ResponseModel{
		Status:     http.StatusOK,
		Message:    "Getting data is successfully",
		DataLength: len(dataContent),
		Data:       dataContent,
	}
	return resp
}
