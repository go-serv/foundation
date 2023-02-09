package event

import (
	"github.com/mesh-master/foundation/pkg/z"
	"net/http"
)

type keyType int

const (
	NetClientNewContext keyType = iota + 1
	// NetClientBeforeDial event is being fired before a client connection will be established.
	NetClientBeforeDial
	// WebProxyNewHttpServer event is being fired when an HTTP server is created by the web proxy.
	WebProxyNewHttpServer
)

type NetClientBeforeDialArgs struct {
	Client     z.NetworkClientInterface
	TlsEnabled bool
}

type WebProxyNewHttpServerArgs struct {
	Server *http.Server
	Mux    *http.ServeMux
}
