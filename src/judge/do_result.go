package judge

import (
	"JudgeHost/src/util"
)

func compileError(msg string) *JudgeResponse {
	res := JudgeResponse{
		Result: -4,
		Msg:    msg,
	}
	return &res
}

func doneProblem(
	message string,
	memory,
	cpuTime int,
	result int,
	testcase string,
	responseVo *JudgeResponse) *JudgeResponse {
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
		responseVo.AppendCase(&CaseStatus{TestCase: testcase})
	}
	return responseVo
}

func acProblem(memory, cpuTime int, responseVo *JudgeResponse) *JudgeResponse {
	responseVo.Msg = "Accept"
	responseVo.Result = 0
	responseVo.Memory = memory
	responseVo.CpuTime = cpuTime
	return responseVo
}

func doneCase(result string,
	time,
	memory int,
	testcase,
	caseData,
	outputData,
	userOutput string,
	rep *JudgeResponse) {
	rep.AppendCase(&CaseStatus{
		Result:     result,
		Time:       time,
		Memory:     memory,
		TestCase:   testcase,
		CaseData:   caseData,
		OutputData: outputData,
		UserOutput: userOutput,
	})
}
