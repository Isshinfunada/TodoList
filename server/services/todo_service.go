package services

import (
	"context"

	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/jackc/pgx/v5/pgtype"
)

type TodoService struct {
	Queries *models.Queries
}

func (s *TodoService) GetTodos(ctx context.Context, userID int32) ([]models.Todo, error) {
	return s.Queries.ListTodos(ctx, pgtype.Int4{Int32: userID})
}

func (s *TodoService) CreateTodo(ctx context.Context, userID int32, text, status string) (models.Todo, error) {
	return s.Queries.CreateTodo(ctx, models.CreateTodoParams{
		UserID: pgtype.Int4{Int32: userID}, // Changed this line
		Text:   text,
		Status: status,
	})
}
