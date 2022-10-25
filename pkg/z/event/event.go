package event

import "github.com/go-serv/foundation/pkg/z"

type keyType int

const (
	NetClientNewContext keyType = iota + 1
	// NetClientBeforeDial event is being fired before a client connection will be established.
	NetClientBeforeDial
)

type NetClientBeforeDialArgs struct {
	Client     z.NetworkClientInterface
	TlsEnabled bool
}
