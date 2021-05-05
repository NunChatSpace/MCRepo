package Delivery

import (
	"controller/Interface"

	"github.com/gofiber/fiber"
)

type RequestHandler struct {
	UD Interface.IUserData
}

var handler RequestHandler

func SetupRoute(app *fiber.App, ud Interface.IUserData) {
	handler = RequestHandler{
		UD: ud,
	}
	app.Get("/", func(c *fiber.Ctx) {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	app.Get("/userinfo", func(c *fiber.Ctx) {
		resp := handler.UD.GetHistory(c)
		c.JSON(resp)
	})
}
