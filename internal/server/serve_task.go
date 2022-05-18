package server

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/service/internal/logger"
)

func serverStartedInfo(address string) {
	info := job.Logger(logger.Info)
	info("started to serve requests on %s", address)
}

//func (e *localEndpoint) ServeTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
//	init := func(task job.TaskInterface) {
//		err := e.Listen()
//		task.Assert(err)
//	}
//	run := func(task job.TaskInterface) {
//		e.serverStartedInfo(e.Address())
//		err := e.unixServe()
//		task.Assert(err)
//		task.Done()
//	}
//	fin := func(task job.TaskInterface) {
//	}
//	return init, run, fin
//}

func (e *tcp4Endpoint) ServeTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	init := func(task job.TaskInterface) {
		err := e.Listen()
		task.Assert(err)
	}
	run := func(task job.TaskInterface) {
		serverStartedInfo(e.Address())
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
		serverStartedInfo(e.Address())
		err := e.tcpServe()
		task.Assert(err)
		task.Done()
	}
	fin := func(task job.TaskInterface) {
	}
	return init, run, fin
}
