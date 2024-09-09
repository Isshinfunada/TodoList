package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"github.com/Isshinfunada/TodoList/server/config"
	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/Isshinfunada/TodoList/server/routes"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

// initDBは、データベース接続を初期化します。
func initDB(cfg *config.Config) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), "postgres://"+cfg.DBUser+":"+cfg.DBPassword+"@"+cfg.DBHost+":"+cfg.DBPort+"/"+cfg.DBName+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return dbpool, nil
}

// initFirebaseは、Firebase Admin SDKを初期化します。
func initFirebase() (*firebase.App, error) {
	opt := option.WithCredentialsFile("firebase-adminsdk.json") // この行を適切なパスに変更してください
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}

// startServerは、Echoサーバーを起動します。
func startServer(dbpool *pgxpool.Pool, firebaseApp *firebase.App) {
	e := echo.New()
	queries := models.New(dbpool)

	// Firebase Auth クライアントを初期化
	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error getting Auth client: %v\n", err)
	}

	routes.InitRoutes(e, queries, authClient)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

func main() {
	// 設定を読み込みます。
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// データベース接続を初期化します。
	dbpool, err := initDB(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer dbpool.Close()

	// Firebase Admin SDKを初期化します。
	firebaseApp, err := initFirebase()
	if err != nil {
		log.Fatalf("Could not initialize Firebase: %v", err)
	}

	// サーバーを起動します。
	startServer(dbpool, firebaseApp)
}
