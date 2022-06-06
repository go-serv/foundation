package net

func NewMiddlewareGroup() *netMwGroup {
	g := new(netMwGroup)
	return g
}

func (mw *netMwGroup) newRequestChain() *requestChain {
	r := new(requestChain)
	r.mwGroup = mw
	return r
}

func (mw *netMwGroup) newResponseChain() *responseChain {
	r := new(responseChain)
	r.mwGroup = mw
	return r
}
