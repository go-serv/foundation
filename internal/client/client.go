package client

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/internal/ancillary"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
	"net"
)

type client struct {
	svcName  string
	codec    z.CodecInterface
	mw       z.MiddlewareInterface
	meta     z.MetaInterface
	endpoint z.EndpointInterface
	conn     net.Conn
	dialOpts []grpc.DialOption
	ancillary.MethodMustBeImplemented
}

func (c *client) ServiceName() string {
	return c.svcName
}

func (c *client) Codec() z.CodecInterface {
	return c.codec
}

func (s *client) WithCodec(cc z.CodecInterface) {
	s.codec = cc
}

func (c *client) Middleware() z.MiddlewareInterface {
	return c.mw
}

func (c *client) WithMiddleware(mw z.MiddlewareInterface) {
	c.mw = mw
}

func (c *client) Meta() z.MetaInterface {
	return c.meta
}

func (c *client) WithMeta(meta z.MetaInterface) {
	c.meta = meta
}

func (c *client) Endpoint() z.EndpointInterface {
	return c.endpoint
}

func (c *client) WithDialOption(opts grpc.DialOption) {
	c.dialOpts = append(c.dialOpts, opts)
}

func (c *client) DialOptions() []grpc.DialOption {
	return c.dialOpts
}

func (c *client) OnConnect(cc grpc.ClientConnInterface) {
	c.MethodMustBeImplemented.Panic()
}

func (c *client) ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	return nil, nil, nil
}

func (c *client) WithOptions(any) {
	c.MethodMustBeImplemented.Panic()
}
