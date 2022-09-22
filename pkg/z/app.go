package z

import (
	job "github.com/AgentCoop/go-work"
	"google.golang.org/grpc"
	"net/http"
)

type EventHandlerFn func(...any) bool

type AppInterface interface {
	Job() job.JobInterface
	Middleware() MiddlewareInterface
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
	// PathPrefix returns current URL path prefix of the admin dashboard. Defaults to /dashboard/.
	PathPrefix() string

	// WithPathPrefix sets the URL path prefix for the admin dashboard.
	WithPathPrefix(string)

	ServeHTTP(res http.ResponseWriter, req *http.Request)
}
