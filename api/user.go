package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/snail24365/hivocab-server/db/sqlc"
	"github.com/snail24365/hivocab-server/util"
)

type UsernamePassword struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required"`
}

	func (server *Server) CreateUser(ctx *gin.Context) {
	var req UsernamePassword
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	existing, err := server.store.GetUserByUsername(ctx, req.Username)

	if err != nil && err != sql.ErrNoRows  {
		ctx.JSON(http.StatusBadGateway, errorResponse(err))
		return 		
	}
	
	if existing.Username == req.Username {
		ctx.JSON(http.StatusConflict, errorResponse(errors.New("username already exists")))
		return
	} 

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := server.store.InsertUser(ctx, db.InsertUserParams{
		Username: req.Username,
		Password: hashedPassword,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, errorResponse(err))
	}
	
	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

func (server *Server) Login(ctx *gin.Context) {
	var req UsernamePassword 
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError , errorResponse(err))
		return 
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		Username: user.Username,
	}

	ctx.JSON(http.StatusOK, rsp)
}