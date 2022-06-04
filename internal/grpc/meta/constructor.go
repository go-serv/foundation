package meta

func NewServerMeta() *serverMeta {
	s := new(serverMeta)
	return s
}

func NewClientMeta() *clientMeta {
	c := new(clientMeta)
	return c
}
