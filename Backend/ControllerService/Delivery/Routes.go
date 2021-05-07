package Delivery

import (
	"controller/Interface"

	"github.com/gofiber/fiber/v2"
)

type RequestHandler struct {
	UD Interface.IUserData
}

var handler RequestHandler

func SetupRoute(app *fiber.App, ud Interface.IUserData) {
	handler = RequestHandler{
		UD: ud,
	}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	app.Get("/userinfo", func(c *fiber.Ctx) error {
		resp := handler.UD.GetHistory(c)
		return c.JSON(resp)
	})
}
