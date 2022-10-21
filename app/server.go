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

func (srv *server) selectEndpoints() []z.EndpointInterface {
	var found bool
	results := make([]z.EndpointInterface, 0)
	for _, svc := range runtime.Runtime().Services() {
		for _, ep := range svc.Endpoints() {
			found = false
			for _, k := range results {
				if k == ep {
					found = true
					break
				}
			}
			if !found {
				results = append(results, ep)
			}
		}
	}
	return results
}

func (srv *server) Start() {
	services := runtime.Runtime().Services()
	if len(services) == 0 {
		panic("application has no services to run, use AddService method")
	}

	for _, svc := range services {
		svc.BindApp(srv)
	}

	eps := srv.selectEndpoints()
	for _, ep := range eps {
		ep.BindAppServer(srv)
		srv.mainJob.AddTask(ep.ServeTask)
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
