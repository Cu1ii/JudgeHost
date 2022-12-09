package judger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

type Result struct {
	CpuTime  int `json:"cpu_time"`
	Signal   int `json:"signal"`
	Memory   int `json:"memory"`
	ExitCode int `json:"exit_code"`
	Result   int `json:"result"`
	Error    int `json:"error"`
	RealTime int `json:"real_time"`
}

func Run(max_cpu_time,
	max_real_time,
	max_memory,
	max_stack,
	max_output_size,
	max_process_number,
	uid,
	gid,
	memory_limit_check_only int,
	args,
	env []string,
	exe_path,
	input_path,
	output_path,
	error_path,
	log_path,
	seccomp_rule_name string) (*Result, error) {

	proc_args := []string{"echo \"370802wsl\" | sudo -S /usr/lib/judger/libjudger.so"}

	//proc_args = append(proc_args, args...)
	if len(args) > 0 {
		for _, arg := range args {
			proc_args = append(proc_args, "--args="+arg)
		}
	}
	proc_args = append(proc_args, env...)
	proc_args = append(proc_args, fmt.Sprintf("--max_cpu_time=%d", max_cpu_time))
	proc_args = append(proc_args, fmt.Sprintf("--max_real_time=%d", max_real_time))
	proc_args = append(proc_args, fmt.Sprintf("--max_memory=%d", max_memory))
	proc_args = append(proc_args, fmt.Sprintf("--max_stack=%d", max_stack))
	proc_args = append(proc_args, fmt.Sprintf("--max_output_size=%d", max_output_size))
	proc_args = append(proc_args, fmt.Sprintf("--max_process_number=%d", max_process_number))
	proc_args = append(proc_args, fmt.Sprintf("--exe_path=%s", exe_path))
	proc_args = append(proc_args, fmt.Sprintf("--input_path=%s", input_path))
	proc_args = append(proc_args, fmt.Sprintf("--output_path=%s", output_path))
	proc_args = append(proc_args, fmt.Sprintf("--error_path=%s", error_path))

	proc_args = append(proc_args, fmt.Sprintf("--log_path=%s", log_path))
	proc_args = append(proc_args, fmt.Sprintf("--uid=%d", uid))
	proc_args = append(proc_args, fmt.Sprintf("--gid=%d", gid))
	proc_args = append(proc_args, fmt.Sprintf("--memory_limit_check_only=%d", memory_limit_check_only))

	if len(seccomp_rule_name) > 0 {
		proc_args = append(proc_args, fmt.Sprintf("--seccomp_rule_name=%s", seccomp_rule_name))
	}
	s := strings.Join(proc_args, " ")
	data, err := ExecShell(s)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(data))
	fmt.Println(s)
	var result Result
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func ExecShell(s_cmd string) ([]byte, error) {

	cmd := exec.Command("/bin/bash", "-c", s_cmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	cmd.Start()

	data, err := ioutil.ReadAll(stdout)
	if err != nil {
		return nil, err
	}
	return data, nil
}
