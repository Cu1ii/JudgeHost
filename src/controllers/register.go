package controllers

import (
	"github.com/gin-gonic/gin"
)

func LoadControllers(e *gin.Engine) {
	LoadBaseController(e)
	LoadFileController(e)
	LoadCommonController(e)
	LoadJudgeControllers(e)
}
