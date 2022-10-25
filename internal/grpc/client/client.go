package client

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/internal/ancillary"
	net_meta "github.com/go-serv/foundation/internal/grpc/meta/net"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"net"
)

type client struct {
	svcName  string
	codec    encoding.Codec
	mw       z.ClientMiddlewareInterface
	meta     z.MetaInterface
	endpoint z.EndpointInterface
	conn     net.Conn
	dialOpts []grpc.DialOption
	ancillary.MethodMustBeImplemented
}

func (c *client) ServiceName() string {
	return c.svcName
}

func (c *client) Codec() encoding.Codec {
	return c.codec
}

func (c *client) WithCodec(codec encoding.Codec) {
	c.codec = codec
}

func (c *client) Middleware() z.ClientMiddlewareInterface {
	return c.mw
}

func (c *client) WithMiddleware(mw z.ClientMiddlewareInterface) {
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

func (c *client) SessionId() z.SessionId {
	dic := c.Meta().Dictionary().(*net_meta.HttpDictionary)
	return dic.SessionId
}

func (c *client) Reset() {
	dic := c.Meta().Dictionary().(*net_meta.HttpDictionary)
	dic.SessionId = 0
}
