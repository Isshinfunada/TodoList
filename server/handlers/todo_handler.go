package handlers

import (
	"net/http"
	"strconv"

	"github.com/Isshinfunada/TodoList/server/services"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	TodoService *services.TodoService
}

func (h *TodoHandler) GetTodos(c echo.Context) error {
	userID, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	todos, err := h.TodoService.GetTodos(c.Request().Context(), int32(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch todos"})
	}

	return c.JSON(http.StatusOK, todos)
}

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
