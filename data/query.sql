-- name: CreateUser :one
INSERT INTO users ( username, email, password, display_name, bio )
VALUES ( ?, ?, ?, ?, ? )
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;
