package net

import (
	"github.com/go-serv/service/internal/grpc/request"
	"github.com/go-serv/service/internal/service"
)

type NetworkServiceInterface interface {
	service.BaseServiceInterface
	Service_OnNewSession(req request.RequestInterface) error
	// Service_InfoNewSession returns timeout in seconds for a new session. Zero means no new session is required
	Service_InfoNewSession(methodName string) int32
	Service_InfoMsgEncryption(methodName string) bool
}
