package controllers

import (
	"JudgeHost/src/common"
	"JudgeHost/src/config/configuration"
	"JudgeHost/src/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoadFileController(e *gin.Engine) {
	fileGroup := e.Group("/file")
	fileGroup.GET("/submission/:submissionId", DownloadSubmissionById)
	fileGroup.DELETE("/submission_path", ClearSubmissionFiles)
	fileGroup.DELETE("/solution_path", ClearSolutionPath)
}

var FileServce = service.NewFileService(configuration.JudgeEnvironmentConfigurationEntity)

// DownloadSubmissionById 下载某次提交的相关信息
func DownloadSubmissionById(context *gin.Context) {
	submissionId := context.Param("submissionId")
	zippedPath, err := FileServce.GetSubmissionDataById(submissionId)
	if err != nil {
		log.Println(err.Error())
	}
	context.Header("Content-Type", "application/octet-stream")
	context.Header("Content-Disposition", "attachment; filename="+submissionId+".zip")
	context.File(zippedPath)
}

// ClearSubmissionFiles 删除所有用户提交的代码文件
func ClearSubmissionFiles(context *gin.Context) {
	if _, err := FileServce.ClearSubmissionPath(); err != nil {
		context.JSON(http.StatusBadRequest, common.NewUnifiedResponseMessgaeData("B0001", "删除失败"))
		return
	}
	context.JSON(http.StatusOK, common.NewUnifiedResponseMessgae("删除成功"))
}

// ClearSolutionPath 删除所有用户提交的代码文件
func ClearSolutionPath(context *gin.Context) {
	if _, err := FileServce.ClearSolutionPath(); err != nil {
		context.JSON(http.StatusBadRequest, common.NewUnifiedResponseMessgaeData("B0001", "删除失败"))
	}
	context.JSON(http.StatusOK, common.NewUnifiedResponseMessgae("删除所有用户提交的代码文件 删除成功"))
}
