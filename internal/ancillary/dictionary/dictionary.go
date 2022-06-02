package dictionary

import (
	"fmt"
	"reflect"
	"unsafe"
)

type OpType int

const (
	HydrateOp OpType = iota + 1
	DehydrateOp
)

type (
	TypeHandler     func(op OpType, name, alias string, value reflect.Value)
	typeHandlersMap map[any]TypeHandler
)

type Dictionary struct {
	typeHandlers typeHandlersMap
}

func (d *Dictionary) RegisterTypeHandler(typ any, handler TypeHandler) {
	if d.typeHandlers == nil {
		d.typeHandlers = make(typeHandlersMap)
	}
	d.typeHandlers[typ] = handler
}

func (d *Dictionary) iterateOver(target interface{}, fn func(name, alias string, t reflect.Type, v reflect.Value)) error {
	st := reflect.TypeOf(target).Elem()
	n := st.NumField()
	for i := 1; i < n; i++ {
		field := st.Field(i)
		if !field.IsExported() {
			return fmt.Errorf("dictionary: field '%s' must be exported", field.Name)
		}
		name, ok := field.Tag.Lookup("name")
		if !ok {
			return fmt.Errorf("dictionary: an item name must be provided for the field '%s' with tag `name:`", field.Name)
		}
		alias, _ := field.Tag.Lookup("alias")
		fieldPtr := reflect.ValueOf(target).Pointer() + field.Offset
		fieldVal := reflect.NewAt(field.Type, unsafe.Pointer(fieldPtr))
		fn(name, alias, field.Type, fieldVal)
	}
	return nil
}

func (d *Dictionary) Hydrate(target interface{}) error {
	return d.iterateOver(target, func(name, alias string, t reflect.Type, v reflect.Value) {
		handler, ok := d.typeHandlers[t]
		if !ok {
			return
		}
		handler(HydrateOp, name, alias, v.Elem())
	})
}

func (d *Dictionary) Dehydrate(target interface{}) error {
	return d.iterateOver(target, func(name, alias string, t reflect.Type, v reflect.Value) {
		handler, ok := d.typeHandlers[t]
		if !ok {
			return
		}
		handler(DehydrateOp, name, alias, v.Elem())
	})
}
