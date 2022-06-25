package session

import "github.com/go-serv/service/pkg/z"

func ServerInit(mw z.NetMiddlewareInterface) {
	mw.AddRequestHandler(serverReqHandler)
	mw.AddResponseHandler(serverResHandler)
}

func ClientInit(mw z.NetMiddlewareInterface) {
	mw.AddRequestHandler(clientReqHandler)
	mw.AddResponseHandler(clientResHandler)
}
