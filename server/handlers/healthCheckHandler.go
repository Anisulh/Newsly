package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) HealthCheck(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
