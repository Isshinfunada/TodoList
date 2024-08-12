package routes

import (
	"github.com/Isshinfunada/TodoList/server/handlers"
	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/Isshinfunada/TodoList/server/services"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *models.Queries) {
	e.POST("/users", handlers.RegisterUser(db))
	e.POST("/login", handlers.Login(db))

	// TodoHandler のインスタンスを作成
	todoHandler := handlers.TodoHandler{
		TodoService: &services.TodoService{
			Queries: db,
		},
	}

	// TodoHandler の GetTodos メソッドを呼び出す
	e.GET("/todos/:user_id", todoHandler.GetTodos)
	e.POST("/todos", todoHandler.CreateTodo)
}
