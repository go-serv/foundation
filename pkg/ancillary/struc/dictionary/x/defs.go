package x

import "reflect"

type (
	ImportHandlerFn func(target DictionaryInterface, name, alias string, value reflect.Value) error
	ExportHandlerFn func(target DictionaryInterface, name, alias string, value reflect.Value) error
)
