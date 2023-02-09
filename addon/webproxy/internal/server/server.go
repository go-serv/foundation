package server

import (
	"github.com/mesh-master/foundation/pkg/z"
	"google.golang.org/grpc"
	"net/http"
	"strings"
)

type webproxy struct {
	httpMw   z.WebProxyMiddlewareInterface
	httpSrv  *http.Server
	httpMux  *http.ServeMux
	grpcSrv  *grpc.Server
	endpoint z.NetEndpointInterface
}

func (wp *webproxy) IsGrpcRequest(req *http.Request) bool {
	return req.ProtoMajor == 2 && strings.HasPrefix(req.Header.Get("Content-Type"), "application/grpc")
}

func (wp *webproxy) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if wp.IsGrpcRequest(req) {
		wp.grpcSrv.ServeHTTP(resp, req)
	} else {
		wp.httpMux.ServeHTTP(resp, req)
	}
}

func (wp *webproxy) Start() (err error) {
	if wp.endpoint.TlsConfig() == nil {
		panic("")
	}
	wp.httpMux = http.NewServeMux()
	wp.httpSrv = &http.Server{
		Addr:              wp.endpoint.Address(),
		Handler:           wp,
		TLSConfig:         wp.endpoint.TlsConfig(),
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
	err = wp.httpSrv.ListenAndServeTLS("", "")
	return
}

func (wp *webproxy) Stop() {

}
