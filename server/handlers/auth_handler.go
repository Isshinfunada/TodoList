package handlers

import (
	"context"
	"net/http"

	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/Isshinfunada/TodoList/server/services"

	"github.com/labstack/echo/v4"
)

/**
 * @api {post} /login Authenticate a user
 * @apiName Login
 * @apiGroup User
 * @apiVersion 1.0.0
 * @apiParam {String} email User's email
 * @apiParam {String} password User's password
 * @apiSuccess (200) {String} token Authentication token
 * @apiError (400) {Object} error Invalid request
 * @apiError (401) {Object} error Unauthorized
 * @apiExample {curl} Example usage:
 *     curl -X POST 'http://localhost:8080/login' -H 'Content-Type: application/json' -d '{"email": "testuser1@example.com", "password": "password1"}'
 */
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
