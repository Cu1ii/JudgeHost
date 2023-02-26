package judge

type Result struct {
	CpuTime  int `json:"cpu_time"`
	Signal   int `json:"signal"`
	Memory   int `json:"memory"`
	ExitCode int `json:"exit_code"`
	Result   int `json:"result"`
	Error    int `json:"error"`
	RealTime int `json:"real_time"`
}
