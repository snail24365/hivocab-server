-- name: GetWritingsByUserIdAndUsecaseId :many
SELECT * FROM writing WHERE user_id = $1 AND usecase_id = $2;

-- name: InsertWriting :one
INSERT INTO writing (
  writing,
  usecase_id,
  user_id
) VALUES (
  $1,
  $2,
  $3
)
RETURNING *;


