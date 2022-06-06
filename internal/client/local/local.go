package local

import "github.com/go-serv/service/pkg/z"

type localClient struct {
	z.ClientInterface
	svc z.LocalServiceInterface
}

func (c *localClient) LocalService() z.LocalServiceInterface {
	return c.svc
}
