package handlers

import (
	"context"
	"net/http"

	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/Isshinfunada/TodoList/server/services"

	"github.com/labstack/echo/v4"
)

func Login(db *models.Queries) echo.HandlerFunc {
	return func(c echo.Context) error {
		var credentials struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.Bind(&credentials); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なリクエスト"})
		}

		token, err := services.AuthenticateUser(context.Background(), db, credentials.Email, credentials.Password)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]string{"token": token})
	}
}
