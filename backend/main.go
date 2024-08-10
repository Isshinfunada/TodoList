package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var db *sql.DB

func main() {
	e := echo.New()

	var err error
	db, err = sql.Open("postgres", "postgres://user:password@db:5432/todos?sslmode=disable")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/api/todos", getTodos)

	e.Logger.Fatal(e.Start(":8080"))
}

func getTodos(c echo.Context) error {
	rows, err := db.Query("SELECT id, text FROM todos")
	if err != nil {
		return err
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Text); err != nil {
			return err
		}
		todos = append(todos, todo)
	}
	return c.JSON(http.StatusOK, todos)
}
