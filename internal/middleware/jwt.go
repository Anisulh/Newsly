package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// JWTProtected creates a middleware to protect routes with JWT validation.
func (m *Middleware) JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the JWT token from the cookie
		tokenString := c.Cookies("token")

		// If the cookie is empty, return an unauthorized error
		if tokenString == "" {
			c.Set("HX-Redirect", "/")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthorized",
				"message": "No token provided",
			})
		}

		// Parse the token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(m.JWTSecret), nil
		})
		// Check if the token is valid
		if err != nil {
			c.Set("HX-Redirect", "/")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthorized",
				"message": "Invalid token",
				"details": err.Error(),
			})
		}

		// Ensure the token was successfully validated
		if !token.Valid {
			c.Set("HX-Redirect", "/")
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: Invalid token")
		}

		// Extracting User ID from Token Claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			userIDFloat, ok := claims["sub"].(float64) // Ensure type assertion is safe
			if !ok {
				c.Set("HX-Redirect", "/")
				return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: User ID invalid")
			}

			userID := uint(userIDFloat) // Convert float64 to uint
			c.Locals("account", userID) // Setting the user ID in Fiber's context
		} else {
			c.Set("HX-Redirect", "/")
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: Invalid token")
		}

		return c.Next()
	}
}
