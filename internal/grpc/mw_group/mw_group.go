package mw_group

import (
	i "github.com/go-serv/service/internal"
)

type RequestMiddlewareHandler func(r i.RequestInterface, svc interface{}) error
type ResponseMiddlewareHandler func(r i.ResponseInterface, svc interface{}) error

type GroupItem struct {
	ReqHandler RequestMiddlewareHandler
	ResHandler ResponseMiddlewareHandler
}

type MwGroup struct {
	Items  []*GroupItem
	Target interface{}
}

func (mw *MwGroup) AddItem(item *GroupItem) {
	mw.Items = append(mw.Items, item)
}
