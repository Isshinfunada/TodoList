package handlers

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
)

// AuthClientInterface は、認証に必要なメソッドを定義するインターフェースです
type AuthClientInterface interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
}

/**
 * @api {post} /login Authenticate a user with Firebase token
 * @apiName Login
 * @apiGroup User
 * @apiVersion 2.0.0
 * @apiParam {String} token Firebase authentication token
 * @apiSuccess (200) {Object} user User information
 * @apiError (400) {Object} error Invalid request
 * @apiError (401) {Object} error Unauthorized
 * @apiExample {curl} Example usage:
 *     curl -X POST 'http://localhost:8080/login' -H 'Authorization: Bearer <firebase_token>'
 */

func Login(authClient AuthClientInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the token from the Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		idToken := strings.TrimPrefix(authHeader, "Bearer ")

		if idToken == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "トークンが提供されていません"})
		}

		// Verify the Firebase token
		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "無効なトークン"})
		}

		// Token is valid, return user information
		return c.JSON(http.StatusOK, map[string]interface{}{
			"user": map[string]interface{}{
				"uid":   token.UID,
				"email": token.Claims["email"],
				// Add any other user information you want to return
			},
		})
	}
}
