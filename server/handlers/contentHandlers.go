package handlers

import (
	"github.com/Anisulh/content_personalization/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetContent(c *fiber.Ctx) error {
	var contents []models.Content
	if err := h.DB.Find(&contents).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot fetch content"})
	}

	return c.JSON(contents)
}


func (h *Handler) GetContentCategories(c *fiber.Ctx) error {
	var categories []models.Category
	if err := h.DB.Find(&categories).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot fetch categories"})
	}

	return c.JSON(categories)
}
