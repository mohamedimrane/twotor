-- name: CreateTwoot :one
INSERT INTO twoots ( contents, user_ID, twoot_ID, created_at )
VALUES ( ?, ?, ?, DATETIME("now") )
RETURNING *;

-- name: GetTwootById :one
SELECT sqlc.embed(twoots), sqlc.embed(users) FROM twoots
INNER JOIN users ON twoots.user_ID = users.ID
WHERE twoots.ID = ? LIMIT 1;

-- name: GetTwootsFromTwootWithId :many
SELECT sqlc.embed(twoots), sqlc.embed(users) FROM twoots
INNER JOIN users ON twoots.user_ID = users.ID
WHERE twoots.twoot_ID = ?;

-- name: GetTwootsByUserId :many
SELECT * FROM twoots
WHERE user_ID = ?;

-- name: ListTwoots :many
SELECT sqlc.embed(twoots), sqlc.embed(users) FROM twoots
INNER JOIN users ON twoots.user_ID = users.ID
WHERE twoots.twoot_ID IS ?;

