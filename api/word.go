package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/snail24365/hivocab-server/db/sqlc"
)

type GetWordsParams struct{
	PageSize 		int 	`json:"pageSize"`
	PageOffset 	int 	`json:"pageOffset"`
}

func (server *Server) GetCountTotalWord(ctx *gin.Context) {
	count, err := server.store.CountAllWord(ctx);

	if err != nil {
		ctx.JSON(http.StatusBadGateway, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, count)
}

func (server *Server) GetWord(ctx *gin.Context) {
	page, err := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10 ,32) 
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	offset := (page - 1)  * pageSize
	words, err := server.store.GetWordByPage(ctx, db.GetWordByPageParams{
		Limit: pageSize,
		Offset: int32(offset),
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, words)
}

const pageSize = 20