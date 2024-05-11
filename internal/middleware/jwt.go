package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// JWTProtected creates a middleware to protect routes with JWT validation.
func JWTProtected(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the JWT token from the cookie
		tokenString := c.Cookies("jwt")

		// If the cookie is empty, return an unauthorized error
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
				"message": "No token provided",
			})
		}

		// Parse the token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Ensure the token's algorithm matches the expected HMAC algorithm
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(secret), nil
		})

		// Check if the token is valid
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
				"message": "Invalid token",
				"details": err.Error(),
			})
		}

		// Ensure the token was successfully validated
		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: Invalid token")
		}

		// Extracting User ID from Token Claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, ok := claims["user_id"].(string) // Ensure type assertion is safe
			if !ok {
				return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: User ID invalid")
			}
			c.Locals("userID", userID) // Setting the user ID in Fiber's context
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: Invalid token")
		}

		return c.Next()
	}
}