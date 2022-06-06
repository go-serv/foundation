// Package net
// The implementation of the network middleware group.
// Request handlers chain: wire data -> codec middleware: m1 -> m2 -> marshaler -> req -> network middleware: h1 -> h2 -> gRPC call
// Response handlers chain: gRPC call -> response -> h1 -> h2 -> codec middleware -> m2 -> m1 -> unmarshaler -> wire data
package net

import (
	z "github.com/go-serv/service/internal"
)

type netMwGroup struct {
	preStreamHandlers []z.NetPreStreamHandlerFn
	reqHandlers       []z.NetRequestHandlerFn
	resHandlers       []z.NetResponseHandlerFn
}

type chain struct {
	mwGroup *netMwGroup
}

type responseChain struct {
	chain
}

type requestChain struct {
	chain
}

func (m *netMwGroup) AddPreStreamHandler(h z.NetPreStreamHandlerFn) {
	m.preStreamHandlers = append(m.preStreamHandlers, h)
}

func (m *netMwGroup) AddRequestHandler(h z.NetRequestHandlerFn) {
	m.reqHandlers = append(m.reqHandlers, h)
}

func (m *netMwGroup) AddResponseHandler(h z.NetResponseHandlerFn) {
	m.resHandlers = append(m.resHandlers, h)
}

func (t *requestChain) passThrough(call z.NetContextInterface) (res z.ResponseInterface, err error) {
	var curr z.NetChainElement
	invokeHandler := func(next z.NetChainElement, _ z.RequestInterface, res z.ResponseInterface) (err error) {
		var payload interface{}
		if payload, err = call.Invoke(); err != nil {
			return
		}
		res.WithPayload(payload)
		return
	}
	// Build a middleware chain so that the first added handler will be called first.
	ch := append(t.mwGroup.reqHandlers, invokeHandler)
	l1 := len(ch)
	for i := l1 - 1; i >= 0; i-- {
		handler := ch[i]
		next := curr
		curr = func(req z.RequestInterface, res z.ResponseInterface) (el z.NetChainElement, err error) {
			err = handler(next, req, res)
			if err != nil {
				return
			}
			return curr, nil
		}
	}
	_, err = curr(call.Request(), call.Response())
	return call.Response(), err
}

func (t *responseChain) passThrough(res z.ResponseInterface) (out interface{}, err error) {
	var curr z.NetChainElement
	tailCall := func(_ z.NetChainElement, res z.ResponseInterface) (err error) {
		out = res.Payload()
		return
	}
	//
	ch := append([]z.NetResponseHandlerFn{tailCall}, t.mwGroup.resHandlers...)
	l1 := len(ch)
	for i := 0; i < l1; i++ {
		handler := ch[i]
		next := curr
		curr = func(_ z.RequestInterface, res z.ResponseInterface) (z.NetChainElement, error) {
			err = handler(next, res)
			if err != nil {
				return nil, err
			}
			return curr, nil
		}
	}
	_, err = curr(nil, res)
	return
}
