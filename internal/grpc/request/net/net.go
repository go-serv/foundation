package net

import "github.com/go-serv/service/internal/grpc/request"

type clientInfo struct {
}

type netRequest struct {
	request.Request
	clientInfo *clientInfo
}
