package Service

import (
	"moonservice/Model"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
)

func Test_HowManyMoonCoinForMe_Success(t *testing.T) {
	mcc := new(MoonCoinCalculator)
	mcModel := Model.RTDBMoonCoinModel{
		ExchangeRate: 0.02,
		Remaining:    1000,
	}
	body := Model.RequestModel{
		Username:            "test",
		BuyWith:             50.0,
		CurrentExchangeRate: 0.02,
		SlippageRateMin:     0,
		SlippageRateMax:     100,
	}

	buyLog, moonCoin, THBTUsed := mcc.HowManyMoonCoinForMe(mcModel, &body)

	utils.AssertEqual(t, 1, len(buyLog))
	utils.AssertEqual(t, float64(1), moonCoin)
	utils.AssertEqual(t, float64(50), THBTUsed)
}

func Test_HowManyMoonCoinForMe_Success2(t *testing.T) {
	mcc := new(MoonCoinCalculator)
	mcModel := Model.RTDBMoonCoinModel{
		ExchangeRate: 0.02,
		Remaining:    1000,
	}
	body := Model.RequestModel{
		Username:            "test",
		BuyWith:             501,
		CurrentExchangeRate: 0.02,
		SlippageRateMin:     0,
		SlippageRateMax:     100,
	}

	buyLog, moonCoin, THBTUsed := mcc.HowManyMoonCoinForMe(mcModel, &body)

	utils.AssertEqual(t, 2, len(buyLog))
	utils.AssertEqual(t, float64(10.018181818181818), moonCoin)
	utils.AssertEqual(t, float64(501), THBTUsed)
}

func Test_CurrentExchangeRate_50(t *testing.T) {
	mcc := new(MoonCoinCalculator)

	currEx := mcc.CurrentExchangeRate(0)

	utils.AssertEqual(t, int(50), int(currEx))
}

func Test_CurrentExchangeRate_55(t *testing.T) {
	mcc := new(MoonCoinCalculator)

	currEx := mcc.CurrentExchangeRate(1)

	utils.AssertEqual(t, int(55), int(currEx))
}

func Test_CurrentExchangeRate_18Step(t *testing.T) {
	mcc := new(MoonCoinCalculator)

	currEx := mcc.CurrentExchangeRate(18)

	utils.AssertEqual(t, int(277), int(currEx))
}
