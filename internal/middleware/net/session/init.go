package session

import "github.com/go-serv/service/pkg/z"

func ServerInit(mwGroup z.NetMiddlewareInterface) {
	mwGroup.AddRequestHandler(serverSessionHandler)
}

func ClientInit(mwGroup z.NetMiddlewareInterface) {
	mwGroup.AddRequestHandler(clientSessionHandler)
}
