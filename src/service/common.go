package service

import (
	"JudgeHost/src/config/configuration"
	"JudgeHost/src/models/bo"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"runtime"
	"time"
)

type CommonService struct {
	JudgeEnvironmentConfiguration *configuration.JudgeEnvironmentConfiguration
	JudgeExecutorPool             *ants.PoolWithFunc
	Port                          int
	Version                       string
}

func NewCommonService(judgeEnvironmentConfiguration *configuration.JudgeEnvironmentConfiguration, judgeExecutorPool *ants.PoolWithFunc, port int, version string) *CommonService {
	return &CommonService{
		JudgeEnvironmentConfiguration: judgeEnvironmentConfiguration,
		JudgeExecutorPool:             judgeExecutorPool,
		Port:                          port,
		Version:                       version,
	}
}

func (c *CommonService) GetJudgeHostConfiguration() *bo.JudgeHostConfigurationBO {
	judgeHostConfiguration := &bo.JudgeHostConfigurationBO{
		ScriptPath:     c.JudgeEnvironmentConfiguration.JudgeEnvironment.ScriptPath,
		ResolutionPath: c.JudgeEnvironmentConfiguration.JudgeEnvironment.ResolutionPath,
		Port:           c.Port,
		Version:        c.Version,
	}
	return judgeHostConfiguration
}

func (c *CommonService) GetJudgeHostCondition() *bo.JudgeHostConditionBO {
	cpuCoreAmount := runtime.NumGoroutine()
	workingAmount := c.JudgeExecutorPool.Running()
	cpu := c.GetCpuCostPercentage()
	memory := c.GetMemoryCostPercentage()
	maxWorkingAmount := c.JudgeExecutorPool.Cap()
	return &bo.JudgeHostConditionBO{
		WorkingAmount:        workingAmount,
		CpuCoreAmount:        cpuCoreAmount,
		MemoryCostPercentage: memory,
		CpuCostPercentage:    cpu,
		MaxWorkingAmount:     maxWorkingAmount,
	}
}

// GetCpuCostPercentage 获取系统 cpu 占用率
func (c *CommonService) GetCpuCostPercentage() int {
	totalPercent, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		// TODO logger
		fmt.Println(err.Error())
		return -1
	}
	return int(totalPercent[0] * 100)
}

// GetMemoryCostPercentage 获取系统内存占用率
func (c *CommonService) GetMemoryCostPercentage() int {
	memory, err := mem.VirtualMemory()
	if err != nil {
		// TODO logger
		fmt.Println(err.Error())
		return -1
	}
	totalVirtualMemory := memory.Total
	freePhysicalMemorySize := memory.Free
	value := 1.0 * freePhysicalMemorySize / totalVirtualMemory
	return int((1 - value) * 100)
}

func (c *CommonService) SetJudgeHostWorkingAmount(corePoolSize int, isForceSet bool) {
	isHavingWorkingNode := (c.JudgeExecutorPool.Running() != 0)
	if !isForceSet && isHavingWorkingNode {
		// TODO logger
		fmt.Println("has task in working set fail")
		return
	}
	c.JudgeExecutorPool.Tune(corePoolSize)
	fmt.Println(c.JudgeExecutorPool.Cap())
}
