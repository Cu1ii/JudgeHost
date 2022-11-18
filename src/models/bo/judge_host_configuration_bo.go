package bo

// JudgeHostConfigurationBO 判题服务器配置bo对象
type JudgeHostConfigurationBO struct {
	WorkPath       string
	ScriptPath     string
	ResolutionPath string
	Port           int
	Version        string
}
