package Service

import (
	"math"
	"moonservice/Model"
)

type MoonCoinCalculator struct {
}

func (mcc *MoonCoinCalculator) HowManyMoonCoinForMe(mcModel Model.RTDBMoonCoinModel, body *Model.RequestModel) (buyLog []Model.BuyLogModel, moonCoin float64, THBTUsed float64) {
	// Start with 1 ... 100
	baseIndex := 100 - int((mcModel.Remaining / 10))
	buyWith := body.BuyWith
	moonCoin = 0.0
	for i := baseIndex; i <= 100; i++ {
		baseMoonCoinStep := 1000 - (10 * i)

		baseExchangeRate := mcc.CurrentExchangeRate(i)
		MaxTHBT := 0.0
		if !((1/baseExchangeRate) >= body.CurrentExchangeRate*(body.SlippageRateMin) && (1/baseExchangeRate) <= body.CurrentExchangeRate*(body.SlippageRateMax)) {
			break
		}

		if float64(mcModel.Remaining-float64(baseMoonCoinStep)) == 0.0 {
			MaxTHBT = baseExchangeRate * 10.0
		} else {
			MaxTHBT = baseExchangeRate * float64(mcModel.Remaining-float64(baseMoonCoinStep))
		}
		if buyWith <= MaxTHBT {
			// fmt.Printf("Last Round : MaxTHBT is %.15f, buy with %.15f ", MaxTHBT, buyWith)
			moonTook := (buyWith * (1 / baseExchangeRate))
			THBTUsed = THBTUsed + buyWith
			buyLog = append(buyLog, Model.BuyLogModel{
				ExchangeRate: baseExchangeRate,
				MoonTook:     moonTook,
				THBTUsed:     buyWith,
			})
			moonCoin = moonCoin + moonTook
			mcModel.Remaining = mcModel.Remaining - moonTook
			// fmt.Printf("===> %.15f\n", moonTook)
			break
		} else {
			// fmt.Printf("MaxTHBT is %.15f, buy with %.15f ", MaxTHBT, MaxTHBT)
			moonTook := (MaxTHBT * (1 / baseExchangeRate))
			buyWith = buyWith - MaxTHBT
			THBTUsed = THBTUsed + MaxTHBT
			buyLog = append(buyLog, Model.BuyLogModel{
				ExchangeRate: baseExchangeRate,
				MoonTook:     moonTook,
				THBTUsed:     THBTUsed,
			})
			moonCoin = moonCoin + moonTook
			mcModel.Remaining = mcModel.Remaining - 10.0
			// fmt.Printf("===> %.15f\n", moonTook)
		}
	}

	return buyLog, moonCoin, THBTUsed
}

func (mcc *MoonCoinCalculator) CurrentExchangeRate(step int) float64 {
	// For every 10 coins of MOON sold, The price will rise by 10%
	ex := (50.0 * math.Pow(1.1, float64(step)))
	return ex
}
