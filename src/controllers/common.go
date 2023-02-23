package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadCommonController(e *gin.Engine) {
	commonGroup := e.Group("/common")
	commonGroup.GET("/connect", connect)
}

func connect(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "connect success!"})
}
