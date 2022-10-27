package net

import (
	"encoding/base64"
	net_io "github.com/go-serv/foundation/pkg/ancillary/net/io"
	"github.com/go-serv/foundation/pkg/ancillary/struc/dictionary"
	"github.com/go-serv/foundation/pkg/ancillary/struc/dictionary/x"
	"github.com/go-serv/foundation/pkg/z"
	dic_defs "github.com/go-serv/foundation/pkg/z/dictionary"
	"reflect"
)

func uint64FromBase64(in string) (out uint64, err error) {
	var bytes []byte
	if bytes, err = base64.StdEncoding.DecodeString(in); err != nil {
		return
	}
	reader := net_io.NewReader(bytes)
	out, err = net_io.GenericNetReader[uint64](reader)
	return
}

func uint64ToBase64(in uint64) (out string, err error) {
	writer := net_io.NewWriter()
	if err = net_io.GenericNetWriter[uint64](writer, in); err != nil {
		return
	}
	out = base64.StdEncoding.EncodeToString(writer.Bytes())
	return
}

func init() {
	base64StringImp := func(target x.DictionaryInterface, name, alias string, v reflect.Value) (err error) {
		meta := target.(z.MetaInterface)
		if base64Str, has := meta.Get(name); has {
			var decoded []byte
			if decoded, err = base64.StdEncoding.DecodeString(base64Str); err != nil {
				v.SetBytes(decoded)
			}
		}
		return
	}
	base64StringExp := func(target x.DictionaryInterface, name, alias string, v reflect.Value) (err error) {
		return
	}
	dictionary.RegisterTypeHandlers(reflect.TypeOf((*dic_defs.Base64String)(nil)).Elem(), base64StringImp, base64StringExp)

	// Session ID header.
	sessImp := func(target x.DictionaryInterface, name, alias string, v reflect.Value) (err error) {
		meta := target.(z.MetaInterface)
		if idStr, has := meta.Get(name); has {
			var sId uint64
			if sId, err = uint64FromBase64(idStr); err != nil {
				return
			}
			v.Set(reflect.ValueOf(z.SessionId(sId)))
		}
		return
	}
	sessExp := func(target x.DictionaryInterface, name, alias string, v reflect.Value) (err error) {
		if v.IsZero() {
			return
		}
		var idStr string
		if idStr, err = uint64ToBase64(v.Uint()); err != nil {
			return
		}
		meta := target.(z.MetaInterface)
		meta.Set(name, idStr)
		return
	}
	dictionary.RegisterTypeHandlers(reflect.TypeOf((*z.SessionId)(nil)).Elem(), sessImp, sessExp)
}
