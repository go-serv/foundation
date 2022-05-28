package local

import i "github.com/go-serv/service/internal"

type localClient struct {
	i.ClientInterface
	svc i.LocalServiceInterface
}

func (c *localClient) LocalService() i.LocalServiceInterface {
	return c.svc
}
