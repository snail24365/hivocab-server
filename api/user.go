package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/snail24365/hivocab-server/db/sqlc"
)

func (server *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest 
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	user, err := createUser(ctx, server.store,req)

	if err != nil {
		ctx.JSON(http.StatusConflict, errorResponse(err))
	}
	
	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}



type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func createUser(ctx context.Context, store db.Store, user CreateUserRequest) (db.User, error) {	
	
	store.GetUserByUsername(ctx, user.Username)

	return store.InsertUser(ctx, db.InsertUserParams{
		Username: user.Username,
		Password: user.Password,
	})

}