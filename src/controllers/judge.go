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
	"sync"
)

var JudgeService *service.JudgeService = service.NewJudgeService(configuration.JudgeEnvironmentConfigurationEntity)

type JudgeTask struct {
	JudgeDTO    *dto.JudgeDTO
	JudgeResult []*dto.SingleJudgeResultDTO
	Message     string
}

func (j *JudgeTask) Do() {
	j.JudgeResult = JudgeService.RunJudge(j.JudgeDTO)
	if len(j.JudgeResult) > 0 {
		j.JudgeResult[0].SetMessage()
		j.Message = j.JudgeResult[0].Message
	}

}

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
	judgeTask := JudgeTask{
		JudgeDTO: &judgeDTO,
	}
	judgeTask.JudgeResult = make([]*dto.SingleJudgeResultDTO, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	judgeTaskWrop := configuration.NewTaskWrop(&judgeTask, &wg)
	configuration.JudgeExecutorPool.Invoke(judgeTaskWrop)
	wg.Wait()
	context.JSON(http.StatusOK, common.NewUnifiedResponseMessgaeData("judge result", judgeTask.JudgeResult))
}
