package meta

import (
	z "github.com/go-serv/service/internal"
	ancillary_net "github.com/go-serv/service/internal/ancillary/net"
	"net"
)

type CommonDictionary struct {
	z.MetaDictionary
	SessionId z.SessionId `header:"gserv-session-id"`
}

type ServiceDictionary struct {
	CommonDictionary
	// The Content-Type representation header is used to indicate the original media type of the resource
	// (prior to any content encoding applied for sending).
	//

	// The X-Forwarded-For (XFF) request header is a de-facto standard header for identifying the originating IP address
	// of a client connecting to a web server through a proxy server.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For
	ClientIp net.IPAddr `name:"x-forwarded-for"`

	// The X-Forwarded-Host (XFH) header is a de-facto standard header for identifying the original host requested by the
	// client in the Host HTTP request header.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-Host
	ForwardedHost ancillary_net.Host `name:"x-forwarded-host"`

	// The X-Forwarded-Proto (XFP) header is a de-facto standard header for identifying the protocol (HTTP or HTTPS)
	// that a client used to connect to your proxy or load balancer.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-Proto
	ForwardedProto string `name:"x-forwarded-proto"`
}

type ClientDictionary struct {
	CommonDictionary
}
