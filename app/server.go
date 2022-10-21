package app

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/internal/service"
	"github.com/go-serv/foundation/pkg/y/kv"
	"github.com/go-serv/foundation/pkg/z"
)

type server struct {
	mainJob    job.JobInterface
	middleware z.ServerMiddlewareInterface
	dashboard  z.DashboardInterface
	wp         z.WebProxyInterface
	kv         kv.KeyValueStorageInterface
}

func (srv *server) Job() job.JobInterface {
	return srv.mainJob
}

func (srv *server) Middleware() z.ServerMiddlewareInterface {
	return srv.middleware
}

func (srv *server) AddService(svc z.ServiceInterface) {
	rf := service.Reflection()
	rf.AddService(svc.Name())
	rf.Populate()
	runtime.Runtime().RegisterService(svc)
}

func (srv *server) Start() {
	services := runtime.Runtime().Services()
	if len(services) == 0 {
		panic("application has no services to run, use AddService method")
	}

	for _, svc := range services {
		svc.BindApp(srv)
		for _, ep := range svc.Endpoints() {
			ep.BindService(svc)
			srv.mainJob.AddTask(ep.ServeTask)
		}
	}
	<-srv.mainJob.Run()
}

func (srv *server) Stop(reason any) {
	srv.mainJob.Cancel(reason)
}

func (srv *server) WebProxy() z.WebProxyInterface {
	return srv.wp
}

func (srv *server) FetchConfig() {
	panic("")
}
