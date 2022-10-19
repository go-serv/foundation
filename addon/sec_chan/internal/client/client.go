package client

import (
	"github.com/go-serv/foundation/addon/sec_chan/internal/codec"
	"github.com/go-serv/foundation/addon/sec_chan/x"
	proto "github.com/go-serv/foundation/internal/autogen/net/sec_chan"
	grpc_client "github.com/go-serv/foundation/internal/client"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
)

var (
	serviceName = proto.SecureChannel_ServiceDesc.ServiceName
)

type impl struct {
	grpc_client.NetOptions
	c *client
}

type client struct {
	z.NetworkClientInterface
	impl
	grpcClient proto.SecureChannelClient
}

func (c *client) OnConnect(cc grpc.ClientConnInterface) {
	c.grpcClient = proto.NewSecureChannelClient(cc, codec.Name)
	c.impl.c = c
}

func (c *client) Create(lifetime int, nonceLen int, algoType x.KeyExchangeAlgoType) (pubKey []byte, nonce []byte, err error) {
	var (
		res *proto.Create_Response
	)
	// Populate the request data and make call.
	req := &proto.Create_Request{}
	req.Lifetime = uint32(lifetime)
	req.NonceLength = uint32(nonceLen)
	if res, err = c.impl.Create(req, algoType); err != nil {
		return
	}
	pubKey = res.GetPubKey()
	nonce = res.GetNonce()
	return
}

func (c *client) Close() (err error) {
	var (
		res *proto.Close_Response
	)
	req := &proto.Close_Request{}
	if res, err = c.impl.Close(req); err != nil {
		return
	}
	// Notice: left for future use.
	_ = res
	return
}
