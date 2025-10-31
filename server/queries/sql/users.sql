-- queries/sql/users.sql

-- name: ListUsers :many
SELECT id, name, email, creation_date FROM users ORDER BY id;

-- name: GetUserByID :one
SELECT id, name, email, creation_date FROM users WHERE id = ?;

-- name: SearchUsersByName :many
SELECT id, name, email, creation_date FROM users WHERE LOWER(name) LIKE LOWER(?);

-- name: GetUserAuthByName :one
SELECT id, name, password FROM users WHERE name = ?;
