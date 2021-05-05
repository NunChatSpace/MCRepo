package Model

type UserLog struct {
	BuyDate  string  `json:"BuyDate"`
	Username string  `json:"Username"`
	THBT     float64 `json:"THBT"`
	MOON     float64 `json:"MOON"`
	Rate     string  `json:"Rate"`
}
