package codec

import "github.com/go-serv/service/internal/autogen/proto/net"

func init() {
	net.RegisterMessageWrapper(MessageWrapperHandler())
}
