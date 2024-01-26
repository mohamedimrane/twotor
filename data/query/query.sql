-- name: CreateUser :one
INSERT INTO users ( username, email, password, display_name, bio )
VALUES ( ?, ?, ?, ?, ? )
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;
