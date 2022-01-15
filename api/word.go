package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createWordRequest struct {
	Spelling string `json:"spelling" binding:"required"`
}

func (server *Server) createWord(ctx *gin.Context) {
	var req createWordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	word, _ := server.store.CreateWord(ctx, req.Spelling)
	ctx.JSON(http.StatusOK, word)
}