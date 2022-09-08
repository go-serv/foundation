package copyable

import "reflect"

type Shallow struct{}

func (Shallow) Merge(dst any, src any) {
	dstv := reflect.Indirect(reflect.ValueOf(dst))
	srcv := reflect.Indirect(reflect.ValueOf(src))
	if dstv.Type() != srcv.Type() {
		panic("copyable structures must be of the same type")
	}
	for i := 0; i < dstv.NumField(); i++ {
		srcf := srcv.Field(i)
		if srcf.Type().Kind() == reflect.Struct {
			Shallow{}.Merge(reflect.Indirect(dstv.Field(i).Addr().Elem()).Interface(),
				reflect.Indirect(srcf.Addr().Elem()).Interface())
		}
		n := srcf.Type().Name()
		_ = n
		if !srcf.IsZero() {
			dstv.Field(i).Addr().Elem().Set(srcf.Addr().Elem())
		}
	}
}
