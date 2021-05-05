package Service

import (
	"context"
	"fmt"
	"moonservice/Database"
	"moonservice/Interface"
	"moonservice/Model"
	"net/http"
	"time"

	"github.com/gofiber/fiber"
)

type MoonService struct {
	RTDB Interface.IRTDB
}

func (ms *MoonService) BuyMOON(c *fiber.Ctx) (resp Model.ResponseModel) {
	respData := Model.ResponseDataModel{}
	body := new(Model.RequestModel)
	err := c.BodyParser(body)
	if err != nil {
		resp = Model.ResponseModel{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return resp
	}

	mcModel, err := ms.RTDB.GetMoonCoinFromRTDB()
	if err != nil {
		resp = Model.ResponseModel{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return resp
	}
	mcc := new(MoonCoinCalculator)

	buyLog, realMoonCoin, thbtUsed := mcc.HowManyMoonCoinForMe(mcModel, body)
	if len(buyLog) == 0 {

		resp = Model.ResponseModel{
			Status:  500,
			Message: "Exchange rate is exceed the slippage telorance",
		}
		c.JSON(resp)

		return resp
	}
	step := 100 - int(((mcModel.Remaining - realMoonCoin) / 10))
	exRate := mcc.CurrentExchangeRate(step)
	_, err = ms.RTDB.DecreaseMoonCoinToRTDB(realMoonCoin, exRate)
	if err != nil {
		resp = Model.ResponseModel{
			Status:  500,
			Message: err.Error(),
		}
		return resp
	}

	respData = Model.ResponseDataModel{
		BuyingStaus: "Success",
		MOON:        realMoonCoin,
		THBT:        thbtUsed,
	}
	resp = Model.ResponseModel{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    respData,
	}

	ms.UpdateMoon(time.Now(), body.Username, body.BuyWith, buyLog)

	return resp
}

func (ms *MoonService) UpdateMoon(date time.Time, username string, thbt float64, buyLog []Model.BuyLogModel) {
	dbStruct := Database.GetMongoDBStruct()
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	for i := range buyLog {
		context := context.Background()
		exchangeRateTxt := fmt.Sprintf("1 MOON = %.15f | %.15f", 1/buyLog[i].ExchangeRate, buyLog[i].ExchangeRate)
		log := Model.UserLog{
			BuyDate:  currentTime,
			Username: username,
			THBT:     buyLog[i].THBTUsed,
			MOON:     buyLog[i].MoonTook,
			Rate:     exchangeRateTxt,
		}
		dbStruct.UserLogCollection.InsertOne(context, log)
	}
}

func NewMoonService(rtdb Interface.IRTDB) Interface.IMoonService {
	return &MoonService{
		RTDB: rtdb,
	}
}
