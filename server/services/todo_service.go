package services

import (
	"context"
	"errors"
	"log"

	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/jackc/pgx/v5/pgtype"
)

type TodoService struct {
	Queries *models.Queries
}

func (s *TodoService) GetTodos(ctx context.Context, userID int32) ([]models.Todo, error) {
	todos, err := s.Queries.ListTodos(ctx, pgtype.Int4{Int32: userID, Valid: true})
	if err != nil {
		log.Printf("Error in ListTodos: %v", err)
		return nil, err
	}

	modelTodos := ([]models.Todo)(todos)
	if modelTodos == nil {
		log.Printf("todos: %#v", todos)
		log.Printf("Type assertion failed in GetTodos")
		return nil, errors.New("internal server error")
	}

	return modelTodos, nil
}

func (s *TodoService) CreateTodo(ctx context.Context, userID int32, text, status string) (models.Todo, error) {
	return s.Queries.CreateTodo(ctx, models.CreateTodoParams{
		UserID: pgtype.Int4{Int32: userID, Valid: true},
		Text:   text,
		Status: status,
	})
}
