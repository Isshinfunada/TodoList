package handlers

import (
	"net/http"

	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/Isshinfunada/TodoList/server/services"

	"github.com/labstack/echo/v4"
)

func RegisterUser(db *models.Queries) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(models.User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なリクエスト"})
		}

		if err := services.RegisterUser(c.Request().Context(), db, user); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, user)
	}
}
