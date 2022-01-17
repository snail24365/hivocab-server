-- name: InsertExample :one
INSERT INTO example (
  id,
  usecase_id,
  sentence
) VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: ListExampleByUsecase :many
SELECT * FROM example
WHERE usecase_id = $1;

-- name: CountAllExample :one
SELECT COUNT(*) FROM example;

-- name: DeleteAllExample :exec
DELETE FROM example;