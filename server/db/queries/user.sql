-- name: CreateUser :exec
INSERT INTO users (username, email, password) VALUES ($1,$2,$3);

-- name: GetUserByEmail :one
SELECT id, username, email, password FROM users WHERE email = $1;