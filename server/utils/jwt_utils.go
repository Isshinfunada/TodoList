package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// GetUserIDFromToken extracts the user ID from the JWT token present in the request context.
func GetUserIDFromToken(c echo.Context) (int, error) {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok || user == nil {
		return 0, errors.New("no token found or invalid token type")
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user ID not found in token")
	}

	return int(userID), nil
}
