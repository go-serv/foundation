package dictionary

import (
	"fmt"
	"reflect"
)

var (
	typeHandlers typeHandlersMap
)

type (
	ImportHandlerFn func(target DictionaryInterface, name, alias string, value reflect.Value) error
	ExportHandlerFn func(target DictionaryInterface, name, alias string, value reflect.Value) error
)

type (
	typeHandlersMapItem struct {
		imp ImportHandlerFn
		exp ExportHandlerFn
	}
	typeHandlersMap map[any]typeHandlersMapItem
)

type Dictionary struct{}

func RegisterTypeHandlers(typ any, imp ImportHandlerFn, exp ExportHandlerFn) {
	typeHandlers[typ] = typeHandlersMapItem{imp, exp}
}

func (d Dictionary) iterateOver(struc interface{}, fn func(name, alias string, t reflect.Type, v reflect.Value) error) (err error) {
	var strucType reflect.Type
	strucVal := reflect.ValueOf(struc)
	indir := reflect.Indirect(strucVal)

	switch strucVal.Kind() {
	case reflect.Pointer:
		strucType = reflect.TypeOf(struc).Elem()
	default:
		strucType = strucVal.Type()
	}

	for i := 0; i < strucType.NumField(); i++ {
		strucField := strucType.Field(i)
		if !strucField.IsExported() {
			continue
		}
		if strucField.Type == reflect.TypeOf((*Dictionary)(nil)).Elem() { // base type Dictionary, skip
			continue
		}
		fv := indir.Field(i).Interface()
		fk := indir.Field(i).Kind()
		if fk == reflect.Pointer || fk == reflect.Interface {
			if err = d.iterateOver(fv, fn); err != nil {
				return err
			}
			continue
		}
		name, ok := strucField.Tag.Lookup("name")
		if !ok {
			return fmt.Errorf("dictionary: an item name must be provided for the field '%s' with tag `name:`", strucField.Name)
		}
		alias, _ := strucField.Tag.Lookup("alias")
		err = fn(name, alias, strucField.Type, indir.Field(i).Addr())
	}
	return nil
}

func (d Dictionary) Import(target DictionaryInterface) (err error) {
	err = d.iterateOver(target, func(name, alias string, t reflect.Type, v reflect.Value) error {
		item, ok := typeHandlers[t]
		if !ok || item.imp == nil {
			return nil
		}
		return item.imp(target, name, alias, v.Elem())
	})
	return
}

func (d Dictionary) Export(target DictionaryInterface) (err error) {
	err = d.iterateOver(target, func(name, alias string, t reflect.Type, v reflect.Value) error {
		item, ok := typeHandlers[t]
		if !ok || item.exp == nil {
			return nil
		}
		return item.exp(target, name, alias, v.Elem())
	})
	return
}
