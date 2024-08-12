package routes

import (
	"github.com/Isshinfunada/TodoList/server/handlers"
	"github.com/Isshinfunada/TodoList/server/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo, db *models.Queries) {
	// CORS設定を追加
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},       // フロントエンドのオリジン
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT}, // 許可するHTTPメソッド
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.POST("/users", handlers.RegisterUser(db))
	e.POST("/login", handlers.Login(db))
}
