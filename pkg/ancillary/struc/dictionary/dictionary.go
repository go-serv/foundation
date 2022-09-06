package dictionary

import (
	"fmt"
	"github.com/go-serv/foundation/pkg/z"
	"reflect"
)

type (
	typeHandlersMap map[any]z.DictionaryTypeHandlerFn
)

type Dictionary struct {
	typeHandlers typeHandlersMap
}

func (d *Dictionary) RegisterTypeHandler(typ any, handler z.DictionaryTypeHandlerFn) {
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
		if fld.Type == reflect.TypeOf((*Dictionary)(nil)).Elem() { // base type Dictionary, skip
			continue
		}
		if !fld.IsExported() {
			return fmt.Errorf("dictionary: field '%s' must be exported", fld.Name)
		}
		fldVal := indir.Field(i).Addr().Interface()
		if _, ok := fldVal.(z.DictionaryInterface); ok {
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
		handler(z.HydrateOp, name, alias, v.Elem())
	})
}

func (d *Dictionary) Dehydrate(target interface{}) error {
	return d.iterateOver(target, func(name, alias string, t reflect.Type, v reflect.Value) {
		handler, ok := d.typeHandlers[t]
		if !ok {
			return
		}
		handler(z.DehydrateOp, name, alias, v.Elem())
	})
}
