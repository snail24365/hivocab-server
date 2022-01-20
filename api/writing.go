package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/snail24365/hivocab-server/db/sqlc"
)

func (server *Server) FetchWriting(ctx *gin.Context) {
	store := server.store
	userId, exists := ctx.Get(authorizationUserIdKey)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, nil)
		return 
	}
	
	usecaseId, err := strconv.ParseInt(ctx.Query("usecaseId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return 
	}

	uid := userId.(int64)


	writings, err := store.GetWritingsByUserIdAndUsecaseId(ctx, db.GetWritingsByUserIdAndUsecaseIdParams{
		UserID: uid,
		UsecaseID:  usecaseId,
	})	

	if err != nil && err != sql.ErrNoRows {
		ctx.JSON(http.StatusBadGateway, nil)
		return
	}
	ctx.JSON(http.StatusOK, writings)
}

func (server *Server) EnrollWriting(ctx *gin.Context) {
	store := server.store

	var req db.InsertWritingParams
	userId, exists := ctx.Get(authorizationUserIdKey)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, nil)
		return 
	}

	// usecaseId, err := strconv.ParseInt(ctx.Query("usecaseId"), 10, 64)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, nil)
	// 	return 
	// }
		
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	req.UserID = userId.(int64)
	result, err := store.InsertWriting(ctx, req)
	fmt.Print(result)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, result)
}

func (server *Server) DeleteWriting(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}