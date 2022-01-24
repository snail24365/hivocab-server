-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: InsertUser :one
INSERT INTO users (
  username,
  password
) VALUES (
  $1,
  $2
)
RETURNING *;

-- name: MoveNextExercise :one
UPDATE users 
SET study_index = $1
WHERE username = $2
RETURNING *;


-- name: GetStudyInfoById :one
SELECT study_goal,study_amount,latest_visit FROM users WHERE id = $1;


