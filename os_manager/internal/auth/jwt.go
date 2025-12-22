package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("super-secret-key") // depois mover para env
func GenerateToken(userID int64, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
