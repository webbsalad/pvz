package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/webbsalad/pvz/internal/model"
)

func ExtractClaimsFromToken(token, secret string) (model.Role, error) {
	if secret == "" {
		return "", fmt.Errorf("JWT secret is not set")
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", model.ErrJwtExpired
		}
		return "", fmt.Errorf("invalid token: %w", err)
	}

	if !parsedToken.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("could not extract claims from token")
	}

	jwtRole, ok := claims["role"].(string)
	if !ok {
		return "", fmt.Errorf("role not found in token")
	}

	role, err := model.NewRole(jwtRole)
	if err != nil {
		return "", fmt.Errorf("convert jwt role to model: %w", err)
	}

	return role, nil
}
