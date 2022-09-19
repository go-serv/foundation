package app

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/internal/service"
	"github.com/go-serv/foundation/pkg/z"
)

type app struct {
	mainJob   job.JobInterface
	dashboard z.DashboardInterface
	wp        z.WebProxyInterface
}

func (a *app) Job() job.JobInterface {
	return a.mainJob
}

func (a *app) AddService(svc z.ServiceInterface) {
	rf := service.Reflection()
	rf.AddService(svc.Name())
	rf.Populate()
	runtime.Runtime().RegisterService(svc)
}

func (a *app) Start() {
	services := runtime.Runtime().Services()
	if len(services) == 0 {
		panic("application has no services to run, use AddService method")
	}
	for _, svc := range services {
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

func (a *app) WebProxy() z.WebProxyInterface {
	return a.wp
}

func (a *app) FetchConfig() {
	panic("")
}
