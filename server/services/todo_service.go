package services

import (
	"context"
	"fmt"
	"log"

	"github.com/Isshinfunada/TodoList/server/models"
)

// TodoServiceInterface は TodoService のインターフェースです。
type TodoServiceInterface interface {
	GetTodos(ctx context.Context, firebaseUID string) ([]models.Todo, error)
	GetTodoByID(ctx context.Context, id int32) (models.Todo, error)
	CreateTodo(ctx context.Context, firebaseUID string, text, status string) (models.Todo, error)
	EditTodo(ctx context.Context, id int32, text string) (models.Todo, error)
	DeleteTodo(ctx context.Context, id int32) error
	UpdateTodoStatus(ctx context.Context, id int32, status string) (models.Todo, error)
}

// TodoService はTodo関連のビジネスロジックを提供します。
type TodoService struct {
	Queries *models.Queries
}

// NewTodoService は新しい TodoService を初期化します。
func NewTodoService(q *models.Queries) *TodoService {
	return &TodoService{Queries: q}
}

// GetTodos はfirebaseUIDに基づいてTodosを取得します。
func (s *TodoService) GetTodos(ctx context.Context, firebaseUID string) ([]models.Todo, error) {
	todos, err := s.Queries.ListTodos(ctx, firebaseUID)
	if err != nil {
		log.Printf("Error in ListTodos: %v", err)
		return nil, err
	}

	// todos が nil の場合、空のスライスを返す
	if todos == nil {
		log.Println("No todos found for user")
		return []models.Todo{}, nil
	}

	return todos, nil
}

// GetTodoByID は特定のTodoを取得します。
func (s *TodoService) GetTodoByID(ctx context.Context, id int32) (models.Todo, error) {
	todo, err := s.Queries.GetTodoByID(ctx, id)
	if err != nil {
		log.Printf("Error in GetTodoByID: %v", err)
		return models.Todo{}, err
	}
	return todo, nil
}

// CreateTodo は新しいTodoを作成します。
func (s *TodoService) CreateTodo(ctx context.Context, firebaseUID string, text, status string) (models.Todo, error) {
	todo, err := s.Queries.CreateTodo(ctx, models.CreateTodoParams{
		UserID: firebaseUID,
		Text:   text,
		Status: status,
	})
	if err != nil {
		log.Printf("Error in CreateTodo: %v", err)
		return models.Todo{}, err
	}
	return todo, nil
}

// EditTodo は既存のTodoを編集します。
func (s *TodoService) EditTodo(ctx context.Context, id int32, text string) (models.Todo, error) {
	todo, err := s.Queries.EditTodo(ctx, models.EditTodoParams{
		ID:   id,
		Text: text,
	})
	if err != nil {
		log.Printf("Error in EditTodo: %v", err)
		return models.Todo{}, err
	}
	return todo, nil
}

// DeleteTodo はTodoを削除します。
func (s *TodoService) DeleteTodo(ctx context.Context, id int32) error {
	err := s.Queries.DeleteTodo(ctx, id)
	if err != nil {
		log.Printf("Error deleting todo ID %d: %v", id, err)
		return fmt.Errorf("failed to delete todo: %w", err)
	}
	return nil
}

// UpdateTodoStatus はTodoのステータスを更新します。
func (s *TodoService) UpdateTodoStatus(ctx context.Context, id int32, status string) (models.Todo, error) {
	todo, err := s.Queries.UpdateTodoStatus(ctx, models.UpdateTodoStatusParams{
		ID:     id,
		Status: status,
	})
	if err != nil {
		log.Printf("Error in UpdateTodoStatus: %v", err)
		return models.Todo{}, err
	}
	return todo, nil
}
