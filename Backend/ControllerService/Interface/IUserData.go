package Interface

import (
	"controller/Model"

	"github.com/gofiber/fiber/v2"
)

type IUserData interface {
	GetHistory(c *fiber.Ctx) (resp Model.ResponseModel)
}
