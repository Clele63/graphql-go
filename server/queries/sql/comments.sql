-- name: CreateComment :exec
INSERT INTO comments (id, content, author_id, task_id)
VALUES (?, ?, ?, ?);

-- name: GetCreatedComment :one
SELECT * FROM comments WHERE id = LAST_INSERT_ID();

-- name: GetComment :one
SELECT * FROM comments
WHERE id = ? LIMIT 1;

-- name: UpdateComment :exec
UPDATE comments
SET 
  content = COALESCE(sqlc.narg(content), content)
WHERE id = sqlc.arg(id);

-- name: ListComments :many
SELECT * FROM comments;

-- name: DeleteComment :exec
DELETE FROM comments WHERE id = ?;

----------------------------------

-- name: GetCommentForTask :one
SELECT * FROM comments
WHERE task_id = ?
ORDER BY created_at ASC
LIMIT 1;

-- name: ListCommentsByTask :many
SELECT * FROM comments
WHERE task_id = ?
ORDER BY created_at ASC;

-- name: GetCommentAuthor :one
SELECT author_id FROM comments
WHERE id = ? LIMIT 1;