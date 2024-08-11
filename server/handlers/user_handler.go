package handlers

import (
	"net/http"

	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/Isshinfunada/TodoList/server/services"

	"github.com/labstack/echo/v4"
)

/**
 * @api {post} /users Register a new user
 * @apiName RegisterUser
 * @apiGroup User
 * @apiVersion 1.0.0
 * @apiParam {String} Username User's username
 * @apiParam {String} Email User's email
 * @apiParam {String} Password User's password
 * @apiSuccess (201) {Object} user The registered user
 * @apiError (400) {Object} error Invalid request
 * @apiError (500) {Object} error Internal server error
 * @apiExample {curl} Example usage:
 *     curl -X POST 'http://localhost:8080/users' -H 'Content-Type: application/json' -d '{"Username": "testuser1", "Email": "testuser1@example.com", "Password": "password1"}'
 */
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
