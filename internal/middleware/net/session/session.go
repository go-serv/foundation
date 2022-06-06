package session

import z "github.com/go-serv/service/internal"

func ServerSessionHandler(next z.NetChainElement, req z.RequestInterface, res z.ResponseInterface) (err error) {
	_, err = next(req, res)
	return
}
