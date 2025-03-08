package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/om13rajpal/gobackend/config"
	"github.com/om13rajpal/gobackend/internal/models"
)

func GenerateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	tokenData := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := tokenData.SignedString([]byte(config.JWT_SECRET))

	if err != nil {
		return "", err
	}

	return string(signedToken), nil
}
