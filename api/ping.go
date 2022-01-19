package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) AuthPing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}