package service

import (
	"JudgeHost/src/config/configuration"
	"JudgeHost/src/config/global"
	"JudgeHost/src/models/bo"
	"JudgeHost/src/models/dto"
	"JudgeHost/src/util"
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"sync/atomic"
)

type JudgeService struct {
	JudgeEnvironmentConfiguration *configuration.JudgeEnvironmentConfiguration
	SolutionStdInPathKey          string
	SolutionExpectedStdOutPathKey string
	EnableJudgeCoreGuard          int
	DisableJudgeCoreGuard         int
	UseRootUid                    int
	UseDefaultUid                 int
	CompileOutMaxSize             int
}

func NewJudgeService(environmentConfiguration *configuration.JudgeEnvironmentConfiguration) *JudgeService {
	return &JudgeService{
		JudgeEnvironmentConfiguration: environmentConfiguration,
		SolutionStdInPathKey:          "stdIn",
		SolutionExpectedStdOutPathKey: "expectedStdOut",
		EnableJudgeCoreGuard:          1,
		DisableJudgeCoreGuard:         0,
		UseRootUid:                    0,
		UseDefaultUid:                 6666,
		CompileOutMaxSize:             100000,
	}
}

// CompileSubmission 读取 compile.sh 生成脚本
func (s *JudgeService) CompileSubmission() ([]string, error) {
	// 编译脚本
	compileScript := util.GetCompileScriptPath()

	// 本次提交工作目录
	submissionWorkingPath := util.GetSubmissionWorkingPath()

	// 判题核心脚本
	judgeCoreScript := util.GetJudgeCoreScriptPath()

	judgeDTO := util.GetJudgeConfig()

	//用户代码
	codePath := util.GetCodePath(global.Language[judgeDTO.Language])

	// 对应语言的编译脚本
	buildScript, err := global.GetBuildScriptByRunningPath(judgeDTO.Language, submissionWorkingPath, codePath)
	if err != nil {
		logrus.Debug("get build script running path error: ", err.Error())
		return nil, err
	}
	useId := s.UseDefaultUid
	if isJava := judgeDTO.Language == "JAVA"; isJava {
		useId = s.UseRootUid
	}

	compileCommand := exec.Command(compileScript,
		submissionWorkingPath,
		codePath,
		judgeDTO.SubmissionCode,
		buildScript,
		judgeCoreScript,
		strconv.FormatInt(int64(useId), 10),
		strconv.FormatInt(int64(s.CompileOutMaxSize), 10),
	)

	if err := compileCommand.Run(); err != nil {
		logrus.Debug("compileCommand.Run", err.Error())
		return nil, err
	}
	return s.ReadFile(submissionWorkingPath + "/" + util.CompileStdErrName)
	// [DEBUG]: TEST
	// return s.ReadFile("/home/cu1/test/submission" + "/" + util.CompileStdErrName)
}

// RunJudge 执行判题
func (s *JudgeService) RunJudge(judgeDTO *dto.JudgeDTO) ([]*dto.SingleJudgeResultDTO, error) {
	curId := atomic.AddInt64(&global.GlobalSubmissionId, 1)
	judgeConfigurationBO := bo.JudgeConfigurationBO{
		SubmissionId:   curId,
		JudgeConfig:    judgeDTO,
		SubmissionPath: s.JudgeEnvironmentConfiguration.JudgeEnvironment.SubmissionPath,
		WorkPath:       s.JudgeEnvironmentConfiguration.JudgeEnvironment.SubmissionPath,
		ScriptPath:     s.JudgeEnvironmentConfiguration.JudgeEnvironment.ScriptPath,
		ResolutionPath: s.JudgeEnvironmentConfiguration.JudgeEnvironment.ResolutionPath,
	}
	util.InitJudgeConfiguration(&judgeConfigurationBO)

	// 编译用户提交
	compileResult, err := s.CompileSubmission()
	if err != nil {
		logrus.Debug("compile submission error", err.Error())
		return nil, err
	}
	util.SetExtraInfo(compileResult)
	result := []*dto.SingleJudgeResultDTO{}
	if s.IsCompileSuccess(compileResult) {
		totalSolutions := judgeDTO.Solutions

		// DEBUG DATA
		// totalSolutions = append(judgeDTO.Solutions, &dto.SolutionDTO{})

		for index, solution := range totalSolutions {
			singleJudgeResult, err := s.RunForSingleJudge(solution, index+1)
			if err != nil {
				result = append(result, singleJudgeResult)
				return result, err
			}
			isAccept := singleJudgeResult.Condition == global.JudgeResult["ACCEPT"]
			result = append(result, singleJudgeResult)
			if !isAccept && judgeDTO.IsAcmMode() {
				break
			}
		}
	} else {
		resolution := dto.SingleJudgeResultDTO{}
		resolution.Condition = global.JudgeResult["COMPILE_ERROR"]
		resolution.SetMessage()
		result = append(result, &resolution)
	}
	return result, nil
}

