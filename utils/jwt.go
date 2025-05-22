package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT generates a JWT token using user ID and email
func GenerateJWT(userID int, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // token valid 1 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ambil secret dari environment variable
	secret := os.Getenv("JWT_SECRET")

	// Generate token string
	return token.SignedString([]byte(secret))
}
