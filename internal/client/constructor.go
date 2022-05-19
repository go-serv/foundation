package client

func NewLocalClient() *localClient {
	c := new(localClient)
	return c
}

func NewNetClient() *netClient {
	c := new(netClient)
	return c
}
