package handler

import (
	"github.com/GuiCezaF/todo-api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func HandlerHealthCheck(ctx *fiber.Ctx) error {

	utils.SendSuccess(ctx, "I'm Live", nil)
	return nil
}