func (s *JudgeService) IsCompileSuccess(compileResult []string) bool {
	language := util.GetJudgeConfig().Language
	// c语言家族（c && cpp）
	isCppFamily := language == "C" || language == "C_PLUS_PLUS"
	// java
	isJava := language == "JAVA"

	for _, str := range compileResult {
		isBad := strings.Contains(str, "error:") || strings.Contains(str, "错误:") || strings.Contains(str, "Error:")
		if isCppFamily && isBad {
			return false
		}
		if isJava && isBad {
			return false
		}
	}
	return true
}

func (s *JudgeService) ReadFile(filePath string) ([]string, error) {
	return util.ReadFileByLines(filePath)
}

func (s *JudgeService) RunForSingleJudge(solutionDTO *dto.SolutionDTO, index int) (*dto.SingleJudgeResultDTO, error) {
	singleJudgeRunningName := "running_" + strconv.FormatInt(int64(index), 10)
	input, output := s.GetResolutionInputAndOutputFile(solutionDTO)
	judging, err := s.StartJudging(input, singleJudgeRunningName)

	// fmt.Println("[DEBUG] service/judge.go:191 ", *judging)

	if err != nil {
		logrus.Debug("StartJudging error: ", err)
		return nil, err
	}
	judgeCoreStdErr, err := s.ReadFile(judging.StdErrPath)
	if err != nil {
		logrus.Debug("ReadFile error: ", err)
		return nil, err
	}

	if len(judgeCoreStdErr) == 0 {

		//fmt.Println("[DEBUG] service/judge.go:207 ", len(judgeCoreStdErr))

		isSuccess := judging.Condition == 1
		// [DEBUG]:  TEST
		// output = "/home/cu1/test/submission/exp.out"
		isPass, err := s.CompareOutputWithResolutions(judging.StdOutPath, output)

		// DEBUG
		if !isPass {
			judging.Condition = global.JudgeResult["WRONG_ANSWER"]
		}

		//logrus.Info("is pass ? ", isPass, "stdout:",
		// judging.StdOutPath, " output: ", output)
		if err != nil {
			logrus.Debug("CompareOutputWithResolutions error: ", err)
			return nil, err
		}
		if isSuccess && isPass {
			judging.Condition = global.JudgeResult["ACCEPT"]
		}
	} else {
		util.SetExtraInfo(judgeCoreStdErr)
		judging.Condition = global.JudgeResult["RUNTIME_ERROR"]
	}
	judging.SetMessage()
	return judging, nil
}

// GetResolutionInputAndOutputFile 获取输入文件和期望的输出文件，供后续判题使用
func (s *JudgeService) GetResolutionInputAndOutputFile(solution *dto.SolutionDTO) (string, string) {
	return solution.StdIn, solution.ExpectedStdOut
	// TODO 利用 redis 做本地缓存 文件在远程存储时选择实现
}

