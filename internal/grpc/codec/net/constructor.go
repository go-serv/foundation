package codec

import "github.com/go-serv/service/internal/service"

func NewNetCodec(svc service.NetworkServiceInterface) *codec {
	c := new(codec)
	c.svc = svc
	return c
}
