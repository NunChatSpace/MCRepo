package Interface

import "moonservice/Model"

type IRTDB interface {
	GetMoonCoinFromRTDB() (mcModel Model.RTDBMoonCoinModel, err error)
	DecreaseMoonCoinToRTDB(moonCoin float64, exchangeRate float64) (mcModel Model.RTDBMoonCoinModel, err error)
}
