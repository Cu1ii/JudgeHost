package judge

import (
	"JudgeHost/src/global"
	"JudgeHost/src/models/vo"
	"JudgeHost/src/util"
	"errors"
	"fmt"
	"github.com/jiang4869/go-judger"
	"github.com/sirupsen/logrus"
	"io"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const (
	Python2Path      = "/usr/bin/python2"
	Python3Path      = "/usr/bin/python3"
	JavaPath         = "/usr/bin/java"
	SpecialJudgePath = "./scripts/spj.sh"
)

func judge(submitId,
	memoryLimit,
	timeLimit int,
	resolutionPath,
	code,
	language string,
	problem int,
	isOI int,
	spj bool,
	res *string) error {

	rep := vo.ResponseVo{}
	logrus.Info("Begin to Compile!")
	if success, r := compile(submitId, code, global.JudgeEnvironmentSetting.SubmissionPath, language); !success {
		logrus.Debug("compile ", problem, " fail")
		*res = r
		return errors.New("compile" + strconv.FormatInt(int64(problem), 10) + "fail")
	}
	// submitTime := submittime.Unix() // 似乎并没有什么用
	//runPath := fmt.Sprintf("%s/%d/%d.out", XojSubmissionPath, id, id)
	maxMemory := 0
	maxTime := 0

	myResult := 100
	myTestcase := ""
	myTime := 0
	myMemory := 0
	if resolutionPath != "" {
		// 此时去远程下载测试用例
	} else {
		// 若为空则默认采用本地的测试用例进行测试
		resolutionPath = fmt.Sprintf("%s/%d", global.JudgeEnvironmentSetting.ResolutionPath, problem)
	}
	files, err := filepath.Glob(resolutionPath + "/*")
	if err != nil || len(files) == 0 {
		*res = doneProblem("get resolution error!",
			0,
			0,
			5,
			"?",
			&rep)
		logrus.Error(err)
		return err
	}
	zipPath := fmt.Sprintf("%s/%d.zip", resolutionPath, problem)
	// resolutionPath + "/" + problem + ".zip"
	if util.IsFileIn(zipPath) {
		if isSuccess, err := util.UnZipInDictionary(zipPath, resolutionPath); !isSuccess || err != nil {
			*res = doneProblem("unzip "+strconv.FormatInt(int64(problem), 10)+".zip error!",
				0,
				0,
				5,
				"?",
				&rep)
			return err
		}
		if isSuccess, err := util.DeleteFile(zipPath, true); !isSuccess || err != nil {
			*res = doneProblem("delete resolution zip error!",
				0,
				0,
				5,
				"?",
				&rep)
			return err
		}
	}
	logrus.Info("resolution path = ", resolutionPath)
	logrus.Info("len files = ", len(files))
	logrus.Info("files = ", files)

	inputFiles := []string{}
	exoutputFiles := []string{}

	// 这里通过 Contains 的方法判断文件扩展名是不严谨的
	for _, file := range files {
		if strings.Contains(file, ".in") {
			inputFiles = append(inputFiles, file)
		} else if strings.Contains(file, ".out") {
			exoutputFiles = append(exoutputFiles, file)
		}
	}

	sort.Strings(inputFiles)
	sort.Strings(exoutputFiles)

	errorPath := fmt.Sprintf("%s/%d/%d%s", global.JudgeEnvironmentSetting.SubmissionPath, submitId, submitId, "error.out")
	logPath := fmt.Sprintf("%s/%d/%d%s", global.JudgeEnvironmentSetting.SubmissionPath, submitId, submitId, "log.out")

	for idx, in := range inputFiles {
		logrus.Infof("Judging!! %s/%d.in", resolutionPath, idx)

		outputPath := fmt.Sprintf("%s/%d/%d%s", global.JudgeEnvironmentSetting.SubmissionPath, submitId, idx, "temp.out")
		inputPath := fmt.Sprintf("%s/%d.in", resolutionPath, idx)
		result, err := singleJudge(
			timeLimit,
			memoryLimit,
			submitId,
			inputPath,
			outputPath,
			errorPath,
			logPath,
			language,
		)
		if err != nil {
			logrus.Errorf("Error occurred while calling judger: %v", err)
			result.Result = 5
		}
		maxInt := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}
		maxMemory = maxInt(maxMemory, result.Memory)
		maxTime = maxInt(maxTime, result.CpuTime)
		expectOutputPath := fmt.Sprintf("%s/%d.out", resolutionPath, idx)
		userOutputData := ""
		caseData := ""
		outputData := ""
		caseData, err = util.ReadFileByByte(inputPath, 300)
		if err != nil && err != io.EOF {
			result.Result = 5
		}
		outputData, err = util.ReadFileByByte(expectOutputPath, 300)
		if err != nil && err != io.EOF {
			result.Result = 5
		}
		userOutputData, err = util.ReadFileByByte(outputPath, 300)
		if err != nil && err != io.EOF {
			result.Result = 5
		}
		if result.Result != 0 {
			// 据 LOJ 源码中反应, QDU 判 Memory Exceed 有谜之 Bug, 这里还没有确认, 所以先使用 LPOJ 处理方式
			if (result.Result == 4 && result.ExitCode == 127 && result.Signal == 0) ||
				(result.Result == 4 && result.ExitCode == 0 && result.Signal == 31) {
				if myResult == 100 {
					myResult = 3
					myTestcase = in[strings.LastIndex(in, "/")+1:]
					myTime = result.CpuTime
					myMemory = result.Memory
				}
				doneCase(
					"Memory Limit Exceeded",
					result.CpuTime,
					result.Memory/1024/1024,
					in,
					caseData,
					outputData,
					userOutputData,
					&rep,
				)
			} else {
				if myResult == 100 {
					myResult = result.Result
					myTestcase = in[strings.LastIndex(in, "/")+1:]
					myTime = result.CpuTime
					myMemory = result.Memory
				}
				resultStr := util.TransformResultToString(result.Result)
				//"unknown"
				//if result.Result == 2 || result.Result == 1 {
				//	resultStr = "Time Limit Exceeded"
				//}
				//if result.Result == 3 {
				//	resultStr = "Memory Limit Exceeded"
				//}
				//if result.Result == 4 {
				//	resultStr = "Runtime Error"
				//}
				//if result.Result == 5 {
				//	resultStr = "System Error"
				//}
				doneCase(
					resultStr,
					result.CpuTime,
					result.Memory/1024/1024,
					in[strings.LastIndex(in, "/")+1:],
					caseData,
					outputData,
					userOutputData,
					&rep,
				)
			}
		} else {
			// isSpj := ""
			res := 0 // 0 ac -3 wrong -5 presentation
			spjPath := fmt.Sprintf("%s/checker", resolutionPath)
			// 若存在 checker 文件, 则说明为 spj 问题
			if util.IsFileIn(spjPath) {
				logrus.Info("Begin to special judge!")
				// isSpj = " (This test case is Special Judge) "
				res = specialjudge(spjPath, inputPath, outputPath, expectOutputPath)
			} else {
				logrus.Info("Comparing output!")
				std, stdErr := util.ReadFileByLines(outputPath)
				ans, ansErr := util.ReadFileByLines(expectOutputPath)
				if stdErr != nil {
					logrus.Error(err.Error())
					res = -3
				} else if ansErr != nil {
					logrus.Error(err.Error())
					res = -3
				} else {
					if len(std) != len(ans) {
						res = -3
					} else {
						isCorrect := true
						for i, stdLine := range std {
							ansLine := ans[i]
							// 先判断在不除去末尾 "\n \t\r" 这些符号时是否相同, 这时有可能出现末尾 "\n \t\r" 不同导致的对比失败
							if stdLine != ansLine {
								res = -3
							}
							// 再判断 "\n \t\r" 除去后是否相同, 此时若是不相同则表示结果错误
							stdLine = strings.TrimRight(stdLine, "\n \t\r")
							ansLine = strings.TrimRight(ansLine, "\n \t\r")
							if stdLine != ansLine {
								res = -3
								isCorrect = false
							}
						}
						// 如果最后未出现非空白符错误 但同时 result == -3 则表示格式错误
						if isCorrect && res == -3 {
							res = -5
						}
					}
				}
			}
			// 单组样例评判完毕，保存单组样例评判结果
			if res != 0 {
				if myResult == 100 {
					myResult = res
					myTestcase = in[strings.LastIndex(in, "/")+1:]
					myTime = result.CpuTime
					myMemory = result.Memory
				}
				resultStr := util.TransformResultToString(res)
				//"Unknown"
				//if res == -5 {
				//	resultStr = "Presentation Error"
				//}
				//if res == -3 {
				//	resultStr = "Wrong Answer"
				//}
				//if res == 5 {
				//	resultStr = "System Error"
				//}
				doneCase(
					resultStr,
					result.CpuTime,
					result.Memory/1024/1024,
					in[strings.LastIndex(in, "/")+1:],
					caseData,
					outputData,
					userOutputData,
					&rep,
				)
				if isOI == 0 {
					if result.Result == 5 || result.Result == 4 {
						rep.Error = result.Error
					}
					break
				}
			} else {
				doneCase(
					"Accepted",
					result.CpuTime,
					result.Memory/1024/1024,
					in[strings.LastIndex(in, "/")+1:],
					caseData,
					outputData,
					userOutputData,
					&rep,
				)
			}
		}
		logrus.Info("Done one data!")
	}
	// 汇总所有结果
	if myResult == 100 {
		*res = acProblem(maxMemory/1024/1024, maxTime, &rep)
	} else {
		*res = doneProblem("",
			myMemory/1024/1024,
			myTime,
			myResult,
			"",
			&rep)
	}
	// 历史遗留问题
	fmt.Sprintf("%s", myTestcase)

	logrus.Info("All done!")
	return nil
}

