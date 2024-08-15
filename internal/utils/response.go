package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SendSuccess(ctx *fiber.Ctx, message string, data interface{}) {
	ctx.Response().Header.SetContentType("application/json")
	ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}

func SendError(ctx *fiber.Ctx, statusCode int, message string) {
	ctx.Response().Header.SetContentType("application/json")
	ctx.Status(statusCode).JSON(fiber.Map{
		"message":   message,
		"errorCode": statusCode,
	})
}
