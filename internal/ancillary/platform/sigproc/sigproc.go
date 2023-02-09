package sigproc

/*
//#include "_cgo_export.h"
extern void go_callback_int(void*, int);
*/
import "C"
import (
	"fmt"
	"github.com/mesh-master/foundation/internal/ancillary/platform"
)

type SignalTyp uint

const (
	SharedMemRelease SignalTyp = iota + 1
)

type (
	SigCodeTyp  uint
	SigValueTyp uint
)

func (c SigCodeTyp) Validate() SigCodeTyp {
	return c
}

func (v SigValueTyp) Validate() SigValueTyp {
	return v
}

func (v SigValueTyp) pack(c SigCodeTyp) uint64 {
	return (uint64(c) << 54) | (uint64(v) & (^uint64(0) >> 10))
}

const (
	SharedMemoryRelease SignalTyp = iota + 1
)

//export sigproc_event_hander
func sigproc_event_hander(code int, extra_val int) {
	fmt.Printf("event code %d value %d\n", code, extra_val)
	//runtime.Runtime().TriggerEvent(SigCodeTyp(code).Validate(), SigValueTyp(extra_val).Validate())
}

type platformTyp platform.PlatformType
