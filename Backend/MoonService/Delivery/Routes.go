package Delivery

import (
	"moonservice/Interface"

	"github.com/gofiber/fiber"
)

type RequestHandler struct {
	MS Interface.IMoonService
}

var handler RequestHandler

func SetupRoute(app *fiber.App, ms Interface.IMoonService) {

	handler = RequestHandler{
		MS: ms,
	}

	app.Get("/", func(c *fiber.Ctx) {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	app.Post("/buy", func(c *fiber.Ctx) {
		resp := handler.MS.BuyMOON(c)
		c.JSON(resp)
	})
}
