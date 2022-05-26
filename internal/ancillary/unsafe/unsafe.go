package unsafe

import "unsafe"

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
