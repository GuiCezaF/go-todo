package routes

import (
	"github.com/GuiCezaF/todo-api/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) *fiber.App {
	toDoRouter(app)

	return app
}
func toDoRouter(app *fiber.App) *fiber.App {

	app.Get("/health", handler.HandlerHealthCheck)

	return app
}
