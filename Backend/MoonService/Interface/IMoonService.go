package Interface

import (
	"moonservice/Model"
	"time"

	"github.com/gofiber/fiber"
)

type IMoonService interface {
	// Recieve API
	// 1. Get Moon by THBT
	// 2. If not over sliprate it's OK
	//    2.1 Update Realtime DB
	//    2.2 Update to DB

	BuyMOON(c *fiber.Ctx) (resp Model.ResponseModel)
	UpdateMoon(date time.Time, username string, thbt float64, buyLog []Model.BuyLogModel)
}
