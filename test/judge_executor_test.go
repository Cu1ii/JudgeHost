package test

import (
	"JudgeHost/src/config/configuration"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type SumTask struct {
	index int
	nums  []int
	sum   int
	l     int
	r     int
}

func (s *SumTask) Do() {
	for _, num := range s.nums {
		s.sum += num
	}
	time.Sleep(2 * time.Second)
	fmt.Println(s.index, " = : l = ", s.l, " r = ", s.r, " sum = ", s.sum)
}

func TestJudgeExcutorPool(t *testing.T) {

	configuration.SetJudgeNamePrefix("JudgeExecutorTest")

	const (
		DataSize    = 10000
		DataPerTask = 100
	)

	nums := make([]int, DataSize, DataSize)
	for i := range nums {
		nums[i] = rand.Intn(1000)
	}

	var wg sync.WaitGroup
	wg.Add(DataSize / DataPerTask)
	tasks := make([]*configuration.TaskWrop, 0, DataSize/DataPerTask)
	for i := 0; i < DataSize/DataPerTask; i++ {
		task := configuration.NewTaskWrop(
			&SumTask{index: i + 1,
				nums: nums[i*DataPerTask : (i+1)*DataPerTask],
				l:    i * DataPerTask,
				r:    (i + 1) * DataPerTask,
			}, &wg)
		tasks = append(tasks, task)
		err := configuration.JudgeExecutorPool.Invoke(task)
		if err != nil {
			fmt.Printf("task:%d err:%v\n", i, err)
			wg.Done()
		}
	}
	fmt.Println("available core number", configuration.CORE_POOL_SIZE)
	fmt.Printf("running goroutines: %d\n", configuration.JudgeExecutorPool.Running())
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())

	defer configuration.JudgeExecutorPool.Release()
}
