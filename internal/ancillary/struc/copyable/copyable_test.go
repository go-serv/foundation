package copyable

import "testing"

type A struct {
	Copyable
	One     string
	Num     int
	BoolVal bool
}

func defaultA() *A {
	a := new(A)
	a.One = "one"
	return a
}

func TestMerge(t *testing.T) {
	a := defaultA()
	b := &A{}
	b.Num = 1
	b.BoolVal = true
	Copyable{}.Merge(a, b)
	if !a.BoolVal || a.One != "one" || a.Num != 1 {
		t.Fatal("merge failed")
	}
}

func TestWrongType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("wrong type check failed")
		}
	}()
	a := &A{}
	Copyable{}.Merge(a, struct{ a int }{0})
}
