package controllers

import (
	"JudgeHost/src/common"
	"JudgeHost/src/config/configuration"
	"JudgeHost/src/models/dto"
	"JudgeHost/src/service"
	"JudgeHost/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var JudgeService *service.JudgeService = service.NewJudgeService(configuration.JudgeEnvironmentConfigurationEntity)

func LoadJudgeControllers(e *gin.Engine) {
	judgeGroup := e.Group("/judge")
	judgeGroup.POST("/result", RunJudge)
}

func RunJudge(context *gin.Context) {
	judgeDTO := dto.JudgeDTO{}
	if err := context.ShouldBind(&judgeDTO); err != nil {
		context.JSON(500, gin.H{"msg": err})
		return
	}
	if err := util.ValidateStructCheck(&judgeDTO); err != nil {
		// TODO logger
		fmt.Println(err)
		context.JSON(500, gin.H{"msg": err})
		return
	}
	context.JSON(http.StatusOK, common.NewUnifiedResponseMessgaeData("Respone sueccess", judgeDTO))
}
