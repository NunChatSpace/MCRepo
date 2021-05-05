package Model

type RequestModel struct {
	Username            string  `json:"Username"`
	BuyWith             float64 `json:"BuyWith"`
	CurrentExchangeRate float64 `json:"CurrentExchangeRate"`
	SlippageRateMin     float64 `json:"SlippageRateMin"` // if 5% SlippageRate = 0.95
	SlippageRateMax     float64 `json:"SlippageRateMax"` // if 5% SlippageRate = 1.05
}
