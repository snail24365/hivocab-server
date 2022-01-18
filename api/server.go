package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/snail24365/hivocab-server/db/sqlc"
	"github.com/snail24365/hivocab-server/token"
	"github.com/snail24365/hivocab-server/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.SQLStore)	 (*Server, error)  {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	
	router := gin.Default()
	router.POST("/user", server.CreateUser)
	router.POST("/login", server.Login)
	

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/exercise", server.GetExercise)
	authRoutes.GET("/word", server.GetWords)
	//authRoutes.POST("/logout", server.Logout)
	//authRoutes.GET("/report", server.AnalysisStudy)
	//authRoutes.POST("/writing", server.PostWriting)
	//authRoutes.DELETE("/writing", server.DeleteWriting)
	

	server.router = router
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

