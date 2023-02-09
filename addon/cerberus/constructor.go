package cerberus

import (
	mw "github.com/mesh-master/foundation/addon/cerberus/internal/middleware"
	"github.com/mesh-master/foundation/addon/cerberus/internal/server"
	"github.com/mesh-master/foundation/addon/cerberus/x"
	"github.com/mesh-master/foundation/pkg/z"
)

func NewCerberusService(app z.AppServerInterface, eps []z.EndpointInterface, cfg server.ConfigInterface) ServiceInterface {
	svc := server.NewCerberusService(eps, cfg)
	app.Middleware().Insert(z.NetworkMwKey, z.InsertAfter, x.CerberusMwKey, mw.ServerReqHandler, nil)
	for _, ep := range eps {
		ep.AddService(svc)
	}
	return svc
}

//func NewSecChanClient(ep z.EndpointInterface, cfg ClientConfigInterface) (cc ClientInterface, jb job.JobInterface) {
//	cc, jb = client.NewClient(ep)
//	//cc.Middleware().Insert(z.SessionMwKey, z.InsertAfter, x.SecChanMwKey, sec_mw.ClientReqHandler, sec_mw.ClientResHandler)
//	return
//}
