package controllers

import (
	"JudgeHost/src/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetWelcomeText(context *gin.Context) {
	context.JSON(http.StatusOK,
		common.NewUnifiedResponseMessgae("Your project is running successfully! O(∩_∩)O"))
}

func LoadBaseController(e *gin.Engine) {
	e.GET("/", GetWelcomeText)
}
