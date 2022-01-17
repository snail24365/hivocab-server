package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/snail24365/hivocab-server/db/sqlc"
)

type GetWordsParams struct{
	PageSize 		int 	`json:"pageSize"`
	PageOffset 	int 	`json:"pageOffset"`
}

func (server *Server) GetWords(ctx *gin.Context) {
	var req GetWordsParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	words, err := GetWords(ctx, server.store, req);
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, words)
}

func GetWords(ctx context.Context, store db.Store, info GetWordsParams) ([]db.Word, error) {	
	offset := info.PageSize * info.PageOffset
	return store.GetWords(ctx, db.GetWordsParams{Limit: int32(info.PageSize), Offset: int32(offset)})
}