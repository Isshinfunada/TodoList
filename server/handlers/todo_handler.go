package handlers

import (
	"log"
	"net/http"
	"strconv"

	"firebase.google.com/go/auth"
	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/Isshinfunada/TodoList/server/services"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	TodoService *services.TodoService
}

/**
 * @api {get} /todos/list Get Todos for the authenticated user
 * @apiName GetTodos
 * @apiGroup Todo
 * @apiVersion 1.0.0
 * @apiSuccess (200) {Array} todos Array of Todo objects
 * @apiSuccess (200) {Number} todos.ID Todo's unique ID
 * @apiSuccess (200) {Number} todos.UserID ID of the user who owns the todo
 * @apiSuccess (200) {String} todos.Text Text content of the todo
 * @apiSuccess (200) {String} todos.Status Status of the todo (e.g., "pending", "completed")
 * @apiSuccess (200) {String} todos.CreatedAt Timestamp when the todo was created
 * @apiSuccess (200) {String} todos.UpdatedAt Timestamp when the todo was last updated
 * @apiError (401) {Object} error Invalid token
 * @apiError (500) {Object} error Failed to fetch todos
 * @apiExample {curl} Example usage:
 *     curl -X GET 'http://localhost:8080/todos/list' -H 'Authorization: Bearer <firebase_token>'
 */
func (h *TodoHandler) GetTodos(c echo.Context) error {
	token := c.Get("user").(*auth.Token)
	userID, err := strconv.Atoi(token.UID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
	}

	todos, err := h.TodoService.GetTodos(c.Request().Context(), int32(userID))
	if err != nil {
		log.Printf("Error in GetTodos: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch todos"})
	}

	if todos == nil {
		todos = []models.Todo{} // 空のスライスを返す
	}

	return c.JSON(http.StatusOK, todos)
}

/**
 * @api {post} /todos Create a new Todo
 * @apiName CreateTodo
 * @apiGroup Todo
 * @apiVersion 1.0.0
 * @apiBody {Number} user_id ID of the user who owns the todo
 * @apiBody {String} text Text content of the todo
 * @apiBody {String} status Status of the todo (e.g., "pending", "completed")
 * @apiSuccess (201) {Object} todo Created Todo object
 * @apiSuccess (201) {Number} todo.ID Todo's unique ID
 * @apiSuccess (201) {Number} todo.UserID ID of the user who owns the todo
 * @apiSuccess (201) {String} todo.Text Text content of the todo
 * @apiSuccess (201) {String} todo.Status Status of the todo
 * @apiSuccess (201) {String} todo.CreatedAt Timestamp when the todo was created
 * @apiSuccess (201) {String} todo.UpdatedAt Timestamp when the todo was last updated
 * @apiError (400) {Object} error Invalid input
 * @apiError (500) {Object} error Failed to create todo
 * @apiExample {curl} Example usage:
 *     curl -X POST 'http://localhost:8080/todos/create' \
 *     -H 'Content-Type: application/json' \
 *     -d '{
 *       "user_id": 1,
 *       "text": "新しいTODO",
 *       "status": "pending"
 *     }'
 */
func (h *TodoHandler) CreateTodo(c echo.Context) error {
	// JWTトークンからユーザーIDを取得
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		log.Printf("Invalid token: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
	}

	var req struct {
		Text   string `json:"text"`
		Status string `json:"status"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	todo, err := h.TodoService.CreateTodo(c.Request().Context(), int32(userID), req.Text, req.Status)
	if err != nil {
		log.Printf("Error in CreateTodo: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create todo"})
	}

	return c.JSON(http.StatusCreated, todo)
}

/**
 * @api {post} /todos/edit Edit an existing Todo
 * @apiName EditTodo
 * @apiGroup Todo
 * @apiVersion 1.0.0
 * @apiBody {Number} id ID of the todo to edit
 * @apiBody {String} text New text content of the todo
 * @apiSuccess (200) {Object} todo Edited Todo object
 * @apiSuccess (200) {Number} todo.ID Todo's unique ID
 * @apiSuccess (200) {Number} todo.UserID ID of the user who owns the todo
 * @apiSuccess (200) {String} todo.Text Text content of the todo
 * @apiSuccess (200) {String} todo.Status Status of the todo
 * @apiSuccess (200) {String} todo.CreatedAt Timestamp when the todo was created
 * @apiSuccess (200) {String} todo.UpdatedAt Timestamp when the todo was last updated
 * @apiError (400) {Object} error Invalid input
 * @apiError (500) {Object} error Failed to edit todo
 * @apiExample {curl} Example usage:
 *     curl -X POST 'http://localhost:8080/todos/edit' \
 *     -H 'Content-Type: application/json' \
 *     -d '{
 *       "id": 2,
 *       "text": "更新されたTODO"
 *     }'
 */
func (h *TodoHandler) EditTodo(c echo.Context) error {
	var req struct {
		ID   int32  `json:"id"`
		Text string `json:"text"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	todo, err := h.TodoService.EditTodo(c.Request().Context(), req.ID, req.Text)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to edit todo"})
	}

	return c.JSON(http.StatusOK, todo)
}

/**
 * @api {post} /todos/delete Delete a Todo
 * @apiName DeleteTodo
 * @apiGroup Todo
 * @apiVersion 1.0.0
 * @apiBody {Number} id ID of the todo to delete
 * @apiSuccess (200) {Object} message Success message
 * @apiSuccess (200) {String} message.message "Todo deleted successfully"
 * @apiError (400) {Object} error Invalid input
 * @apiError (404) {Object} error Todo not found
 * @apiError (500) {Object} error Failed to delete todo
 * @apiExample {curl} Example usage:
 *     curl -X POST 'http://localhost:8080/todos/delete' \
 *     -H 'Content-Type: application/json' \
 *     -d '{
 *       "id": 2
 *     }'
 */
func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	var req struct {
		ID int32 `json:"id"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	err := h.TodoService.DeleteTodo(c.Request().Context(), req.ID)
	if err != nil {
		if err.Error() == "todo not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
		}
		log.Printf("Error in DeleteTodo: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete todo"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Todo deleted successfully"})
}

/**
 * @api {post} /todos/:id/status Update Todo Status
 * @apiName UpdateTodoStatus
 * @apiGroup Todo
 * @apiVersion 1.0.0
 * @apiParam {Number} id ID of the todo to update
 * @apiBody {String} status New status of the todo (e.g., "pending", "completed")
 * @apiSuccess (200) {Object} todo Updated Todo object
 * @apiSuccess (200) {Number} todo.ID Todo's unique ID
 * @apiSuccess (200) {Number} todo.UserID ID of the user who owns the todo
 * @apiSuccess (200) {String} todo.Text Text content of the todo
 * @apiSuccess (200) {String} todo.Status Status of the todo
 * @apiSuccess (200) {String} todo.CreatedAt Timestamp when the todo was created
 * @apiSuccess (200) {String} todo.UpdatedAt Timestamp when the todo was last updated
 * @apiError (400) {Object} error Invalid todo ID or input
 * @apiError (500) {Object} error Failed to update todo status
 * @apiExample {curl} Example usage:
 *     curl -X POST 'http://localhost:8080/todos/2/status' \
 *     -H 'Content-Type: application/json' \
 *     -d '{
 *       "status": "completed"
 *     }'
 */
func (h *TodoHandler) UpdateTodoStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid todo ID"})
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	todo, err := h.TodoService.UpdateTodoStatus(c.Request().Context(), int32(id), req.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update todo status"})
	}

	return c.JSON(http.StatusOK, todo)
}
