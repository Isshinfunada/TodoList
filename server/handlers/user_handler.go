package handlers

import (
	"net/http"
	"strings"

	"context"

	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/Isshinfunada/TodoList/server/services"
	"github.com/labstack/echo/v4"
)

/**
 * @api {post} /users Register a new user
 * @apiName RegisterUser
 * @apiGroup User
 * @apiVersion 2.0.0
 * @apiParam {String} Username User's username
 * @apiParam {String} Email User's email
 * @apiParam {String} Password User's password
 * @apiSuccess (201) {Object} user The registered user
 * @apiError (400) {Object} error Invalid request
 * @apiError (401) {Object} error Unauthorized
 * @apiError (500) {Object} error Internal server error
 * @apiExample {curl} Example usage:
 *     curl -X POST 'http://localhost:8080/users' -H 'Content-Type: application/json' -d '{"Username": "testuser1", "Email": "testuser1@example.com", "Password": "password1"}'
 */
func RegisterUser(db services.UserQueries, authClient AuthClientInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Authorizationヘッダーからトークンを取得
		authHeader := c.Request().Header.Get("Authorization")
		idToken := strings.TrimPrefix(authHeader, "Bearer ")

		if idToken == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "トークンが提供されていません"})
		}

		// Firebaseトークンを検証
		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "無効なトークン"})
		}

		// トークンが有効な場合、ユーザー情報を取得
		user := new(models.User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なリクエスト"})
		}

		// トークンのUIDをユーザー情報に設定
		user.FirebaseUid = token.UID

		// ユーザーを登録
		if err := services.RegisterUser(c.Request().Context(), db, user); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, user)
	}
}
