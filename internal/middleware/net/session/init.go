package session

import (
	z "github.com/go-serv/service/internal"
)

func ServerInit(mwGroup z.NetMiddlewareGroupInterface) {
	mwGroup.AddRequestHandler(ServerSessionHandler)
}
