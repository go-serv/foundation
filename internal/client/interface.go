package client

type clientInterface interface {
}

type NetworkClientInterface interface {
	clientInterface
}

type LocalClientInterface interface {
	clientInterface
}
