package main

import (
	"github.com/GuiCezaF/todo-api/internal/handler/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.AddRoutes(app)

	app.Listen(":3000")
}
