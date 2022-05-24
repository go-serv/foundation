package codec

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"unsafe"
)

type CodecInterceptorHandler func(b []byte) ([]byte, error)

// maps a proto message to its unmarshaler
type msgUnmarshalerMap map[uintptr]UnmarshalerInterface

type ifaceHeader struct {
	typ   unsafe.Pointer
	pdata unsafe.Pointer
}

type pointer struct {
	p unsafe.Pointer
}

func (p pointer) pointerOfEmptyIface(v any) {
	p.p = (*ifaceHeader)(unsafe.Pointer(&v)).pdata
}

type codec struct {
	df                *DataFrame
	interceptorsChain []CodecInterceptorHandler
	msgUnmarshaler    msgUnmarshalerMap
	unMu              sync.Mutex
}

func (c *codec) mapProtoMessage(protoMsg interface{}, un UnmarshalerInterface) {
	c.unMu.Lock()
	defer c.unMu.Unlock()
	ptr := pointer{}
	ptr.pointerOfEmptyIface(protoMsg)
	key := uintptr(ptr.p)
	c.msgUnmarshaler[key] = un
}

func (c *codec) GetUnmarshalerByProtoMessage(protoMsg interface{}) (UnmarshalerInterface, error) {
	c.unMu.Lock()
	defer c.unMu.Unlock()
	ptr := pointer{}
	ptr.pointerOfEmptyIface(protoMsg)
	key := uintptr(ptr.p)
	v, ok := c.msgUnmarshaler[key]
	if !ok {
		return nil, status.Error(codes.Internal, "failed to retrieve unmarshaler for the provided proto message")
	}
	return v, nil
}

func (c *codec) ChainInterceptorHandler(h CodecInterceptorHandler) {
	c.interceptorsChain = append(c.interceptorsChain, h)
}
