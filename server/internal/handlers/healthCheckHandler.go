package handlers

import (
	"github.com/gofiber/fiber/v2"
)


func (h *Handler) HealthCheck(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
