package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snail24365/hivocab-server/api/service"
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
	//router.GET("/word", server.GetWords)
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

func (server *Server) GetExercise(ctx *gin.Context) {
	var req service.GetExerciseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	exercise, err := service.GetExercise(ctx, server.store, req)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req service.CreateUserRequest 
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	user, err := service.CreateUser(ctx, server.store, service.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusConflict, errorResponse(err))
	}
	
	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}

