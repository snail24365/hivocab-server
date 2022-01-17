-- name: InsertUsecase :one
INSERT INTO usecase (
  id,
  word_id,
  description_sentence
) VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: GetUsecaseByWord :many
SELECT * FROM usecase
WHERE word_id = $1;

-- name: GetUsecaseById :one
SELECT * FROM usecase
WHERE id = $1;

-- name: CountAllUsecase :one
SELECT COUNT(*) FROM usecase;

-- name: DeleteAllUsecase :exec
DELETE FROM usecase;
