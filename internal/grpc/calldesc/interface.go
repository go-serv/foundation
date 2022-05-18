package calldesc

import "github.com/go-serv/service/internal/grpc/meta"

type callDescInterface interface {
	Request() RequestInterface
	WithRequest(RequestInterface)
}

type ServerCallDescriptor interface {
	callDescInterface
}

type RequestInterface interface {
	Meta() meta.MetaInterface
	Validate() []error
}

type ResponseInterface interface {
	Meta() meta.MetaInterface
	//ToGrpc() interface{}
}
