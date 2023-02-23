package controllers

import (
	"JudgeHost/src/judge"
	"JudgeHost/src/models/dto"
	"JudgeHost/src/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func LoadJudgeControllers(e *gin.Engine) {
	judgeGroup := e.Group("/judge")
	judgeGroup.POST("/result", RunJudge)
}

func RunJudge(context *gin.Context) {
	judgeRequest := dto.JudgeRequest{}
	if err := context.ShouldBind(&judgeRequest); err != nil {
		context.JSON(500, gin.H{"msg": err})
		return
	}
	if err := util.ValidateStructCheck(&judgeRequest); err != nil {
		logrus.Debug("ValidateStructCheck error", err)
		context.JSON(500, gin.H{"msg": err})
		return
	}
	var res string
	err := judge.RunJudge(&judgeRequest, &res)
	if err != nil {
		logrus.Error("")
	}
	context.JSON(http.StatusOK, gin.H{"msg": res})
}
