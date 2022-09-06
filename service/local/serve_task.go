package local

import (
	job "github.com/AgentCoop/go-work"
)

func (ep *localEndpoint) ServeTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	init := func(task job.TaskInterface) {
	}
	run := func(task job.TaskInterface) {
		err := ""
		task.Assert(err)
		task.Done()
	}
	fin := func(task job.TaskInterface) {
	}
	return init, run, fin
}
