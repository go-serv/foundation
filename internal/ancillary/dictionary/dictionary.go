package dictionary

import (
	"fmt"
	"reflect"
)

const (
	HydrateOp OpType = iota + 1
	DehydrateOp
)

type (
	TypeHandler     func(op OpType, name, alias string, value reflect.Value)
	typeHandlersMap map[any]TypeHandler
	OpType          int
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

func (d *Dictionary) iterateOver(struc interface{}, fn func(name, alias string, t reflect.Type, v reflect.Value)) error {
	var typ reflect.Type
	iv := reflect.ValueOf(struc)
	indir := reflect.Indirect(iv)
	if iv.Kind() == reflect.Pointer {
		typ = reflect.TypeOf(struc).Elem()
	} else {
		typ = iv.Type()
	}
	for i := 0; i < typ.NumField(); i++ {
		fld := typ.Field(i)
		if fld.Type == reflect.TypeOf((*Dictionary)(nil)).Elem() {
			continue
		}
		if !fld.IsExported() {
			return fmt.Errorf("dictionary: field '%s' must be exported", fld.Name)
		}
		fldVal := indir.Field(i).Addr().Interface()
		// tt := reflect.TypeOf(struc).Implements(reflect.TypeOf((*DictionaryInterface)(nil)).Elem())
		// _ = tt
		if _, ok := fldVal.(DictionaryInterface); ok {
			if err := d.iterateOver(fldVal, fn); err != nil {
				return err
			}
			continue
		}
		name, ok := fld.Tag.Lookup("name")
		if !ok {
			return fmt.Errorf("dictionary: an item name must be provided for the field '%s' with tag `name:`", fld.Name)
		}
		alias, _ := fld.Tag.Lookup("alias")
		fn(name, alias, fld.Type, indir.Field(i).Addr())
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