func (s *JudgeService) CompareOutputWithResolutions(submissionOutput, expectedOutput string) (bool, error) {
	compareScript := util.GetCompareScriptPath()

	// [DEBUG]: TEST
	// fmt.Println("[DEBUG] service/judge.go:247 ", compareScript)

	compareCommand := exec.Command(compareScript, submissionOutput, expectedOutput)
	var exitOut bytes.Buffer
	compareCommand.Stdout = &exitOut
	if err := compareCommand.Run(); err != nil {
		logrus.Debug("compareCommand.Wait error: ", err)
		return false, err
	}
	exitCode := exitOut.String()
	// fmt.Println("[DEBUG] service/judge.go:257 ", "0" == exitCode)
	// fmt.Println("[DEBUG] service/judge.go:258 ", exitCode)
	return strings.Contains(exitCode, "0"), nil
}

func (s *JudgeService) StartJudging(stdInPath, name string) (*dto.SingleJudgeResultDTO, error) {
	coreScript := util.GetJudgeCoreScriptPath()
	judgeConfig := util.GetJudgeConfig()
	workingPath := util.GetSubmissionWorkingPath()
	language := util.GetJudgeConfig().Language
	// c语言家族（c && cpp）
	isCppFamily := language == "C" || language == "C_PLUS_PLUS"

	isGuard := s.DisableJudgeCoreGuard
	if isCppFamily {
		isGuard = s.EnableJudgeCoreGuard
	}
	rootCommand := exec.Command("echo", "your root password")
	judgeCommand := exec.Command("sudo", "-S", coreScript,
		"-r", util.GetRunnerScriptPath(),
		"-o", workingPath+"/"+name+".out",
		"-t", strconv.FormatInt(int64(judgeConfig.RealTimeLimit), 10),
		"-c", strconv.FormatInt(int64(judgeConfig.CpuTimeLimit), 10),
		"-m", strconv.FormatInt(int64(judgeConfig.MemoryLimit), 10),
		"-f", strconv.FormatInt(int64(judgeConfig.OutputLimit), 10),
		"-e", workingPath+"/"+name+".err",
		"-i", stdInPath,
		"-g", strconv.FormatInt(int64(isGuard), 10),
	)

	judgeCommand.Stdin, _ = rootCommand.StdoutPipe()

	var stdout bytes.Buffer
	judgeCommand.Stdout = &stdout
	if err := judgeCommand.Start(); err != nil {
		logrus.Debug("cmd.Start error: ", err)
		return nil, err
	}

	if err := rootCommand.Run(); err != nil {
		logrus.Debug("root command run error: ", err.Error())
		return nil, err
	}

	if err := judgeCommand.Wait(); err != nil {
		logrus.Debug("cmd.Wait error: ", err)
		return nil, err
	}

	// [DEBUG]: TEST
	//judgeCommand = exec.Command("/home/cu1/test/judge_test.sh",
	//	"/home/cu1/test/submission")
	resultJson := stdout.String()
	//fmt.Println("[DEBUG] service/judge.go:280 ", resultJson)
	singleJudgeResultDTO := dto.SingleJudgeResultDTO{}
	if err := json.Unmarshal([]byte(resultJson), &singleJudgeResultDTO); err != nil {
		logrus.Debug("json.Unmarshal error: ", err)
		return nil, err
	}
	// fmt.Println("[DEBUG] service/judge.go:287 ", singleJudgeResultDTO)
	// fmt.Println("[DEBUG] service/judge.go:288 ", "back")
	return &singleJudgeResultDTO, nil
}

// ReadStdout 准备废弃
func (s *JudgeService) ReadStdout(stdout io.ReadCloser) ([]string, error) {
	reader := bufio.NewReader(stdout)
	result := []string{}
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		result = append(result, line)
	}
	return result, nil
}

func (s *JudgeService) WriteCodeToWorkingPath(code, extension string) (bool, error) {
	codePath := util.GetCodePath(extension)
	return util.WriteDataToFilePath(code, codePath)
}
