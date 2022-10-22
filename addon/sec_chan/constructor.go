package sec_chan

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/addon/sec_chan/internal/client"
	sec_mw "github.com/go-serv/foundation/addon/sec_chan/internal/middleware"
	"github.com/go-serv/foundation/addon/sec_chan/internal/server"
	"github.com/go-serv/foundation/addon/sec_chan/x"
	"github.com/go-serv/foundation/pkg/z"
)

func NewSecChanService(app z.AppServerInterface, eps []z.EndpointInterface, cfg server.ConfigInterface) ServiceInterface {
	svc := server.NewSecureChanService(eps, cfg)
	for _, ep := range eps {
		ep.AddService(svc)
	}
	app.Middleware().Insert(z.NetworkMwKey, z.InsertAfter, x.SecChanMwKey, sec_mw.ServerReqHandler, sec_mw.ServerResHandler)
	return svc
}

func NewSecChanClient(ep z.EndpointInterface, cfg ClientConfigInterface) (cc ClientInterface, jb job.JobInterface) {
	cc, jb = client.NewClient(ep)
	cc.Middleware().Insert(z.NetworkMwKey, z.InsertAfter, x.SecChanMwKey, sec_mw.ClientReqHandler, sec_mw.ClientResHandler)
	return
}
