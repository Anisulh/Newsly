package handlers

import (
	"Newsly/internal/models"
	"Newsly/internal/utils"
	"Newsly/web/templates/partials"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) UserRegistration(c *fiber.Ctx) error {
	userInfo := &models.User{}
	if err := c.BodyParser(userInfo); err != nil {
		c.Status(fiber.StatusBadRequest)
		return Render(c, partials.RegisterError("Invalid user information"))
	}

	// Check if the user already exists
	var count int64
	h.DB.Model(&models.User{}).Where("email = ?", userInfo.Email).Count(&count)

	if count > 0 {
		c.Status(fiber.StatusConflict)
		return Render(c, partials.RegisterError("Invalid user information"))
	}
	res := h.DB.Create(&userInfo)
	if res.Error != nil {
		c.Set("HX-Redirect", "/internal-server-error")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	token, err := utils.GenerateJWT(userInfo.ID, h.JWTSecret)
	if err != nil {
		c.Set("HX-Redirect", "/internal-server-error")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Expires:  time.Now().Add(72 * time.Hour),
		Secure:   h.Environment == "production",
	}

	c.Cookie(&cookie)
	c.Set("HX-Redirect", "/auth/interest-topics")

	// Sending back a success message
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": "User registered successfully"})
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) UserLogin(c *fiber.Ctx) error {
	// Structure to hold login credentials
	loginInfo := UserLoginRequest{}
	if err := c.BodyParser(&loginInfo); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return Render(c, partials.LoginError())
	}

	// Find user by email
	var user models.User
	if err := h.DB.Where("email = ?", loginInfo.Email).First(&user).Error; err != nil {
		c.Status(fiber.StatusUnauthorized)
		return Render(c, partials.LoginError())
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password)); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return Render(c, partials.LoginError())
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, h.JWTSecret)
	if err != nil {
		c.Set("HX-Redirect", "/internal-server-error")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Set cookie
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Expires:  time.Now().Add(72 * time.Hour),
		Secure:   h.Environment == "production",
		SameSite: "Strict",
	}

	c.Cookie(&cookie)
	// if user's category interests are not set, redirect to interests page
	if len(user.CategoryInterests) == 0 {
		c.Set("HX-Redirect", "/auth/interest-topics")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "User logged in successfully"})
	}
	c.Set("HX-Redirect", "/auth/feed")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "User logged in successfully"})
}

func (h *Handler) UserLogout(c *fiber.Ctx) error {
	c.ClearCookie("token")
	c.Set("HX-Redirect", "/")
	return c.SendStatus(fiber.StatusOK)
}

// Secured Routes
// SaveUserInterestsRequest is the expected JSON payload.
type SaveUserInterestsRequest struct {
	Categories []string `json:"categories"`
}

// SaveUserInterests updates the user's interests based on the selected category keys.
func (h *Handler) SaveUserInterests(c *fiber.Ctx) error {
	// Parse the incoming JSON payload.
	var req SaveUserInterestsRequest
	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Get the current user ID from the context.
	userID, ok := c.Locals("account").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized or invalid user session",
		})
	}

	// Fetch the user from the database.
	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		log.Printf("User not found: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Fetch the categories from the database based on the provided keys.
	var selectedCategories []models.Category
	if err := h.DB.
		Where("key IN ?", req.Categories).
		Find(&selectedCategories).Error; err != nil {
		log.Printf("Error fetching categories: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching categories",
		})
	}

	// Replace the user's current category interests with the selected ones.
	if err := h.DB.Model(&user).Association("CategoryInterests").Replace(selectedCategories); err != nil {
		log.Printf("Error updating interests: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error updating user interests",
		})
	}
	c.Set("HX-Redirect", "/auth/feed")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User interests updated successfully",
	})
}
