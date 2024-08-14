package routes

import (
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) *fiber.App {
	toDoRouter(app)

	return app
}
func toDoRouter(app *fiber.App) *fiber.App {

	type response struct {
		Message string `json:"message"`
	}

	app.Get("/", func(c *fiber.Ctx) error {
		r := response{Message: "Hello World"}
		return c.JSON(r)
	})

	return app
}
