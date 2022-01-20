// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package db

import (
	"context"
)

const getUserById = `-- name: GetUserById :one
SELECT id, username, password, latest_visit, study_amount, study_goal, password_changed_at, created_at, study_index FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.LatestVisit,
		&i.StudyAmount,
		&i.StudyGoal,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.StudyIndex,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, password, latest_visit, study_amount, study_goal, password_changed_at, created_at, study_index FROM users WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.LatestVisit,
		&i.StudyAmount,
		&i.StudyGoal,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.StudyIndex,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO users (
  username,
  password
) VALUES (
  $1,
  $2
)
RETURNING id, username, password, latest_visit, study_amount, study_goal, password_changed_at, created_at, study_index
`

type InsertUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser, arg.Username, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.LatestVisit,
		&i.StudyAmount,
		&i.StudyGoal,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.StudyIndex,
	)
	return i, err
}

const moveNextExercise = `-- name: MoveNextExercise :one
UPDATE users 
SET study_index = $1
WHERE username = $2
RETURNING id, username, password, latest_visit, study_amount, study_goal, password_changed_at, created_at, study_index
`

type MoveNextExerciseParams struct {
	StudyIndex int64  `json:"study_index"`
	Username   string `json:"username"`
}

func (q *Queries) MoveNextExercise(ctx context.Context, arg MoveNextExerciseParams) (User, error) {
	row := q.db.QueryRowContext(ctx, moveNextExercise, arg.StudyIndex, arg.Username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.LatestVisit,
		&i.StudyAmount,
		&i.StudyGoal,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.StudyIndex,
	)
	return i, err
}
