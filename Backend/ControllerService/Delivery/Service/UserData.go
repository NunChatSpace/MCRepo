package Service

import (
	"context"
	"controller/Database"
	"controller/Interface"
	"controller/Model"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserData struct {
}

func NewUserData() Interface.IUserData {
	return &UserData{}
}

func (ud *UserData) GetHistory(ctx *fiber.Ctx) (resp Model.ResponseModel) {
	dbStruct := Database.GetMongoDBStruct()
	context := context.Background()
	findOptions := options.Find()
	// Sort by `price` field descending
	findOptions.SetSort(bson.D{{"BuyDate", -1}})
	cursor, err := dbStruct.UserLogCollection.Find(context, bson.D{}, findOptions)

	if err != nil {
		resp = Model.ResponseModel{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		// fmt.Println(resp)
		return resp
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
			// fmt.Println(resp)
			ctx.JSON(resp)
		}
		dataContent = append(dataContent, content)
	}

	if err := cursor.Err(); err != nil {
		resp = Model.ResponseModel{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		// fmt.Println(resp)
		return resp
	}
	resp = Model.ResponseModel{
		Status:     http.StatusOK,
		Message:    "Getting data is successfully",
		DataLength: len(dataContent),
		Data:       dataContent,
	}
	// fmt.Println(resp)
	return resp
}
