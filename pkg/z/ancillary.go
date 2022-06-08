package z

import "reflect"

type (
	DictionaryTypeHandlerFn func(op DictionaryOp, name, alias string, value reflect.Value)
	DictionaryOp            int
)

const (
	HydrateOp DictionaryOp = iota + 1
	DehydrateOp
)

type DictionaryInterface interface {
	RegisterTypeHandler(typ any, handler DictionaryTypeHandlerFn)
	Hydrate(target interface{}) error
	Dehydrate(target interface{}) error
}
