package handlers

import (
	"strconv"

	"github.com/Anisulh/content_personalization/models"
	"github.com/Anisulh/content_personalization/utils"
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

func LikeContent(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	contentID, err := strconv.ParseUint(c.Params("contentId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid content id"})
	}

	interaction := utils.UserInteraction{
		UserID:    userID,
		ContentID: uint(contentID),
		Type:      "like",
	}

	if kafkaErr := utils.SendMessage([]string{"localhost:9092"}, "user-interactions", interaction); kafkaErr != nil {
		// Handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to send kafka message"})
	}

	return c.SendStatus(400)
	// Rest of your handler logic
	// ...
}

func DislikeContent(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	contentID, err := strconv.ParseUint(c.Params("contentId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid content id"})
	}

	interaction := utils.UserInteraction{
		UserID:    userID,
		ContentID: uint(contentID),
		Type:      "dislike",
	}

	if kafkaErr := utils.SendMessage([]string{"localhost:9092"}, "user-interactions", interaction); kafkaErr != nil {
		// Handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to send kafka message"})
	}

	return c.SendStatus(400)
	// Rest of your handler logic
	// ...
}

func BookmarkContent(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	contentID, err := strconv.ParseUint(c.Params("contentId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid content id"})
	}

	interaction := utils.UserInteraction{
		UserID:    userID,
		ContentID: uint(contentID),
		Type:      "bookmark",
	}

	if kafkaErr := utils.SendMessage([]string{"localhost:9092"}, "user-interactions", interaction); kafkaErr != nil {
		// Handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to send kafka message"})
	}

	return c.SendStatus(400)
	// Rest of your handler logic
	// ...
}
