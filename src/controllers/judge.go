package controllers

import (
	"JudgeHost/src/common"
	"JudgeHost/src/config/configuration"
	"JudgeHost/src/models/dto"
	"JudgeHost/src/service"
	"JudgeHost/src/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func LoadJudgeControllers(e *gin.Engine) {
	judgeGroup := e.Group("/judge")
	judgeGroup.POST("/result", RunJudge)
}

type JudgeTask struct {
	Name        string
	JudgeDTO    *dto.JudgeDTO
	JudgeResult []*dto.SingleJudgeResultDTO
	Message     string
	Err         error
	done        chan int
}

func NewJudgeTask(judgeDTO *dto.JudgeDTO) *JudgeTask {
	return &JudgeTask{
		Name:     configuration.JudgeNamePrefix + strconv.FormatInt(configuration.JudgeIndexId.Add(1), 10),
		JudgeDTO: judgeDTO,
		done:     make(chan int),
	}
}

func (j *JudgeTask) Do() {
	j.JudgeResult, j.Err = JudgeService.RunJudge(j.JudgeDTO)
	if len(j.JudgeResult) > 0 {
		j.JudgeResult[0].SetMessage()
		j.Message = j.JudgeResult[0].Message
	}
	j.done <- 1
}

func (j *JudgeTask) GetName() string {
	return j.Name
}

func (j *JudgeTask) Wait() {
	_ = <-j.done
}

var JudgeService = service.NewJudgeService(configuration.JudgeEnvironmentConfigurationEntity)

func RunJudge(context *gin.Context) {
	judgeDTO := dto.JudgeDTO{}
	if err := context.ShouldBind(&judgeDTO); err != nil {
		context.JSON(500, gin.H{"msg": err})
		return
	}
	if err := util.ValidateStructCheck(&judgeDTO); err != nil {
		logrus.Debug("ValidateStructCheck error", err)
		context.JSON(500, gin.H{"msg": err})
		return
	}
	judgeTask := NewJudgeTask(&judgeDTO)
	configuration.JudgeExecutorPool.Invoke(judgeTask)
	judgeTask.Wait()
	if judgeTask.Err != nil {
		logrus.Debug("ValidateStructCheck error", judgeTask.Err)
		context.JSON(http.StatusInternalServerError,
			common.NewUnifiedResponseMessgaeData("judge result"+judgeTask.Err.Error()+" ", judgeTask.JudgeResult))
	}
	context.JSON(http.StatusOK, common.NewUnifiedResponseMessgaeData("judge result", judgeTask.JudgeResult))
}
