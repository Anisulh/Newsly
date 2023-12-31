package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func JWTProtected() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Retrieve the JWT token from the cookie
		cookie := c.Cookies("jwt")

		// If the cookie is empty, return an unauthorized error
		if cookie == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: No token provided")
		}

		// Parse the token
		token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
			// Ensure the token's algorithm matches the expected algorithm
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			// Return the secret key
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		// Check if the token is valid
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: Invalid token")
		}

		// Extracting User ID from Token Claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["user_id"].(string) // Assuming user_id is stored as a string
			c.Locals("userID", userID)           // Setting the user ID in Fiber's context
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: Invalid token")
		}

		return c.Next()
	}
}
