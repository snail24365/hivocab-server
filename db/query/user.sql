-- name: GetUserById :one
SELECT * FROM user WHERE id = $1;

-- name: CreateUser :one
INSERT INTO user (
  username,
  password
) VALUES (
  $1,
  $2
)
RETURNING *;