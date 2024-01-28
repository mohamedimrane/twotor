-- name: CreateTwoot :one
INSERT INTO twoots ( contents, user_ID, created_at )
VALUES ( ?, ?, DATETIME("now") )
RETURNING *;

-- name: GetTwootById :one
SELECT sqlc.embed(twoots), sqlc.embed(users) FROM twoots
INNER JOIN users ON twoots.user_ID = users.ID
WHERE users.ID = ? LIMIT 1;

-- name: ListTwoots :many
SELECT sqlc.embed(twoots), sqlc.embed(users) FROM twoots
INNER JOIN users ON twoots.user_ID = users.ID;

