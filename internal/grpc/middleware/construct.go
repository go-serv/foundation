package middleware

func NewMiddleware() *mwHandlersChain {
	g := new(mwHandlersChain)
	return g
}
