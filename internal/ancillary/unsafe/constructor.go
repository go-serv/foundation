package unsafe

import "unsafe"

func NewPointer(v any) unsafe.Pointer {
	return unsafe.Pointer(&v)
}
