package utils

import (
	"time"
	"tugas8/app/model"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key")

// GenerateToken membuat JWT berdasarkan data user
func GenerateToken(user model.User) (string, error) {
	claims := &model.JWTClaims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "tugas8",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}