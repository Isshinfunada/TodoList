package services

import (
	"context"
	"errors"

	"github.com/Isshinfunada/TodoList/server/models"
)

func RegisterUser(ctx context.Context, db *models.Queries, user *models.User) error {
	// ユーザーが既に存在するか確認
	existingUser, err := db.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if existingUser != (models.User{}) {
		return errors.New("ユーザーは既に存在します")
	}

	// ユーザーを登録
	params := models.CreateUserParams{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	return db.CreateUser(ctx, params)
}
