-- name: CreateColumn :exec
INSERT INTO columns (id, name, `order`, board_id)
VALUES (?, ?, ?, ?);

-- name: GetCreatedColumn :one
SELECT * FROM columns WHERE id = LAST_INSERT_ID();

-- name: GetColumn :one
SELECT * FROM columns
WHERE id = ? LIMIT 1;

-- name: UpdateColumn :exec
UPDATE columns
SET 
    name = COALESCE(sqlc.narg(name), name),
    `order` = COALESCE(sqlc.narg("order"), `order`),
    board_id = COALESCE(sqlc.narg(board_id), board_id)
WHERE id = sqlc.arg(id);

-- name: ListColumns :many
SELECT * FROM columns
ORDER BY name;

-- name: DeleteColumn :exec
DELETE FROM columns WHERE id = ?;

----------------------------------

-- name: ListColumnsByBoard :many
SELECT * FROM columns
WHERE board_id = ?
ORDER BY `order` ASC;

-- name: GetColumnForTask :one
SELECT c.* FROM columns c
JOIN tasks t ON c.id = t.column_id
WHERE t.id = ?;