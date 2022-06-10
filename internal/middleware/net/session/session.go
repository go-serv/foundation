package session

import "github.com/go-serv/service/pkg/z"

func serverSessionHandler(next z.NetChainElementFn, req z.RequestInterface, res z.ResponseInterface) (err error) {
	_, err = next(req, res)
	if err != nil {
		return
	}
	return
}

func clientSessionHandler(next z.NetChainElementFn, req z.RequestInterface, res z.ResponseInterface) (err error) {
	_, err = next(req, res)
	if err != nil {
		return
	}
	return
}
