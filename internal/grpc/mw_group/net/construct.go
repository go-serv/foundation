package net

func NewMiddlewareGroup(target interface{}) *netMwGroup {
	g := new(netMwGroup)
	g.Target = target
	return g
}
