-- name: GetTask :one
SELECT * FROM tasks
WHERE id = $1 LIMIT 1;

-- name: ListTasks :many
SELECT * FROM tasks
ORDER BY active_at ASC;

-- name: CreateTask :one
INSERT INTO tasks (
    title, active_at
) VALUES (
    $1, $2
)
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;

-- name: UpdateTask :one
UPDATE tasks
SET 
    title = $2,
    active_at = $3,
    done = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING *;