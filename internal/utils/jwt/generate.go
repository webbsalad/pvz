package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/webbsalad/pvz/internal/model"
)

func GenerateTokens(role model.Role, secret string) (string, error) {
	now := time.Now()

	claims := jwt.MapClaims{
		"role": role.String(),
		"exp":  now.Add(24 * time.Hour).Unix(),
	}
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenJWT.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}

	return token, nil
}
