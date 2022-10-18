package middleware

func NewServerMiddleware() *serverMw {
	mw := new(serverMw)
	return mw
}

func NewClientMiddleware() *clientMw {
	mw := new(clientMw)
	return mw
}
