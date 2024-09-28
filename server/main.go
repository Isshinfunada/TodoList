package main

import (
	"context"
	"log"
	"os"

	"database/sql"

	firebase "firebase.google.com/go"
	"github.com/Isshinfunada/TodoList/server/config"
	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/Isshinfunada/TodoList/server/routes"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

// initDBは、データベース接続を初期化します。
func initDB(cfg *config.Config) (*sql.DB, error) { // 戻り値を *sql.DB に変更
	dsn := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=disable"
	db, err := sql.Open("pgx", dsn) // "pgx" ドライバーを使用
	if err != nil {
		return nil, err
	}

	// 接続確認
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
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

func startServer(db *sql.DB, firebaseApp *firebase.App) { // パラメータを *sql.DB に変更
	e := echo.New()

	// models.New に *sql.DB を渡す
	queries := models.New(db)

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
	db, err := initDB(cfg) // 変数名を dbpool から db に変更
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer db.Close()

	// Firebase Admin SDKを初期化します。
	firebaseApp, err := initFirebase()
	if err != nil {
		log.Fatalf("Could not initialize Firebase: %v", err)
	}

	// サーバーを起動します。
	startServer(db, firebaseApp) // db を渡す
}
