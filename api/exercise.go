package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/snail24365/hivocab-server/db/sqlc"
	"github.com/snail24365/hivocab-server/token"
)

type Exercise struct {
	Word 			db.Word 			`json:"word"`
	Usecase 	db.Usecase 		`json:"usecase"`
	Examples  []db.Example   `json:"examples"`
}

func (server *Server) GetExercise(ctx *gin.Context) {
	exercise, err := GetExercise(ctx, server.store)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, exercise)
}


func (server *Server) NextExercise(ctx *gin.Context) {
username := ctx.GetString(authorizationUsernameKey)
	user, err := server.store.GetUserByUsername(ctx, username)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, nil)
		return
	}
	
	numUsecase, err := server.store.CountAllUsecase(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, nil)
		return
	}

	nextIndex := (user.StudyIndex + 1) % numUsecase
	_, err = server.store.MoveNextExercise(ctx, db.MoveNextExerciseParams{
		nextIndex,
		username,
	})

	

	exercise, err := GetExercise(ctx, server.store)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}


func GetExercise(ctx *gin.Context, store db.Store) (Exercise, error) {	
	exercise := Exercise{}
	
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	username := authPayload.Username
	
	user, err := store.GetUserByUsername(ctx, username)
	if err != nil {
		return exercise, err
	}
	
	usecaseId := user.StudyIndex
	exercise.Usecase, err = store.GetUsecaseById(ctx, usecaseId)
	if err != nil {
		return exercise, err
	}
	
	exercise.Word, err = store.GetWordById(ctx, exercise.Usecase.WordID)
	if err != nil {
		return exercise, err
	}

	exercise.Examples, err = store.ListExampleByUsecase(ctx, exercise.Usecase.ID)
	if err != nil {
		return exercise, err
	}

	return exercise, err
}