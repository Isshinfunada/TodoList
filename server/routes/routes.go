package routes

import (
	"github.com/Isshinfunada/TodoList/server/handlers"
	"github.com/Isshinfunada/TodoList/server/models"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *models.Queries) {
	e.POST("/users", handlers.RegisterUser(db))
	e.POST("/login", handlers.Login(db))
}
