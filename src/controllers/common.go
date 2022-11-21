package controllers

import (
	"JudgeHost/src/common"
	"JudgeHost/src/config/configuration"
	"JudgeHost/src/config/global"
	"JudgeHost/src/models/dto"
	"JudgeHost/src/models/vo"
	"JudgeHost/src/service"
	"JudgeHost/src/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func LoadCommonController(e *gin.Engine) {
	commonGroup := e.Group("/common")
	commonGroup.GET("/connection_test", TestConnection)
	commonGroup.PUT("/max_working_amount", SetMaxWorkingAmount)
}

var CommonService = service.NewCommonService(configuration.JudgeEnvironmentConfigurationEntity, configuration.JudgeExecutorPool, global.GlobalApp.App.Port, "v0.1")

func TestConnection(context *gin.Context) {
	hostInfo := CommonService.GetJudgeHostConfiguration()
	conditionBO := CommonService.GetJudgeHostCondition()
	conditionVO := &vo.JudgeHostConditionVO{
		WorkPath:             hostInfo.WorkPath,
		ScriptPath:           hostInfo.ScriptPath,
		ResolutionPath:       hostInfo.ResolutionPath,
		Port:                 hostInfo.Port,
		WorkingAmount:        conditionBO.WorkingAmount,
		CpuCoreAmount:        conditionBO.CpuCoreAmount,
		MemoryCostPercentage: conditionBO.CpuCostPercentage,
		CpuCostPercentage:    conditionBO.CpuCostPercentage,
		QueueAmount:          conditionBO.QueueAmount,
		MaxWorkingAmount:     conditionBO.MaxWorkingAmount,
		Version:              hostInfo.Version,
	}
	context.JSON(http.StatusOK, common.NewUnifiedResponseMessgaeData("JudgeHost running normally", conditionVO))
}

func SetMaxWorkingAmount(context *gin.Context) {
	workingAmountDTO := dto.SetWorkingAmountDTO{}
	if err := context.ShouldBind(&workingAmountDTO); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	if err := util.ValidateStructCheck(workingAmountDTO); err != nil {
		logrus.Debug("ValidateStructCheck error", err)
		context.JSON(500, gin.H{"msg": err})
	}
	CommonService.SetJudgeHostWorkingAmount(workingAmountDTO.MaxWorkingAmount, workingAmountDTO.ForceSet)
}
