package judge

import (
	"JudgeHost/src/util"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func compile(submissionId int, code, submissionPath, language string) (bool, string) {
	switch language {
	case "C":
		return compileC(submissionId, code, submissionPath)
	case "C++":
		return compileCPP(submissionId, code, submissionPath)
	case "Java":
		return compileJava(submissionId, code, submissionPath)
	case "Python3":
		return compilePython3(submissionId, code, submissionPath)
	case "Python2":
		return compilePython2(submissionId, code, submissionPath)
	case "Go":
		return compileGo(submissionId, code, submissionPath)
	case "Swift5.1":
		return compileSwift(submissionId, code, submissionPath)
	default:
		return false, compileError("Unknown Language!")
	}
}

func compileC(id int, code, submissionPath string) (bool, string) {
	submissionPath = fmt.Sprintf("%s/%d", submissionPath, id)
	if _, err := os.Stat(submissionPath); os.IsNotExist(err) {
		if err := os.MkdirAll(submissionPath, os.ModePerm); err != nil {
			logrus.Error("create ", submissionPath, " fail ", err.Error())
		}
	}
	codePath := fmt.Sprintf("%s/Main.c", submissionPath)
	if _, err := util.WriteDataToFilePath(code, codePath); err != nil {
		logrus.Error(err.Error())
	}
	compileCmd := exec.Command("/home/cu1/XOJ/shell/compileC.sh", submissionPath, strconv.FormatInt(int64(id), 10))
	if err := compileCmd.Run(); err != nil {
		logrus.Error(err.Error())
		var msg string
		compileErrPath := fmt.Sprintf("%s/%dce.txt", submissionPath, id)
		lines, err := util.ReadFileByLines(compileErrPath)
		if err != nil {
			msg = "Fatal Compile error!"
			logrus.Error(err.Error())
			return false, compileError(msg)
		}
		if len(lines) == 0 {
			msg = "Compile timeout! Maybe you define too big arrays!"
		} else {
			msg = util.StringListToString(lines)
		}
		return false, compileError(msg)
	}
	return true, ""
}

func compileCPP(id int, code, submissionPath string) (bool, string) {
	submissionPath = fmt.Sprintf("%s/%d", submissionPath, id)
	if _, err := os.Stat(submissionPath); os.IsNotExist(err) {
		if err := os.MkdirAll(submissionPath, os.ModePerm); err != nil {
			logrus.Error("create ", submissionPath, " fail ", err.Error())
		}
	}
	codePath := fmt.Sprintf("%s/Main.cpp", submissionPath)
	if _, err := util.WriteDataToFilePath(code, codePath); err != nil {
		logrus.Error(err.Error())
	}
	compileCmd := exec.Command("/home/cu1/XOJ/shell/compileCPP.sh", submissionPath, strconv.FormatInt(int64(id), 10))
	if err := compileCmd.Run(); err != nil {
		logrus.Error(err.Error())
		var msg string
		compileErrPath := fmt.Sprintf("%s/%dce.txt", submissionPath, id)
		lines, err := util.ReadFileByLines(compileErrPath)
		if err != nil {
			msg = "Fatal Compile error!"
			logrus.Error(err.Error())
			return false, compileError(msg)
		}
		if len(lines) == 0 {
			msg = "Compile timeout! Maybe you define too big arrays!"
		} else {
			msg = util.StringListToString(lines)
		}
		return false, compileError(msg)
	}
	return true, ""
}

func compilePython2(id int, code, submissionPath string) (bool, string) {
	submissionPath = fmt.Sprintf("%s/%d", submissionPath, id)
	if filterWord := pythonFilters(code); filterWord != "0" {
		return false, compileError("Your code has sensitive words " + filterWord)
	}
	if _, err := os.Stat(submissionPath); os.IsNotExist(err) {
		if err := os.MkdirAll(submissionPath, os.ModePerm); err != nil {
			logrus.Error("create ", submissionPath, " fail ", err.Error())
			return false, compileError("System error")
		}
	}
	codePath := fmt.Sprintf("%s/Main.py", submissionPath)
	if _, err := util.WriteDataToFilePath(code, codePath); err != nil {
		logrus.Error(err.Error())
		return false, compileError("System error")
	}
	return true, ""
}

func compilePython3(id int, code, submissionPath string) (bool, string) {
	submissionPath = fmt.Sprintf("%s/%d", submissionPath, id)
	if filterWord := pythonFilters(code); filterWord != "0" {
		return false, compileError("Your code has sensitive words " + filterWord)
	}
	if _, err := os.Stat(submissionPath); os.IsNotExist(err) {
		if err := os.MkdirAll(submissionPath, os.ModePerm); err != nil {
			logrus.Error("create ", submissionPath, " fail ", err.Error())
			return false, compileError("System error")
		}
	}
	codePath := fmt.Sprintf("%s/Main.py", submissionPath)
	if _, err := util.WriteDataToFilePath(code, codePath); err != nil {
		logrus.Error(err.Error())
		return false, compileError("System error")
	}
	return true, ""
}

func compileGo(id int, code, submissionPath string) (bool, string) {
	submissionPath = fmt.Sprintf("%s/%d", submissionPath, id)
	if _, err := os.Stat(submissionPath); os.IsNotExist(err) {
		if err := os.MkdirAll(submissionPath, os.ModePerm); err != nil {
			logrus.Error("create ", submissionPath, " fail ", err.Error())
			return false, compileError("System error")
		}
	}
	codePath := fmt.Sprintf("%s/Main.go", submissionPath)
	if _, err := util.WriteDataToFilePath(code, codePath); err != nil {
		logrus.Error(err.Error())
		return false, compileError("System error")
	}
	compileCmd := exec.Command("/home/cu1/XOJ/shell/compileGo.sh", submissionPath, strconv.FormatInt(int64(id), 10))
	if err := compileCmd.Run(); err != nil {
		logrus.Error(err.Error())
		var msg string
		compileErrPath := fmt.Sprintf("%s/%dce.txt", submissionPath, id)
		lines, err := util.ReadFileByLines(compileErrPath)
		if err != nil {
			logrus.Error(err.Error())
			msg = "Fatal Compile error!"
			return false, compileError(msg)
		}
		if len(lines) == 0 {
			msg = "Compile timeout! Maybe you define too big arrays!"
		} else {
			msg = util.StringListToString(lines)
		}
		return false, compileError(msg)
	}
	return true, ""
}

func compileJava(id int, code, submissionPath string) (bool, string) {
	submissionPath = fmt.Sprintf("%s/%d", submissionPath, id)
	if _, err := os.Stat(submissionPath); os.IsNotExist(err) {
		if err := os.MkdirAll(submissionPath, os.ModePerm); err != nil {
			logrus.Error("create ", submissionPath, " fail ", err.Error())
			return false, compileError("System error")
		}
	}
	codePath := fmt.Sprintf("%s/Main.java", submissionPath)
	if _, err := util.WriteDataToFilePath(code, codePath); err != nil {
		logrus.Error(err.Error())
		return false, compileError("System error")
	}
	compileCmd := exec.Command("/home/cu1/XOJ/shell/compileJava.sh", submissionPath, strconv.FormatInt(int64(id), 10))
	if err := compileCmd.Run(); err != nil {
		logrus.Error(err.Error())
		var msg string
		compileErrPath := fmt.Sprintf("%s/%dce.txt", submissionPath, id)
		lines, err := util.ReadFileByLines(compileErrPath)
		if err != nil {
			logrus.Error(err.Error())
			msg = "Fatal Compile error!"
			return false, compileError(msg)
		}
		msg = util.StringListToString(lines)
		return false, compileError(msg)

	}
	return true, ""
}

// 还没有支持, 当前环境没有装 swift
func compileSwift(id int, code, submissionPath string) (bool, string) {
	submissionPath = fmt.Sprintf("%s/%d", submissionPath, id)
	if _, err := os.Stat(submissionPath); os.IsNotExist(err) {
		if err := os.MkdirAll(submissionPath, os.ModePerm); err != nil {
			logrus.Error("create ", submissionPath, " fail ", err.Error())
			return false, compileError("System error")
		}
	}
	codePath := fmt.Sprintf("%s/Main.swift", submissionPath)
	if _, err := util.WriteDataToFilePath(code, codePath); err != nil {
		logrus.Error(err.Error())
		return false, compileError("System error")
	}
	compileCmd := exec.Command("/home/cu1/XOJ/shell/compileJava.sh", submissionPath, strconv.FormatInt(int64(id), 10))
	if err := compileCmd.Run(); err != nil {
		logrus.Error(err.Error())
		var msg string
		compileErrPath := fmt.Sprintf("%s/%dce.txt", submissionPath, id)
		lines, err := util.ReadFileByLines(compileErrPath)
		if err != nil {
			logrus.Error(err.Error())
			msg = "Fatal Compile error!"
			return false, compileError(msg)
		}
		if len(lines) == 0 {
			msg = "Compile timeout! Maybe you define too big arrays!"
		} else {
			msg = util.StringListToString(lines)
		}
		return false, compileError(msg)

	}
	return true, ""
}

func pythonFilters(code string) string {
	if strings.Contains(code, "thread") {
		return "thread"
	}
	if strings.Contains(code, "process") {
		return "process"
	}
	if strings.Contains(code, "resource") {
		return "resource"
	}
	if strings.Contains(code, "ctypes") {
		return "ctypes"
	}
	if strings.Contains(code, "os") {
		return "os"
	}
	if strings.Contains(code, "__import__") {
		return "__import__"
	}
	if strings.Contains(code, "eval") {
		return "eval"
	}
	if strings.Contains(code, "exec") {
		return "exec"
	}
	if strings.Contains(code, "globals") {
		return "globals"
	}
	if strings.Contains(code, "locals") {
		return "locals"
	}
	if strings.Contains(code, "compile") {
		return "compile"
	}
	if strings.Contains(code, "frame") {
		return "frame"
	}
	return "0"
}
