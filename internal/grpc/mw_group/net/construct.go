package net

func NewMiddleware() *netMiddleware {
	g := new(netMiddleware)
	return g
}

func (mw *netMiddleware) newRequestChain() *requestChain {
	r := new(requestChain)
	r.mw = mw
	return r
}

func (mw *netMiddleware) newResponseChain() *responseChain {
	r := new(responseChain)
	r.mw = mw
	return r
}
