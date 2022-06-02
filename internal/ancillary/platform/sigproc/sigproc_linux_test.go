//go:build linux && cgo

package sigproc

import (
	"testing"
	"time"
)

func TestSignalProcess(t *testing.T) {
	SignalProcess(0, 255, 4456)
	time.Sleep(time.Millisecond)
}
