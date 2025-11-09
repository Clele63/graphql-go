-- name: CreateUser :exec
INSERT INTO users (id, name, email, password, avatar)
VALUES (?, ?, ?, ?, NULL);

-- name: GetCreatedUser :one
SELECT * FROM users WHERE id = LAST_INSERT_ID();

-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: UpdateUser :exec
UPDATE users
SET 
    name = COALESCE(sqlc.narg(name), name),
    email = COALESCE(sqlc.narg(email), email)
WHERE id = sqlc.arg(id);

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;

----------------------------------

-- name: SearchUsers :many
SELECT * FROM users
WHERE name LIKE ? OR email LIKE ?
ORDER BY name;

-- name: GetUserForComment :one
SELECT u.* FROM users u
JOIN comments c ON u.id = c.author_id
WHERE c.id = ?;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;
