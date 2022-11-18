package vo

// JudgeHostConditionVO 描述判题服务器状态的视图层对象
type JudgeHostConditionVO struct {
	WorkPath             string
	ScriptPath           string
	ResolutionPath       string
	Port                 int
	WorkingAmount        int
	CpuCoreAmount        int
	MemoryCostPercentage int
	CpuCostPercentage    int
	QueueAmount          int
	MaxWorkingAmount     int
	Version              string
}
