package slice_test

import (
	"github.com/go-serv/foundation/pkg/ancillary/slice"
	"testing"
)

func TestInsertAfter(t *testing.T) {
	var (
		in, out []int
	)
	in = []int{0}
	out = slice.InsertAfter[int](in, 1, 2)
	if out[0] != 1 {
		t.Fatalf("expected [1], got %v", out)
	}
	in = []int{2, 4, 5}
	out = slice.InsertAfter[int](in, 1, 3)
	if out[0] != 2 || out[1] != 3 || out[2] != 4 || out[3] != 5 {
		t.Fatalf("expected [2, 3, 4, 5], got %v", out)
	}
}
