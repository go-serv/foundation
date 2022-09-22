package client

import (
	"github.com/go-serv/foundation/addon/sec-chan/internal/codec"
	proto "github.com/go-serv/foundation/internal/autogen/sec_chan_mw"
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

func (c *client) NewSession(lifetime int, nonceLen int) (pubKey []byte, nonce []byte, err error) {
	var (
		res *proto.Session_Response
	)
	// Populate the request data and make call.
	req := &proto.Session_Request{}
	req.Lifetime = uint32(lifetime)
	req.NonceLength = uint32(nonceLen)
	if res, err = c.impl.SecureSession(req); err != nil {
		return
	}

	pubKey = res.GetPubKey()
	nonce = res.GetNonce()
	return
}
