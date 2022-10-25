package net

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc/metadata"
	"reflect"
)

type meta struct {
	data *metadata.MD
	dic  z.DictionaryInterface
}

func (m *meta) Copy(target z.MetaInterface) {
	src := m.dic.(*HttpDictionary)
	dst := target.Dictionary().(*HttpDictionary)
	dst.SessionId = src.SessionId
}

func (m *meta) Dictionary() interface{} {
	return m.dic
}

func (m *meta) Hydrate() error {
	return m.dic.Hydrate(m.dic)
}

func (m *meta) Dehydrate() (md metadata.MD, err error) {
	if m.data == nil {
		m.data = &metadata.MD{}
	}
	err = m.dic.Dehydrate(m.dic)
	if err != nil {
		return
	}
	return *m.data, nil
}

func (m *meta) uint64toBase64(v uint64) string {
	var data [8]byte
	binary.LittleEndian.PutUint64(data[:], v)
	return base64.StdEncoding.EncodeToString(data[:])
}

func (m *meta) base64ToUint64(v string) (out uint64, err error) {
	var data []byte
	data, err = base64.StdEncoding.DecodeString(v)
	if err != nil {
		return
	}
	err = binary.Read(bytes.NewReader(data), binary.LittleEndian, &out)
	return
}

func (m *meta) registerTypeHandlers(dic z.DictionaryInterface) {
	dic.RegisterTypeHandler((*Base64String)(nil), func(op z.DictionaryOp, name, alias string, v reflect.Value) {
		switch op {
		case z.HydrateOp:
			values := m.data.Get(name)
			if len(values) > 0 {
				base64Str := values[0]
				if decoded, err := base64.StdEncoding.DecodeString(base64Str); err != nil {
					v.SetBytes(decoded)
				}
			}
		case z.DehydrateOp:
			seq := v.Bytes()
			if len(seq) > 0 {
				encoded := base64.StdEncoding.EncodeToString(seq)
				m.data.Set(name, encoded)
			}
		}
	})
	dic.RegisterTypeHandler(reflect.TypeOf(""), func(op z.DictionaryOp, name, alias string, rv reflect.Value) {
		switch op {
		case z.HydrateOp:
			v := m.data.Get(name)
			if len(v) > 0 {
				rv.SetString(v[0])
			}
		case z.DehydrateOp:
			v := rv.String()
			if len(v) > 0 {
				m.data.Set(name, v)
			}
		}
	})
	dic.RegisterTypeHandler(reflect.TypeOf(z.SessionId(0)), func(op z.DictionaryOp, name, alias string, rv reflect.Value) {
		switch op {
		case z.HydrateOp:
			v := m.data.Get(name)
			if len(v) > 0 {
				sId, _ := m.base64ToUint64(v[0])
				rv.SetUint(sId)
			}
		case z.DehydrateOp:
			v := rv.Uint()
			if v > 0 {
				id := m.uint64toBase64(v)
				m.data.Set(name, id)
			}
		}
	})
}
