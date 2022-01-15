-- name: CreateExample :one
INSERT INTO example (
  usecase_id,
  sentence
) VALUES (
  $1,
  $2
)
RETURNING *;

-- name: ListExampleByUsecase :many
SELECT * FROM example
WHERE usecase_id = $1;