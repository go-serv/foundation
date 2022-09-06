// Package net
// The implementation of the network middleware group.
// Request handlers chain: wire data -> codec middleware: m1 -> m2 -> marshaler -> req -> network middleware: h1 -> h2 -> gRPC call
// Response handlers chain: gRPC call -> response -> h1 -> h2 -> codec middleware -> m2 -> m1 -> unmarshaler -> wire data
package net

import (
	"github.com/go-serv/foundation/pkg/z"
)

type netMiddleware struct {
	client            z.ClientInterface
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

func (m *netMiddleware) Client() z.ClientInterface {
	return m.client
}

func (m *netMiddleware) WithClient(client z.ClientInterface) {
	m.client = client
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

func (t *requestChain) passThrough(ctx z.NetContextInterface) (err error) {
	var curr z.NetChainElementFn
	invokeHandler := func(next z.NetChainElementFn, _ z.NetContextInterface, _ z.RequestInterface) error {
		return ctx.Invoke()
	}
	// Build a middleware chain so that the first added handler will be called first.
	ch := append(t.mw.reqHandlers, invokeHandler)
	l1 := len(ch)
	for i := l1 - 1; i >= 0; i-- {
		handler := ch[i]
		next := curr
		curr = func(req z.RequestInterface, _ z.ResponseInterface) (el z.NetChainElementFn, err error) {
			if err = handler(next, ctx, req); err != nil {
				return
			}
			return curr, nil
		}
	}
	_, err = curr(ctx.Request(), nil)
	return
}

func (t *responseChain) passThrough(ctx z.NetContextInterface) (err error) {
	var curr z.NetChainElementFn
	tailCall := func(_ z.NetChainElementFn, _ z.NetContextInterface, res z.ResponseInterface) error {
		return nil
	}
	//
	ch := append([]z.NetResponseHandlerFn{tailCall}, t.mw.resHandlers...)
	l1 := len(ch)
	for i := 0; i < l1; i++ {
		handler := ch[i]
		next := curr
		curr = func(_ z.RequestInterface, res z.ResponseInterface) (z.NetChainElementFn, error) {
			if err = handler(next, ctx, res); err != nil {
				return nil, err
			}
			return curr, nil
		}
	}
	_, err = curr(nil, ctx.Response())
	return
}
