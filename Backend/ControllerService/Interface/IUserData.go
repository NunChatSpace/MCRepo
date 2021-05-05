package Interface

import (
	"controller/Model"

	"github.com/gofiber/fiber"
)

type IUserData interface {
	GetHistory(c *fiber.Ctx) (resp Model.ResponseModel)
}
