package service

import (
	"JudgeHost/src/config/configuration"
	"JudgeHost/src/util"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type FileService struct {
	JudgeEnvironmentConfiguration *configuration.JudgeEnvironmentConfiguration
}

func NewFileService(judgeEnvironmentConfiguration *configuration.JudgeEnvironmentConfiguration) *FileService {
	return &FileService{JudgeEnvironmentConfiguration: judgeEnvironmentConfiguration}
}

// GetSubmissionDataById 读取某次提交的文件夹，将内容压缩过之后返回
func (f *FileService) GetSubmissionDataById(submissionId string) (string, error) {
	submissionPath := f.GetSubmissionPathById(submissionId)
	return f.ZipSubmissionFolder(submissionPath)
}

// GetSubmissionPathById 获取某次提交的工作目录
func (f *FileService) GetSubmissionPathById(submissionId string) string {
	return f.JudgeEnvironmentConfiguration.JudgeEnvironment.SubmissionPath + "/" + submissionId
}

// ZipSubmissionFolder 返回压缩后的压缩包目录
func (f *FileService) ZipSubmissionFolder(submissionPath string) (string, error) {
	if isDir, err := util.IsDirectory(submissionPath); !isDir {
		return "B1003", err
	}
	zippedPath := submissionPath + "/" + uuid.New().String() + ".zip"
	isZipped, err := util.ZipDictionary(zippedPath, submissionPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	if !isZipped {
		return "B1003", errors.New("can't find file")
	}
	return zippedPath, nil
}

// ClearSubmissionPath 清空提交目录
func (f *FileService) ClearSubmissionPath() (string, error) {
	submissionPath := f.JudgeEnvironmentConfiguration.JudgeEnvironment.SubmissionPath
	if err := util.ClearFileByFolderName(submissionPath); err != nil {
		return "B1007", err
	}
	return "00000", nil
}

// ClearSolutionPath 清空期望输入、输出目录
func (f *FileService) ClearSolutionPath() (string, error) {
	submissionPath := f.JudgeEnvironmentConfiguration.JudgeEnvironment.ResolutionPath
	if err := util.ClearFileByFolderName(submissionPath); err != nil {
		return "B1007", err
	}
	return "00000", nil
}
