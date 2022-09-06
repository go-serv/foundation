package codec

import "github.com/go-serv/foundation/internal/autogen/proto/net"

func init() {
	net.RegisterMessageWrapper(MessageWrapperHandler())
}
