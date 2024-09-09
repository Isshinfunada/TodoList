package routes

import (
	"github.com/Isshinfunada/TodoList/server/handlers"
	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/Isshinfunada/TodoList/server/services"

	"firebase.google.com/go/auth"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo, db *models.Queries, authClient *auth.Client) {
	// CORS設定を追加
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// 認証が不要なルート
	e.POST("/users", handlers.RegisterUser(db))
	e.POST("/login", handlers.Login(authClient))

	// 認証が必要なルート
	authenticated := e.Group("/todos")
	authenticated.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("your_secret_key"),
	}))

	todoHandler := handlers.TodoHandler{
		TodoService: &services.TodoService{
			Queries: db,
		},
	}

	// 認証が必要なルートにJWTミドルウェアを適用
	authenticated.GET("/list", todoHandler.GetTodos)
	authenticated.POST("/create", todoHandler.CreateTodo)
	authenticated.POST("/edit", todoHandler.EditTodo)
	authenticated.POST("/delete", todoHandler.DeleteTodo)
	authenticated.POST("/:id/status", todoHandler.UpdateTodoStatus)
}
