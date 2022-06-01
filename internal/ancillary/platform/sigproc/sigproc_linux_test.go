//go:build linux && cgo

package sigproc

import (
	"testing"
	"time"
)

func TestSignalProcess(t *testing.T) {
	SignalProcess(0, 1)
	time.Sleep(time.Millisecond)
}
