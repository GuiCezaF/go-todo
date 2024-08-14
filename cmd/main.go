package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/GuiCezaF/todo-api/internal/handler/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	app := fiber.New()
	routes.AddRoutes(app)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		app.Shutdown()
	}()
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
