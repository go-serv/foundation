package ancillary_test

import (
	"bytes"
	"github.com/go-serv/service/internal/ancillary"
	"io"
	"testing"
)

func TestNetReaderWriter(t *testing.T) {
	var expected int32 = 0x11223344
	reader := ancillary.NewNetReader([]byte{0x11, 0x22, 0x33, 0x44, 0x3, 0x0})
	got, _ := ancillary.NetReader[int32](reader)
	if got != expected {
		t.Fatalf("NetReader: expected %x, got %x", expected, got)
	}
	// Boolean
	var gotBool bool
	gotBool, _ = ancillary.NetReader[bool](reader)
	if gotBool != true {
		t.Fatalf("NetReader: expected %t, got %t", true, gotBool)
	}
	gotBool, _ = ancillary.NetReader[bool](reader)
	if gotBool != false {
		t.Fatalf("NetReader: expected %t, got %t", false, gotBool)
	}
	// End of input data
	_, eof := ancillary.NetReader[int64](reader)
	if eof != io.EOF {
		t.Fatalf("NetReader: expected end of input data")
	}
}

func TestNetWriter(t *testing.T) {
	w := ancillary.NewNetWriter()
	ancillary.NetWriter[bool](w, false)
	ancillary.NetWriter[uint32](w, (255 << 24))
	ancillary.NetWriter[bool](w, true)
	expected := []byte{0x0, 0xff, 0x0, 0x0, 0x0, 0x1}
	got := w.Bytes()
	if bytes.Compare(expected, got) != 0 {
		t.Fatalf("NetWriter: expected %v, got %v", expected, got)
	}
}
