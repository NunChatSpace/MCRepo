package Service

import (
	"context"
	"controller/Database"
	"controller/Interface"
	"controller/Model"
	"net/http"
	"sort"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// Sort by `BuyDate` field descending
	findOptions.SetSort(bson.D{primitive.E{Key: "BuyDate", Value: -1}})

	cursor, err := dbStruct.UserLogCollection.Find(context, bson.D{}, findOptions)

	if err != nil {
		resp = Model.ResponseModel{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		// fmt.Println(resp)
		return resp
	}

	var dataContent []Model.UserLog
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

	sort.SliceStable(dataContent, func(i, j int) bool {
		dateI, _ := time.Parse("2006-01-02 15:04:05", dataContent[i].BuyDate)
		dateY, _ := time.Parse("2006-01-02 15:04:05", dataContent[j].BuyDate)

		return dateY.Before(dateI)
	})

	resp = Model.ResponseModel{
		Status:     http.StatusOK,
		Message:    "Getting data is successfully",
		DataLength: len(dataContent),
		Data:       dataContent,
	}
	// fmt.Println(resp)
	return resp
}
