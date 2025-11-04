-- queries/sql/mutationsUsers.sql

-- name: CreateUser :exec
INSERT INTO users (name, password, email, creation_date)
VALUES (?, ?, ?, ?);

-- name: GetCreatedUser :one
SELECT * FROM users WHERE id = LAST_INSERT_ID();

-- name: UpdateUser :exec
UPDATE users
SET name = COALESCE(?, name),
    email = COALESCE(?, email)
WHERE id = ?;

-- name: GetUpdatedUser :one
SELECT * FROM users WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
 