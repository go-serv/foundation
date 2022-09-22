package middleware

import "github.com/go-serv/foundation/pkg/z"

func ServerInit(mw z.NetMiddlewareInterface) {
	mw.AddRequestHandler(serverReqHandler)
	mw.AddResponseHandler(serverResHandler)
}

func ClientInit(mw z.NetMiddlewareInterface) {
	mw.AddRequestHandler(clientReqHandler)
	mw.AddResponseHandler(clientResHandler)
}
