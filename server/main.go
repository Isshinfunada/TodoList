package main

import (
	"context"
	"log"

	"github.com/Isshinfunada/TodoList/server/config"
	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/Isshinfunada/TodoList/server/routes"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func initDB(cfg *config.Config) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), "postgres://"+cfg.DBUser+":"+cfg.DBPassword+"@"+cfg.DBHost+":"+cfg.DBPort+"/"+cfg.DBName+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return dbpool, nil
}

func startServer(dbpool *pgxpool.Pool) {
	// log.Println("Server is starting...")
	e := echo.New()
	queries := models.New(dbpool)
	routes.InitRoutes(e, queries)
	e.Logger.Fatal(e.Start(":8080"))
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	dbpool, err := initDB(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer dbpool.Close()

	startServer(dbpool)
}
