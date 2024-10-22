// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Todo struct {
	ID        int32
	UserID    pgtype.Int4
	Text      string
	Status    string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type User struct {
	ID       int32
	Username string
	Email    string
	Password string
}
