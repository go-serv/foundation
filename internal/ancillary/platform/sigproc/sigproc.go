package sigproc

/*
//#include "_cgo_export.h"
extern void go_callback_int(void*, int);
*/
import "C"
import (
	"fmt"
	"github.com/go-serv/service/internal/ancillary/platform"
)

//export sigproc_event_hander
func sigproc_event_hander(e int) {
	fmt.Printf("event %d\n", e)
}

type ProcSignalType uint

const (
	SharedMemRelease ProcSignalType = iota + 1
)

type platformTyp platform.PlatformType
