// Package net
// The implementation of the network middleware group.
// Request handlers chain: wire data -> codec middleware: m1 -> m2 -> marshaler -> req -> network middleware: h1 -> h2 -> gRPC call
// Response handlers chain: gRPC call -> response -> h1 -> h2 -> codec middleware -> m2 -> m1 -> unmarshaler -> wire data
package middleware

import (
	"github.com/go-serv/foundation/pkg/z"
)

type mwHandlersChain struct {
	client            z.ClientInterface
	preStreamHandlers []z.MiddlewarePreStreamHandlerFn
	reqHandlers       []z.MiddlewareRequestHandlerFn
	resHandlers       []z.MiddlewareResponseHandlerFn
}

func (chain *mwHandlersChain) Client() z.ClientInterface {
	return chain.client
}

func (chain *mwHandlersChain) WithClient(client z.ClientInterface) {
	chain.client = client
}

func (chain *mwHandlersChain) AddPreStreamHandler(h z.MiddlewarePreStreamHandlerFn) {
	chain.preStreamHandlers = append(chain.preStreamHandlers, h)
}

func (chain *mwHandlersChain) AddRequestHandler(h z.MiddlewareRequestHandlerFn) {
	chain.reqHandlers = append(chain.reqHandlers, h)
}

func (chain *mwHandlersChain) AddResponseHandler(h z.MiddlewareResponseHandlerFn) {
	chain.resHandlers = append(chain.resHandlers, h)
}

func (dst *mwHandlersChain) MergeWithParent(pi z.MiddlewareInterface) {
	parent := pi.(*mwHandlersChain)
	dst.reqHandlers = append(parent.reqHandlers, dst.reqHandlers...)
	dst.resHandlers = append(parent.resHandlers, dst.resHandlers...)
}

func (chain *mwHandlersChain) requestPassThrough(ctx z.NetContextInterface) (err error) {
	var (
		curr z.MiddlewareChainElementFn
	)

	tailCall := func(next z.MiddlewareChainElementFn, _ z.NetContextInterface, _ z.RequestInterface) error {
		return ctx.Invoke()
	}
	ch := append(chain.reqHandlers, tailCall)

	// Iterate over the request handlers chain. First added handler will be called first.
	for i := len(ch) - 1; i >= 0; i-- {
		handler := ch[i]
		next := curr
		curr = func(req z.RequestInterface, _ z.ResponseInterface) (el z.MiddlewareChainElementFn, err error) {
			if err = handler(next, ctx, req); err != nil {
				return
			}
			return curr, nil
		}
	}
	_, err = curr(ctx.Request(), nil)
	return
}

func (chain *mwHandlersChain) responsePassThrough(ctx z.NetContextInterface) (err error) {
	var (
		curr z.MiddlewareChainElementFn
	)

	// A stub for the last call.
	tailCall := func(_ z.MiddlewareChainElementFn, _ z.NetContextInterface, res z.ResponseInterface) error {
		return nil
	}
	ch := append([]z.MiddlewareResponseHandlerFn{tailCall}, chain.resHandlers...)

	// Iterate over the response handlers chain. First added handler will be called last.
	for i := 0; i < len(ch); i++ {
		handler := ch[i]
		next := curr
		curr = func(_ z.RequestInterface, res z.ResponseInterface) (z.MiddlewareChainElementFn, error) {
			if err = handler(next, ctx, res); err != nil {
				return nil, err
			}
			return curr, nil
		}
	}
	_, err = curr(nil, ctx.Response())
	return
}
