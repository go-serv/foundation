// Package net
// The implementation of the network middleware group.
// Request handlers chain: wire data -> codec middleware: m1 -> m2 -> marshaler -> req -> network middleware: h1 -> h2 -> gRPC call
// Response handlers chain: gRPC call -> response -> h1 -> h2 -> codec middleware -> m2 -> m1 -> unmarshaler -> wire data
package net

import "github.com/go-serv/service/pkg/z"

type netMiddleware struct {
	preStreamHandlers []z.NetPreStreamHandlerFn
	reqHandlers       []z.NetRequestHandlerFn
	resHandlers       []z.NetResponseHandlerFn
}

type chain struct {
	mw *netMiddleware
}

type responseChain struct {
	chain
}

type requestChain struct {
	chain
}

func (m *netMiddleware) AddPreStreamHandler(h z.NetPreStreamHandlerFn) {
	m.preStreamHandlers = append(m.preStreamHandlers, h)
}

func (m *netMiddleware) AddRequestHandler(h z.NetRequestHandlerFn) {
	m.reqHandlers = append(m.reqHandlers, h)
}

func (m *netMiddleware) AddResponseHandler(h z.NetResponseHandlerFn) {
	m.resHandlers = append(m.resHandlers, h)
}

func (t *requestChain) passThrough(call z.NetContextInterface) (res z.ResponseInterface, err error) {
	var curr z.NetChainElementFn
	invokeHandler := func(next z.NetChainElementFn, _ z.RequestInterface, res z.ResponseInterface) (err error) {
		var payload interface{}
		if payload, err = call.Invoke(); err != nil {
			return
		}
		res.WithPayload(payload)
		return
	}
	// Build a middleware chain so that the first added handler will be called first.
	ch := append(t.mw.reqHandlers, invokeHandler)
	l1 := len(ch)
	for i := l1 - 1; i >= 0; i-- {
		handler := ch[i]
		next := curr
		curr = func(req z.RequestInterface, res z.ResponseInterface) (el z.NetChainElementFn, err error) {
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
	var curr z.NetChainElementFn
	tailCall := func(_ z.NetChainElementFn, res z.ResponseInterface) (err error) {
		out = res.Payload()
		return
	}
	//
	ch := append([]z.NetResponseHandlerFn{tailCall}, t.mw.resHandlers...)
	l1 := len(ch)
	for i := 0; i < l1; i++ {
		handler := ch[i]
		next := curr
		curr = func(_ z.RequestInterface, res z.ResponseInterface) (z.NetChainElementFn, error) {
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
