package mw_group

func NewItem(reqHandler RequestMiddlewareHandler, resHandler ResponseMiddlewareHandler) *GroupItem {
	return &GroupItem{reqHandler, resHandler}
}
