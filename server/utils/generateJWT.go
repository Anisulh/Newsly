package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(id uint) (string, error) {
	jwtSecret := []byte(JWTSecret)
if len(jwtSecret) == 0 {
	return "", errors.New("JWT secret is not set")
}
// Create token
token := jwt.New(jwt.SigningMethodHS256)

// Set claims
claims := token.Claims.(jwt.MapClaims)
claims["user_id"] = id
claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

// Generate encoded token and send it as response.
t, err := token.SignedString(jwtSecret)
if err != nil {
	return "", err
}

return t, nil
}