package dictionary

import (
	bb "bytes"
	"github.com/go-serv/foundation/pkg/z"
	"reflect"
	"testing"
)

type PersonTyp struct {
	name string
	age  uint
}

type SampleDictionary struct {
	Dictionary
	Name     string    `name:"name"`
	PrimeNum uint      `name:"prime"`
	Person   PersonTyp `name:"person"`
	Bytes    []byte    `name:"bytes"`
}

type embeddedDic struct {
	SampleDictionary
}

const (
	primeNum   = 7
	authorName = "Andrei"
	authorAge  = 39 // alas not a constant in real life
)

var bytes = []byte{0x1, 0x2, 0x3}

func TestHydrate(t *testing.T) {
	dic := new(embeddedDic)
	dic.RegisterTypeHandler(reflect.TypeOf(uint(0)), func(op z.DictionaryOp, name, alias string, v reflect.Value) {
		switch op {
		case z.HydrateOp:
			v.SetUint(primeNum)
		}
	})
	dic.RegisterTypeHandler(reflect.TypeOf(bytes), func(op z.DictionaryOp, name, alias string, v reflect.Value) {
		switch op {
		case z.HydrateOp:
			v.SetBytes(bytes)
		}
	})
	dic.RegisterTypeHandler(reflect.TypeOf(""), func(op z.DictionaryOp, name, alias string, v reflect.Value) {
		switch op {
		case z.HydrateOp:
			v.SetString(authorName)
		}
	})
	dic.RegisterTypeHandler(reflect.TypeOf(PersonTyp{}), func(op z.DictionaryOp, name, alias string, v reflect.Value) {
		switch op {
		case z.HydrateOp:
			p := PersonTyp{
				name: authorName,
				age:  authorAge,
			}
			v.Set(reflect.ValueOf(p))
		}
	})
	err := dic.Hydrate(dic)
	if err != nil {
		t.Fatalf("hydrate: %v", err)
	}
	if bb.Compare(bytes, dic.Bytes) != 0 {
		t.Fatalf("hydrate: expected %v, got %v", bytes, dic.Bytes)
	}
	if dic.PrimeNum != primeNum {
		t.Fatalf("hydrate: expected uint %d, got %d", primeNum, dic.PrimeNum)
	}
	if dic.Person.name != authorName {
		t.Fatalf("hydrate: expected %s, got %s", authorName, dic.Person.name)
	}
	if dic.Person.age != authorAge {
		t.Fatalf("hydrate: expected %d, got %d", authorAge, dic.Person.age)
	}
}
