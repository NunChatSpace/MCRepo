package Model

type ResponseModel struct {
	Status     int         `json:"Status"`
	Message    string      `json:"Message"`
	DataLength int         `json:"DataLength"`
	Data       interface{} `json:"Data"`
}

type UserLog struct {
	BuyDate  string  `json:"BuyDate"`
	Username string  `json:"Username"`
	THBT     float64 `json:"THBT"`
	MOON     float64 `json:"MOON"`
	Rate     string  `json:"Rate"`
}
