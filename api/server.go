package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/snail24365/hivocab-server/db/sqlc"
)

type Server struct {
	store db.Store
	router *gin.Engine	
}

func NewServer(store db.Store)	*Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/exercise", server.GetExercise)
	router.POST("/user", server.CreateUser)
	//router.POST("/login", server.Login)
	//router.POST("/logout", server.Logout)
	router.GET("/word", server.GetWords)
	//router.GET("/report", server.AnalysisStudy)
	//router.POST("/writing", server.PostWriting)
	//router.DELETE("/writing", server.DeleteWriting)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