func singleJudge(timeLimit,
	memoryLimit,
	id int,
	inputPath,
	outputPath,
	errorPath,
	logPath,
	language string) (*judger.Result, error) {
	switch language {
	case "C":
		return judgeC(timeLimit, memoryLimit, id, inputPath, outputPath, errorPath, logPath)
	case "C++":
		return judgeCPP(timeLimit, memoryLimit, id, inputPath, outputPath, errorPath, logPath)
	case "Java":
		return judgeJava(timeLimit, memoryLimit, id, inputPath, outputPath, errorPath, logPath)
	case "Python3":
		return judgePyhton3(timeLimit, memoryLimit, id, inputPath, outputPath, errorPath, logPath)
	case "Python2":
		return judgePyhton2(timeLimit, memoryLimit, id, inputPath, outputPath, errorPath, logPath)
	case "Go":
		return judgeGo(timeLimit, memoryLimit, id, inputPath, outputPath, errorPath, logPath)
	default:
		return nil, errors.New("unknown language")
	}
}

func judgeC(timeLimit,
	memoryLimit,
	id int,
	inputPath,
	outputPath,
	errorPath,
	logPath string) (*judger.Result, error) {
	execPath := fmt.Sprintf("%s/%d/%d.o", global.JudgeEnvironmentSetting.SubmissionPath, id, id)
	return judger.Run(
		timeLimit,
		timeLimit*10,
		memoryLimit*1024*1024,
		32*1024*1024,
		32*1024*1024,
		10,
		0,
		0,
		0,
		[]string{},
		[]string{},
		execPath,
		inputPath,
		outputPath,
		errorPath,
		logPath,
		"c_cpp",
		//"c_cpp", 不知道为什么沙箱权限打开会有问题
	)
}

