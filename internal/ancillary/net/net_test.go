package net_test

import (
	"bytes"
	"github.com/go-serv/foundation/internal/ancillary/net"
	"io"
	"testing"
)

var (
	utf8Str = "Привет, Мир!"
)

func TestNetReader(t *testing.T) {
	var err error
	strLen := []byte{0x0, 0x0, 0x0, byte(len(utf8Str))}
	testData := append(strLen, []byte(utf8Str)...)
	testData = append(testData, []byte{0x11, 0x22, 0x33, 0x44, 0x3, 0x0}...)
	reader := net.NewReader(testData)
	strGot, err := reader.ReadString()
	if err != nil {
		t.Fatalf("GenericNetReader: %v", err)
	}
	if strGot != utf8Str {
		t.Fatalf("GenericNetReader: expected %s, got %s", utf8Str, strGot)
	}
	var expected int32 = 0x11223344
	got, _ := net.GenericNetReader[int32](reader)
	if got != expected {
		t.Fatalf("GenericNetReader: expected %x, got %x", expected, got)
	}
	// Boolean
	var gotBool bool
	gotBool, _ = net.GenericNetReader[bool](reader)
	if gotBool != true {
		t.Fatalf("GenericNetReader: expected %t, got %t", true, gotBool)
	}
	gotBool, _ = net.GenericNetReader[bool](reader)
	if gotBool != false {
		t.Fatalf("GenericNetReader: expected %t, got %t", false, gotBool)
	}
	// End of input data
	_, eof := net.GenericNetReader[int64](reader)
	if eof != io.EOF {
		t.Fatalf("GenericNetReader: expected end of input data")
	}
}

func TestNetWriter(t *testing.T) {
	w := net.NewWriter()
	w.WriteString(utf8Str)
	w.Write([]byte{0x55})
	net.GenericNetWriter[bool](w, false)
	net.GenericNetWriter[uint32](w, 255<<24)
	net.GenericNetWriter[bool](w, true)
	expectedData := append([]byte{0x0, 0x0, 0x0, byte(len(utf8Str))}, []byte(utf8Str)...)
	expectedData = append(expectedData, []byte{0x55, 0x0, 0xff, 0x0, 0x0, 0x0, 0x1}...)
	got := w.Bytes()
	if bytes.Compare(expectedData, got) != 0 {
		t.Fatalf("GenericNetWriter: expected %v, got %v", expectedData, got)
	}
}
