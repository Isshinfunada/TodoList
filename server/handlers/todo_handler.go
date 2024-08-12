package handlers

import (
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/Isshinfunada/TodoList/server/services"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	TodoService *services.TodoService
}

/**
 * @api {get} /todos/:user_id Get Todos for a User
 * @apiName GetTodos
 * @apiGroup Todo
 * @apiVersion 1.0.0
 * @apiParam {Number} user_id User's unique ID
 * @apiSuccess (200) {Array} todos Array of Todo objects
 * @apiSuccess (200) {Number} todos.ID Todo's unique ID
 * @apiSuccess (200) {Number} todos.UserID ID of the user who owns the todo
 * @apiSuccess (200) {String} todos.Text Text content of the todo
 * @apiSuccess (200) {String} todos.Status Status of the todo (e.g., "pending", "completed")
 * @apiSuccess (200) {String} todos.CreatedAt Timestamp when the todo was created
 * @apiSuccess (200) {String} todos.UpdatedAt Timestamp when the todo was last updated
 * @apiError (400) {Object} error Invalid user ID
 * @apiError (500) {Object} error Failed to fetch todos
 * @apiExample {curl} Example usage:
 *     curl -X GET 'http://localhost:8080/todos/1'
 */
func (h *TodoHandler) GetTodos(c echo.Context) error {
	userIDStr := c.Param("user_id")
	log.Printf("userIDStr: %s", userIDStr)
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Printf("Invalid user ID: %s, error: %v", userIDStr, err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// int32 へのキャストエラーをチェック
	if int64(userID) > math.MaxInt32 || int64(userID) < math.MinInt32 {
		log.Printf("User ID out of range: %d", userID)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID out of range"})
	}

	todos, err := h.TodoService.GetTodos(c.Request().Context(), int32(userID))
	if err != nil {
		// エラーログを追加
		log.Printf("Error in GetTodos: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch todos"})
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
 *     curl -X POST 'http://localhost:8080/todos' \
 *     -H 'Content-Type: application/json' \
 *     -d '{
 *       "user_id": 1,
 *       "text": "新しいTODO",
 *       "status": "pending"
 *     }'
 */
func (h *TodoHandler) CreateTodo(c echo.Context) error {
	var req struct {
		UserID int32  `json:"user_id"`
		Text   string `json:"text"`
		Status string `json:"status"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	todo, err := h.TodoService.CreateTodo(c.Request().Context(), req.UserID, req.Text, req.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create todo"})
	}

	return c.JSON(http.StatusCreated, todo)
}
