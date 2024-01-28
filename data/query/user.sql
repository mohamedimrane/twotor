-- name: CreateUser :one
INSERT INTO users ( username, email, password, display_name, bio, created_at, updated_at )
VALUES ( ?, ?, ?, ?, ?, DATETIME("now"), DATETIME("now") )
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = ? LIMIT 1;
