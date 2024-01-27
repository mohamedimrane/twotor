-- name: CreateTwoot :one
INSERT INTO twoots ( contents )
VALUES ( ? )
RETURNING *;

-- name: GetTwootById :one
SELECT * FROM twoots
WHERE id = ? LIMIT 1;

-- name: ListTwoots :many
SELECT * FROM twoots;
