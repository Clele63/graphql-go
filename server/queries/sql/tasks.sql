-- name: CreateTask :exec
INSERT INTO tasks (id, title, description, column_id)
VALUES (?, ?, ?, ?);

-- name: GetCreatedTask :one
SELECT * FROM tasks WHERE id = LAST_INSERT_ID();

-- name: GetTask :one
SELECT * FROM tasks
WHERE id = ? LIMIT 1;

-- name: UpdateTask :exec
UPDATE tasks
SET 
  title = COALESCE(sqlc.narg(title), title),
  description = COALESCE(sqlc.narg(description), description)
WHERE id = sqlc.arg(id);

-- name: ListTasks :many
SELECT * FROM tasks
ORDER BY title;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = ?;

----------------------------------

-- name: ListTasksForColumnPaginated :many
SELECT * FROM tasks
WHERE 
  column_id = ? 
  AND id > ?
ORDER BY id ASC
LIMIT ?;

-- name: GetTaskForComment :one
SELECT t.* FROM tasks t
JOIN comments c ON t.id = c.task_id
WHERE c.id = ?;

-- name: MoveTask :exec
UPDATE tasks
SET 
  column_id = COALESCE(sqlc.narg(column_id), column_id)
WHERE id = sqlc.arg(id);