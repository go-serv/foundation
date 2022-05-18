//
//

package calldesc

import (
	"context"
	"github.com/go-serv/service/internal/grpc/meta"
)

type callDesc struct {
	context.Context
	RequestInterface
	ResponseInterface
}

type callDescServer struct {
	callDesc
}

type callDescClient struct {
	callDesc
}

type response struct {
	meta meta.MetaInterface
}

type request struct {
	meta meta.MetaInterface
}
