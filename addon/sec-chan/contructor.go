package sec_chan

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/addon/sec-chan/internal/client"
	sec_mw "github.com/go-serv/foundation/addon/sec-chan/internal/middleware"
	"github.com/go-serv/foundation/addon/sec-chan/internal/server"
	session_mw "github.com/go-serv/foundation/internal/middleware/net/session"
	"github.com/go-serv/foundation/pkg/z"
)

func NewSecChanService(app z.AppInterface, eps []z.EndpointInterface, cfg server.ConfigInterface) ServiceInterface {
	svc := server.NewSecureChanService(eps, cfg)
	// This will enable the service middleware for all registered services as application middleware chain
	// is being merged with the service ones during application start.
	sec_mw.ServerInit(app.Middleware())
	return svc
}

func NewSecChanClient(ep z.EndpointInterface, cfg ClientConfigInterface) (cc ClientInterface, jb job.JobInterface) {
	cc, jb = client.NewClient(ep)
	// NOTICE: clients have to specify the list and order of middlewares to use explicitly.
	session_mw.ClientInit(cc.Middleware())
	sec_mw.ClientInit(cc.Middleware())
	return
}
