package service

import (
	job "github.com/AgentCoop/go-work"
)

func (e *unixSocketEndpoint) ServeTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	init := func(task job.TaskInterface) {
	}
	run := func(task job.TaskInterface) {
		err := e.unixServe()
		task.Assert(err)
		task.Done()
	}
	fin := func(task job.TaskInterface) {
	}
	return init, run, fin
}

func (e *tcp4Endpoint) ServeTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	init := func(task job.TaskInterface) {
		err := e.Listen()
		task.Assert(err)
	}
	run := func(task job.TaskInterface) {
		err := e.tcpServe()
		task.Assert(err)
		task.Done()
	}
	fin := func(task job.TaskInterface) {
	}
	return init, run, fin
}

func (e *tcp6Endpoint) ServeTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	init := func(task job.TaskInterface) {
		err := e.Listen()
		task.Assert(err)
	}
	run := func(task job.TaskInterface) {
		task.Done()
	}
	fin := func(task job.TaskInterface) {
	}
	return init, run, fin
}
