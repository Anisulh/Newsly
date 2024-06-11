package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id uint, secret string) (string, error) {
	jwtSecret := []byte(secret)
	if len(jwtSecret) == 0 {
		return "", errors.New("JWT secret is not set")
	}
	// Create a new JWT token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,                               // Subject (user identifier)
		"iss": "Newsly",                         // Issuer
		"aud": "user",                           // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
	})

	tokenString, err := claims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
