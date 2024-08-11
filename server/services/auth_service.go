package services

import (
	"context"
	"errors"
	"time"

	"github.com/Isshinfunada/TodoList/server/models"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func AuthenticateUser(ctx context.Context, db *models.Queries, email, password string) (string, error) {
	user, err := db.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if user == (models.User{}) {
		return "", errors.New("ユーザーが見つかりません")
	}

	if user.Password != password {
		return "", errors.New("パスワードが一致しません")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: int(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
