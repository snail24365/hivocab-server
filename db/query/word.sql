-- name: InsertWord :one
INSERT INTO word (
  id,
  spelling
) VALUES (
  $1,$2
)
RETURNING *;

-- name: GetWordBySpelling :one
SELECT * FROM word
WHERE spelling = $1 LIMIT 1;

-- name: GetWords :many
SELECT * FROM word LIMIT $1 OFFSET $2;

-- name: GetWordById :one
SELECT * FROM word
WHERE id = $1;

-- name: GetWordByPage :many
SELECT * FROM word
LIMIT $1
OFFSET $2;

-- name: CountAllWord :one
SELECT COUNT(*) FROM word;

-- name: DeleteAllWord :exec
DELETE FROM word;
