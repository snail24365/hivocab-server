package service

import (
	"context"

	db "github.com/snail24365/hivocab-server/db/sqlc"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(ctx context.Context, store db.Store, user CreateUserRequest) (db.User, error) {	
	
	store.GetUserByUsername(ctx, user.Username)

	return store.InsertUser(ctx, db.InsertUserParams{
		Username: user.Username,
		Password: user.Password,
	})
	
}