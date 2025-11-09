-- name: CreateBoard :exec
INSERT INTO boards (id, name)
VALUES (?, ?);

-- name: GetCreatedBoard :one
SELECT * FROM boards WHERE id = LAST_INSERT_ID();

-- name: GetBoard :one
SELECT * FROM boards
WHERE id = ? LIMIT 1;

-- name: UpdateBoard :exec
UPDATE boards
SET 
    name = COALESCE(sqlc.narg(name), name)
WHERE id = sqlc.arg(id);

-- name: ListBoards :many
SELECT * FROM boards
ORDER BY name;

-- name: DeleteBoard :exec
DELETE FROM boards WHERE id = ?;

----------------------------------

-- name: GetBoardForColumn :one
SELECT b.*
FROM boards b
JOIN columns c ON b.id = c.board_id
WHERE c.id = ?;
