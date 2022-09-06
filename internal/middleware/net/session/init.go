package session

import "github.com/go-serv/foundation/pkg/z"

func ServerInit(mw z.NetMiddlewareInterface) {
	mw.AddRequestHandler(serverSessionHandler)
}

func ClientInit(mw z.NetMiddlewareInterface) {
	mw.AddRequestHandler(clientSessionHandler)
}
