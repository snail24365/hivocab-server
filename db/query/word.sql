-- name: CreateWord :one
INSERT INTO word (
  spelling
) VALUES (
  $1
)
RETURNING *;

-- name: GetWordBySpelling :one
SELECT * FROM word
WHERE spelling = $1 LIMIT 1;

-- name: ListWordByPage :many
SELECT * FROM word
LIMIT $1
OFFSET $2;

-- name: CountAllWords :one
SELECT COUNT(*) FROM word;
