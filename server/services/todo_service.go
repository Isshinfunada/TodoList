package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/jackc/pgx"
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

	// 型アサーションは不要
	if todos == nil {
		log.Printf("todos: %#v", todos)
		log.Printf("No todos found for user")
		return nil, nil // エラーではなく空のリストを返す
	}

	return todos, nil
}

func (s *TodoService) CreateTodo(ctx context.Context, userID int32, text, status string) (models.Todo, error) {
	return s.Queries.CreateTodo(ctx, models.CreateTodoParams{
		UserID: pgtype.Int4{Int32: userID, Valid: true},
		Text:   text,
		Status: status,
	})
}

func (s *TodoService) EditTodo(ctx context.Context, id int32, text string) (models.Todo, error) {
	return s.Queries.EditTodo(ctx, models.EditTodoParams{
		ID:   id,
		Text: text,
	})
}
func (s *TodoService) DeleteTodo(ctx context.Context, id int32) error {
	err := s.Queries.DeleteTodo(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err) // エラーメッセージをラップ
	}

	// 削除された行数が0の場合は、TODOが存在しなかったと判断
	if err == pgx.ErrNoRows {
		return errors.New("todo not found")
	}

	return nil
}

func (s *TodoService) UpdateTodoStatus(ctx context.Context, id int32, status string) (models.Todo, error) {
	return s.Queries.UpdateTodoStatus(ctx, models.UpdateTodoStatusParams{
		ID:     id,
		Status: status,
	})
}
