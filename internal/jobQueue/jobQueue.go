package jobQueue

import (
	"fmt"
	"github.com/redpkg/formula/log"
	"time"
)

var job chan int
var i int

func init() {
	i = 1
	job = make(chan int, 2)
}

func AddJobTask() {
	log.Info().Msgf("add task, i:%v", i)
	// Add Task to jqb queue worker
	job <- 1 * i
	job <- 2 * i
	job <- 3 * i

	i++
}

func Run() {
	log.Info().Msg("run jub queue worker")
	go worker(job)
}

func worker(jobChan <-chan int) {
	for job := range jobChan {
		fmt.Println("jobChan len:%v", len(jobChan))
		fmt.Println("current job:", job)
		time.Sleep(3 * time.Second)
		fmt.Println("finished job:", job)
	}
}
