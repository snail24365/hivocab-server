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

-- name: CountWritingsGroupByCreateAt :many
SELECT EXTRACT (DAY FROM created_at) as Day, COUNT(user_id) as Count 
FROM 
(SELECT * FROM writing 
ORDER BY EXTRACT (MONTH FROM created_at) ASC, EXTRACT (DAY FROM created_at) ASC
) AS ordered_table
WHERE user_id = $1 
AND created_at >= $2 
AND created_at <= $3
GROUP BY EXTRACT (DAY FROM created_at);

-- name: DeleteWriting :exec
DELETE FROM writing WHERE id = $1;

-- name: GetWritingsById :one
SELECT * FROM writing WHERE id = $1;