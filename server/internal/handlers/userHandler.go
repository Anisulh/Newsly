package handlers

import (
	"time"

	"Newsly/internal/models"
	"Newsly/internal/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) UserRegistration(c *fiber.Ctx) error {
	userInfo := &models.User{}
	if err := c.BodyParser(userInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user information"})
	}

	// Check if the user already exists
	var count int64
	h.DB.Model(&models.User{}).Where("email = ?", userInfo.Email).Count(&count)
	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exists"})
	}
	res := h.DB.Create(&userInfo)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	token, err := utils.GenerateJWT(userInfo.ID, h.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		HTTPOnly: true,
		Expires:  time.Now().Add(72 * time.Hour),
	}

	c.Cookie(&cookie)

	// Sending back a success message
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": "User registered successfully"})

}

func (h *Handler) UserLogin(c *fiber.Ctx) error {
	// Structure to hold login credentials
	loginInfo := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := c.BodyParser(&loginInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid login information"})
	}

	// Find user by email
	var user models.User
	if err := h.DB.Where("email = ?", loginInfo.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Incorrect email or password"})
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Incorrect email or password"})
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, h.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Set cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		HTTPOnly: true,
		Expires:  time.Now().Add(72 * time.Hour),
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "User logged in successfully"})
}

// Secured Routes

func (h *Handler) GetUserProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

func (h *Handler) UpdateUserProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	var updateInfo models.User
	if err := c.BodyParser(&updateInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user information"})
	}

	if err := h.DB.Model(&models.User{}).Where("id = ?", userID).Updates(updateInfo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update user profile"})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) GetUserPreferences(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	var preferences []models.Preference
	if err := h.DB.Where("user_id = ?", userID).Find(&preferences).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch preferences"})
	}

	return c.JSON(preferences)
}

func (h *Handler) UpdateUserPreferences(c *fiber.Ctx) error {
	//userID := c.Locals("userID").(string)

	var newPreferences []models.Preference
	if err := c.BodyParser(&newPreferences); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid preference format"})
	}

	// Update preferences logic
	// ...

	return c.SendStatus(fiber.StatusOK)
}
