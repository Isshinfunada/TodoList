-- name: ListTodos :many
SELECT id, user_id, text, status, created_at, updated_at FROM todos WHERE user_id = $1;

-- name: CreateTodo :one
INSERT INTO todos (user_id, text, status, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING id, user_id, text, status, created_at, updated_at;

-- name: EditTodo :one
UPDATE todos SET text = $2, updated_at = NOW() WHERE id = $1 RETURNING id, user_id, text, status, created_at, updated_at;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1;

-- name: UpdateTodoStatus :one
UPDATE todos SET status = $2, updated_at = NOW() WHERE id = $1 RETURNING id, user_id, text, status, created_at, updated_at;
