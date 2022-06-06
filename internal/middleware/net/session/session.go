package session

import "github.com/go-serv/service/pkg/z"

func ServerSessionHandler(next z.NetChainElementFn, req z.RequestInterface, res z.ResponseInterface) (err error) {
	_, err = next(req, res)
	return
}
