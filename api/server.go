package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	//store *db.Store
	router *gin.Engine	
}