func judgeCPP(timeLimit,
	memoryLimit,
	id int,
	inputPath,
	outputPath,
	errorPath,
	logPath string) (*judger.Result, error) {
	execPath := fmt.Sprintf("%s/%d/%d.o", global.JudgeEnvironmentSetting.SubmissionPath, id, id)
	return judger.Run(
		timeLimit,
		timeLimit*10,
		memoryLimit*1024*1024,
		32*1024*1024,
		32*1024*1024,
		10,
		0,
		0,
		0,
		[]string{},
		[]string{},
		execPath,
		inputPath,
		outputPath,
		errorPath,
		logPath,
		"c_cpp",
		//"c_cpp", 不知道为什么沙箱权限打开会有问题
	)
}

func judgeGo(timeLimit,
	memoryLimit,
	id int,
	inputPath,
	outputPath,
	errorPath,
	logPath string) (*judger.Result, error) {
	execPath := fmt.Sprintf("%s/%d/%d.o", global.JudgeEnvironmentSetting.SubmissionPath, id, id)
	return judger.Run(
		timeLimit,
		timeLimit*10,
		memoryLimit*1024*1024,
		32*1024*1024,
		32*1024*1024,
		10,
		0,
		0,
		0,
		[]string{},
		[]string{},
		execPath,
		inputPath,
		outputPath,
		errorPath,
		logPath,
		"golang",
	)
}

