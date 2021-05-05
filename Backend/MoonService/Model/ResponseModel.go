package Model

type ResponseModel struct {
	Status  int               `json:"Status"`
	Message string            `json:"Message"`
	Data    ResponseDataModel `json:"Data"`
}

type ResponseDataModel struct {
	BuyingStaus string  `json:"BuyingStaus"`
	MOON        float64 `josn:"MOON"`
	THBT        float64 `json:"THBT"`
}
