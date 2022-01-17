package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/snail24365/hivocab-server/db/sqlc"
)

type Exercise struct {
	Word 			db.Word 				`json:"word"`
	Usecase 	db.Usecase 		`json:"usecase"`
	Examples  []db.Example   `json:"examples"`
}

type GetExerciseRequest struct {
	UserId int `json:"user_id"`
}

func (server *Server) GetExercise(ctx *gin.Context) {
	var req GetExerciseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	exercise, err := GetExercise(ctx, server.store, req)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}


func GetExercise(ctx context.Context, store db.Store, req GetExerciseRequest) (Exercise, error) {	
	exercise := Exercise{}
	/*
	user, err := store.GetUserById(ctx, int64(userId))
	if err != nil {
		return exercise, err
	}
	*/
	
	//user.StudyIndex
	var err error
	exercise.Usecase, err = store.GetUsecaseById(ctx, 0)
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