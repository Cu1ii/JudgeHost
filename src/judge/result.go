package judge

import (
	"JudgeHost/src/models/entity"
	"JudgeHost/src/models/vo"
	"JudgeHost/src/util"
)

func compileError(msg string) string {
	res := vo.ResponseVo{
		Result: -4,
		Msg:    msg,
	}
	toJson, _ := util.TransToJson(res)
	return toJson
}

func doneProblem(
	message string,
	memory,
	cpuTime int,
	result int,
	testcase string,
	responseVo *vo.ResponseVo) string {
	// fmt.Println(id, " ", problem, " ", message, " ", mytime, " ", username, " ", contest, " ", result, " ", testcase)

	responseVo.CpuTime = cpuTime
	responseVo.Memory = memory
	responseVo.Result = result

	if message == "" {
		responseVo.Msg = util.TransformResultToString(responseVo.Result)
	} else {
		responseVo.Msg = message
	}
	if testcase != "" {
		responseVo.AppendCase(&entity.CaseStatus{TestCase: testcase})
	}
	toJson, _ := util.TransToJson(responseVo)
	return toJson
}

func acProblem(memory, cpuTime int, responseVo *vo.ResponseVo) string {
	responseVo.Msg = "Accept"
	responseVo.Result = 0
	responseVo.Memory = memory
	responseVo.CpuTime = cpuTime
	toJson, _ := util.TransToJson(responseVo)
	return toJson
}

func doneCase(result string,
	time,
	memory int,
	testcase,
	caseData,
	outputData,
	userOutput string,
	rep *vo.ResponseVo) {
	rep.AppendCase(&entity.CaseStatus{
		Result:     result,
		Time:       time,
		Memory:     memory,
		TestCase:   testcase,
		CaseData:   caseData,
		OutputData: outputData,
		UserOutput: userOutput,
	})
}
