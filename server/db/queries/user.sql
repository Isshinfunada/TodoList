-- name: CreateUser :exec
INSERT INTO users (username, email, password, firebase_uid) VALUES ($1, $2, $3, $4);

-- name: GetUserByFirebaseUID :one
SELECT id, username, email, password, firebase_uid FROM users WHERE firebase_uid = $1;