func judgeJava(timeLimit,
	memoryLimit,
	id int,
	inputPath,
	outputPath,
	errorPath,
	logPath string) (*judger.Result, error) {
	runPath := fmt.Sprintf("%s/%d/%d", global.JudgeEnvironmentSetting.SubmissionPath, id, id)
	// javaArgs := fmt.Sprintf("'%s %s %s %s %s'", "-cp", runPath, "-Djava.security.policy==policy", "-Djava.awt.headless=true", "Main")
	return judger.Run(
		timeLimit,
		timeLimit*10,
		memoryLimit*1024*1024,
		32*1024*1024,
		32*1024*1024,
		10,
		0,
		0,
		1,
		[]string{"-cp", runPath, "-Djava.security.policy==policy", "-Djava.awt.headless=true", "Main"},
		[]string{},
		JavaPath,
		inputPath,
		outputPath,
		errorPath,
		logPath,
		"general", // general
	)
}

func judgePyhton2(timeLimit,
	memoryLimit,
	id int,
	inputPath,
	outputPath,
	errorPath,
	logPath string) (*judger.Result, error) {
	runPath := fmt.Sprintf("%s/%d/Main.py", global.JudgeEnvironmentSetting.SubmissionPath, id)
	return judger.Run(
		timeLimit,
		timeLimit*10,
		memoryLimit*1024*1024,
		32*1024*1024,
		32*1024*1024,
		10,
		0,
		0,
		1,
		[]string{runPath},
		[]string{},
		Python2Path,
		inputPath,
		outputPath,
		errorPath,
		logPath,
		"general", // general
	)
}

func judgePyhton3(timeLimit,
	memoryLimit,
	id int,
	inputPath,
	outputPath,
	errorPath,
	logPath string) (*judger.Result, error) {
	runPath := fmt.Sprintf("%s/%d/Main.py", global.JudgeEnvironmentSetting.SubmissionPath, id)
	return judger.Run(
		timeLimit,
		timeLimit*10,
		memoryLimit*1024*1024,
		32*1024*1024,
		32*1024*1024,
		10,
		0,
		0,
		1,
		[]string{runPath},
		[]string{},
		Python3Path,
		inputPath,
		outputPath,
		errorPath,
		logPath,
		"general", // general
	)
}

func specialjudge(checkerPath, inputPath, outputPath, expectOutputPath string) int {
	spjCmd := fmt.Sprintf("%s %s %s %s %s ", SpecialJudgePath, checkerPath, inputPath, outputPath, expectOutputPath)
	res, err := judger.ExecShell(spjCmd)
	if err != nil {
		return 5
	}
	// 传来的 res 数组里有多余的空白符不知道为什么...
	r := strings.TrimRight(string(res), "\t \n")
	if r == "1" {
		return -3
	} else if r == "0" {
		return 0
	}
	return 5
}
