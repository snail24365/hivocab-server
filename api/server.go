package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/snail24365/hivocab-server/db/sqlc"
)

type Server struct {
	store *db.Store
	router *gin.Engine	
}

// 서버 생성 및 라우트 설정
func NewServer(store *db.Store)	*Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/word", server.createWord)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}