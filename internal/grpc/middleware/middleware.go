// Package net
// The implementation of the network middleware group.
// Request handlers chain: wire data -> codec middleware: m1 -> m2 -> marshaler -> req -> network middleware: h1 -> h2 -> gRPC call
// Response handlers chain: gRPC call -> response -> h1 -> h2 -> codec middleware -> m2 -> m1 -> unmarshaler -> wire data
package middleware

import (
	"fmt"
	"github.com/go-serv/foundation/pkg/ancillary/slice"
	"github.com/go-serv/foundation/pkg/z"
)

type chainElement struct {
	key             any
	insertTargetKey any
	insertOp        z.InsertOp
	req             z.RequestHandlerFn
	res             z.ResponseHandlerFn
	disabled        bool
}

type middleware struct {
	els        []*chainElement
	serviceEls map[string][]*chainElement
}

func (m *middleware) Append(newKey any, req z.RequestHandlerFn, res z.ResponseHandlerFn) {
	newEl := &chainElement{key: newKey, req: req, res: res}
	m.els = append(m.els, newEl)
}

func (m *middleware) AppendToServiceChain(service string, newKey any, req z.RequestHandlerFn, res z.ResponseHandlerFn) {
	if _, has := m.serviceEls[service]; !has {
		m.serviceEls[service] = make([]*chainElement, 0)
	}
	newEl := &chainElement{key: newKey, req: req, res: res}
	m.serviceEls[service] = append(m.serviceEls[service], newEl)
}

func (m *middleware) Insert(targetKey any, op z.InsertOp, newKey any, req z.RequestHandlerFn, res z.ResponseHandlerFn) {
	newEl := &chainElement{key: newKey, req: req, res: res, insertTargetKey: targetKey, insertOp: op}
	m.els = append(m.els, newEl)
}

func (m *middleware) requestPassThrough(ctx z.NetContextInterface, service string) (err error) {
	var (
		curr z.NextHandlerFn
	)
	tailCall := func(next z.NextHandlerFn, _ z.NetContextInterface, _ z.RequestInterface) error {
		return ctx.Invoke()
	}
	handlers := m.requestHandlers(service)
	handlers = append(handlers, tailCall)
	// Iterate over the request handlers m. First added handler will be called first.
	for i := len(handlers) - 1; i >= 0; i-- {
		handler := handlers[i]
		next := curr
		curr = func(req z.RequestInterface, _ z.ResponseInterface) (el z.NextHandlerFn, err error) {
			if err = handler(next, ctx, req); err != nil {
				return
			}
			return curr, nil
		}
	}
	_, err = curr(ctx.Request(), nil)
	return
}

func (m *middleware) responsePassThrough(ctx z.NetContextInterface, service string) (err error) {
	var (
		curr z.NextHandlerFn
	)
	//
	tailCall := func(_ z.NextHandlerFn, _ z.NetContextInterface, res z.ResponseInterface) error {
		return nil
	}
	handlers := m.responseHandlers(service)
	handlers = append(handlers, tailCall)
	// Iterate over the response handlers m. First added handler will be called last.
	for i := 0; i < len(handlers); i++ {
		handler := handlers[i]
		next := curr
		curr = func(_ z.RequestInterface, res z.ResponseInterface) (z.NextHandlerFn, error) {
			if err = handler(next, ctx, res); err != nil {
				return nil, err
			}
			return curr, nil
		}
	}
	_, err = curr(nil, ctx.Response())
	return
}

func (m *middleware) findElementByKey(search any, els []*chainElement) int {
	for i, el := range els {
		if el.key == search {
			return i
		}
	}
	return -1
}

func chainElementNotFound(key any) {
	panic(fmt.Sprintf("middleware: failed to find chain element with key '%v'", key))
}

func (m *middleware) orderChainElements(unordered []*chainElement) (ordered []*chainElement) {
	ordered = make([]*chainElement, 0)
	// First append elements that do not require reordering.
	for _, el := range unordered {
		if el.insertOp == 0 {
			ordered = append(ordered, el)
		}
	}
	// Append elements that require a specific order.
	for _, el := range unordered {
		switch el.insertOp {
		case z.InsertBefore:
			if i := m.findElementByKey(el.key, ordered); i != -1 {
				chainElementNotFound(el.key)
			} else {
				ordered = slice.InsertBefore[*chainElement](ordered, el, i)
			}
		case z.InsertAfter:
			if i := m.findElementByKey(el.key, ordered); i != -1 {
				chainElementNotFound(el.key)
			} else {
				ordered = slice.InsertAfter[*chainElement](ordered, el, i)
			}
		}
	}
	return
}

func (m *middleware) serviceChain(service string) (unordered []*chainElement) {
	for _, el := range m.els {
		unordered = append(unordered, el)
	}
	if _, has := m.serviceEls[service]; has {
		for _, sEl := range m.serviceEls[service] {
			unordered = append(unordered, sEl)
		}
	}
	return
}

func (m *middleware) requestHandlers(currentService string) (handlers []z.RequestHandlerFn) {
	ordered := m.orderChainElements(m.serviceChain(currentService))
	for _, el := range ordered {
		if el.disabled {
			continue
		}
		handlers = append(handlers, el.req)
	}
	return
}

func (m *middleware) responseHandlers(currentService string) (handlers []z.ResponseHandlerFn) {
	ordered := m.orderChainElements(m.serviceChain(currentService))
	for _, el := range ordered {
		if el.disabled {
			continue
		}
		handlers = append(handlers, el.res)
	}
	return
}
