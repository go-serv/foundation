package app

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/internal/service"
	"github.com/go-serv/foundation/pkg/z"
)

type (
	registryKey string
	registry    map[registryKey]interface{}
)

type app struct {
	platform z.PlatformInterface
	mainJob  job.JobInterface
	services []z.ServiceInterface
}

func (a *app) Platform() z.PlatformInterface {
	return a.platform
}

func (a *app) Job() job.JobInterface {
	return a.mainJob
}

func (a *app) AddService(svc z.ServiceInterface) {
	rf := service.Reflection()
	rf.AddService(svc)
	rf.Populate()
	runtime.Runtime().RegisterService(svc)
	a.services = append(a.services, svc)
}

func (a *app) Start() {
	if len(a.services) == 0 {
		panic("application has no services to run, use AddService method")
	}
	for _, svc := range a.services {
		svc.BindApp(a)
		for _, ep := range svc.Endpoints() {
			ep.BindService(svc)
			a.mainJob.AddTask(ep.ServeTask)
		}
	}
	<-a.mainJob.Run()
}

func (a *app) Stop(reason any) {
	a.mainJob.Cancel(reason)
}

func (a *app) FetchConfig() {
	panic("")
}
