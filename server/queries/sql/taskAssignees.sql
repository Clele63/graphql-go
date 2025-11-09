-- name: AddTaskAssignee :exec
INSERT INTO task_assignees (task_id, user_id)
VALUES (?, ?);

-- name: GetTaskAssignee :one
SELECT * FROM task_assignees
WHERE task_id = ? AND user_id = ? LIMIT 1;

-- name: ListAssigneesByTask :many
SELECT u.* FROM users u
JOIN task_assignees ta ON u.id = ta.user_id
WHERE ta.task_id = ?;

-- name: ListTaskAssignsByTask :many
SELECT * FROM task_assignees
WHERE task_id = ?;

-- name: ListTaskAssignsByUser :many
SELECT * FROM task_assignees
WHERE user_id = ?;

-- name: DeleteTaskAssignee :exec
DELETE FROM task_assignees
WHERE task_id = ? AND user_id = ?;

-- name: ClearTaskAssigneesByTask :exec
DELETE FROM task_assignees
WHERE task_id = ?;

-- name: ClearTaskAssigneesByUser :exec
DELETE FROM task_assignees
WHERE user_id = ?;