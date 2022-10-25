package event

type keyType int

const (
	NetClientNewContext keyType = iota + 1
	// NetClientBeforeDial event is being fired before a client connection will be established.
	// List of handler arguments:
	// 	args[0]<NetworkClientInterface>. Client instance.
	// 	args[1]<bool>. If true, client will use a TLS connection.
	NetClientBeforeDial
)
