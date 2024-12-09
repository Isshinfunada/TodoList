package services

import (
	"context"
	"fmt"
	"reflect"

	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/jackc/pgconn"
)

type UserQueries interface {
	CreateUser(ctx context.Context, arg models.CreateUserParams) error
}

func RegisterUser(ctx context.Context, db UserQueries, user *models.User) error {
	// ユーザーを登録
	params := models.CreateUserParams{
		Username:    user.Username,
		Email:       user.Email,
		Password:    user.Password,
		FirebaseUid: user.FirebaseUid,
	}
	err := db.CreateUser(ctx, params)
	if err != nil {
		// エラーを直接型アサーションでキャスト
		if pgErr, ok := err.(*pgconn.PgError); ok {
			// ユニーク制約違反の場合
			if pgErr.Code == "23505" && pgErr.ConstraintName == "users_email_key" {
				return fmt.Errorf("ユーザーは既に存在します: %v", user.Email)
			}
		}

		return fmt.Errorf("ユーザーの登録中にエラーが発生しました: %w, エラーの詳細: %v, エラーの型: %v", err, err, reflect.TypeOf(err))
	}
	return nil
}
