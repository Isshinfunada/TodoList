package services

import (
	"context"
	"errors"

	"github.com/Isshinfunada/TodoList/server/models"
)

func AuthenticateUser(ctx context.Context, db *models.Queries, firebaseUID string) (*models.User, error) {
	row, err := db.GetUserByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, err
	}
	if row == (models.GetUserByFirebaseUIDRow{}) {
		return nil, errors.New("ユーザーが見つかりません")
	}

	user := models.User{
		ID:          row.ID,
		Username:    row.Username,
		Email:       row.Email,
		Password:    row.Password,
		FirebaseUid: row.FirebaseUid,
	}

	return &user, nil
}
