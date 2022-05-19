package mw_group

func NewMiddlewareGroup() *mwGroup {
	g := new(mwGroup)
	return g
}

func NewItem(reqHandler RequestMiddlewareHandler, resHandler ResponseMiddlewareHandler) *GroupItem {
	return &GroupItem{reqHandler, resHandler}
}
