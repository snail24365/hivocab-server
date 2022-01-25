package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/snail24365/hivocab-server/db/sqlc"
	"github.com/snail24365/hivocab-server/util"
)

type UsernamePassword struct {
	Username string `json:"username" binding:"required,alphanum,min=5"`
	Password string `json:"password" binding:"required,min=8"`
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req UsernamePassword
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	existing, err := server.store.GetUserByUsername(ctx, req.Username)

	if err != nil && err != sql.ErrNoRows  {
		ctx.JSON(http.StatusBadGateway, errorResponse(err))
		return 		
	}
	
	if existing.Username == req.Username {
		ctx.JSON(http.StatusConflict, errorResponse(errors.New("username already exists")))
		return
	} 

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := server.store.InsertUser(ctx, db.InsertUserParams{
		Username: req.Username,
		Password: hashedPassword,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, errorResponse(err))
	}
	
	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

func (server *Server) Login(ctx *gin.Context) {
	var req UsernamePassword 
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError , errorResponse(err))
		return 
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		Username: user.Username,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) GetStudyInfo(ctx *gin.Context) {
	userId, exist := ctx.Get(authorizationUserIdKey)
	if !exist {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}
	studyInfo, err := server.store.GetStudyInfoById(ctx, userId.(int64))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, nil)
		return
	}
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Seoul")
	todayEnd := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, loc);
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc);
	lastWeek := todayStart.AddDate(0, 0, -7)

	a1 := todayEnd.Day()
	b1 := todayEnd.Hour()
	fmt.Print(a1,b1)
	
	weekWritingStatistics, err := server.store.CountWritingsGroupByCreateAt(ctx, db.CountWritingsGroupByCreateAtParams{
		UserID: userId.(int64), 
		CreatedAt: lastWeek,
		CreatedAt_2: todayEnd,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, nil)
		return
	}
	
	weekAgo := time.Now().AddDate(0, 0, -6)
	var attendanceTable []AttendanceCell
	
	oneWeek := 7
	for offset := 0; offset < oneWeek; offset++ {
		day := weekAgo.AddDate(0, 0, offset).Day()
		
		attendanceTable = append(attendanceTable, AttendanceCell{
			Day: day,
			Count: 0,
		})
	}

	for _, dayWritingStatistics := range weekWritingStatistics {
		for i, x := range attendanceTable {
			if int(dayWritingStatistics.Day) == x.Day {
				attendanceRef := &attendanceTable[i]
				attendanceRef.Count = int(dayWritingStatistics.Count)
				//attendanceTable[i].Count = x.Count
			}
		}			
	}

	todayStudyAmount := attendanceTable[len(attendanceTable)-1].Count


	ctx.JSON(http.StatusOK, GetStudyInfoResponse{
		StudyGoal: studyInfo.StudyGoal,
		StudyAmount: int32(todayStudyAmount),
		LatestVisit: studyInfo.LatestVisit,
		CountByDay: attendanceTable,
	})
}

type AttendanceCell struct {
	Day   int `json:"day"`
	Count int   `json:"count"`
}

type GetStudyInfoResponse struct {
	StudyGoal   int32     `json:"study_goal"`
	StudyAmount int32     `json:"study_amount"`
	LatestVisit time.Time `json:"latest_visit"`
	CountByDay []AttendanceCell `json:"writing_count_per_day"`
}

func (server *Server) UpdateGoal(ctx *gin.Context) {
	userId, exist := ctx.Get(authorizationUserIdKey)
	if !exist {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	var req UpdateGoalReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	
	updated, err := server.store.UpdateGoal(ctx, db.UpdateGoalParams{
		ID: userId.(int64),
		StudyGoal: req.StudyGoal,
	})
	
	if err != nil {
		ctx.JSON(http.StatusBadGateway, nil)
		return
	}
	updated.Password = ""
	ctx.JSON(http.StatusOK, updated)
}

type UpdateGoalReq struct {
	StudyGoal   int32     `json:"study_goal"`
}