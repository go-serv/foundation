package webproxy

import (
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/service/net"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type webproxy struct {
	cfg         *WebProxyConfig
	wrappedGrpc *grpcweb.WrappedGrpcServer
}

func (wp *webproxy) Config() z.WebProxyConfigInterface {
	return wp.cfg
}

type WebProxyConfig struct {
	UseWebsocket          bool
	WsPingInterval        time.Duration
	WsReadLimit           int64
	PemCert               *net.X509PemPair
	AllowedOrigins        []string
	AllowedOriginsFn      func(string) bool
	AllowedRequestHeaders []string
	ApiUrlPath            string
	adminDashboard        z.DashboardInterface
}

func (cfg *WebProxyConfig) Dashboard() z.DashboardInterface {
	return cfg.adminDashboard
}

func (wp *webproxy) BuildHttpServer(grpc *grpc.Server, ep z.NetEndpointInterface) (srv *http.Server, err error) {
	options := make([]grpcweb.Option, 0)

	// Request origin
	if len(wp.cfg.AllowedOrigins) > 0 {
		options = append(options, grpcweb.WithOriginFunc(func(origin string) bool {
			for _, v := range wp.cfg.AllowedOrigins {
				// Check if a regular expression is being used to match request origin.
				if strings.HasPrefix(v, "^") || strings.HasSuffix(v, "$") {
					if ok := regexp.MustCompile(v).MatchString(origin); ok {
						return true
					}
				} else { // Direct match
					if strings.ToLower(v) == strings.ToLower(origin) {
						return true
					}
				}
			}
			return false
		}))
	} else if wp.cfg.AllowedOriginsFn != nil {
		options = append(options, grpcweb.WithOriginFunc(wp.cfg.AllowedOriginsFn))
	}

	wp.wrappedGrpc = grpcweb.WrapServer(grpc, options...)
	serveMux := http.NewServeMux()
	serveMux.Handle(wp.cfg.ApiUrlPath, wp.wrappedGrpc)
	srv = &http.Server{
		Addr:              ep.Address(),
		Handler:           serveMux,
		TLSConfig:         ep.TlsConfig(),
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	// Admin dashboard.
	dashboard := wp.cfg.Dashboard()
	if dashboard != nil {
		serveMux.Handle(dashboard.PathPrefix(), dashboard)
	}

	return
}
