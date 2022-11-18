package bo

// JudgeHostConditionBO 描述判题服务器状态的业务对象
type JudgeHostConditionBO struct {
	WorkingAmount        int
	CpuCoreAmount        int
	MemoryCostPercentage int
	CpuCostPercentage    int
	QueueAmount          int
	MaxWorkingAmount     int
}

func NewJudgeHostConditionBO(workingAmount,
	cpuCoreAmount,
	memoryCostPercentage,
	cpuCostPercentage,
	queueAmount,
	maxWorkingAmount int) *JudgeHostConditionBO {

	return &JudgeHostConditionBO{
		WorkingAmount:        workingAmount,
		CpuCoreAmount:        cpuCoreAmount,
		MemoryCostPercentage: memoryCostPercentage,
		CpuCostPercentage:    cpuCostPercentage,
		QueueAmount:          queueAmount,
		MaxWorkingAmount:     maxWorkingAmount,
	}
}
