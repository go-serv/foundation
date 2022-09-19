package z

import (
	job "github.com/AgentCoop/go-work"
	"google.golang.org/grpc"
	"net/http"
)

type EventHandlerFn func(...any) bool

type AppInterface interface {
	Job() job.JobInterface
	FetchConfig()
	AddService(ServiceInterface)
	//Services() AppInterface
	Start()
	Stop(reason any)
	WebProxy() WebProxyInterface
}

type WebProxyInterface interface {
	BuildHttpServer(grpc *grpc.Server, ep NetEndpointInterface) (srv *http.Server, err error)
	Config() WebProxyConfigInterface
}

type WebProxyConfigInterface interface {
	Dashboard() DashboardInterface
}

type DashboardInterface interface {
	IsFeatureOn() bool
	UrlPath() string
	ServeHTTP(res http.ResponseWriter, req *http.Request)
}
