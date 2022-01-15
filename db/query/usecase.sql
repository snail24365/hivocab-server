-- name: CreateUsecase :one
INSERT INTO usecase (
  word_id,
  description_sentence
) VALUES (
  $1,
  $2
)
RETURNING *;

-- name: ListUsecaseByWord :many
SELECT * FROM usecase
WHERE word_id = $1;