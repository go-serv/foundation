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

func (s *meta) Copy(target z.MetaInterface) {
	src := s.dic.(*HttpDictionary)
	dst := target.Dictionary().(*HttpDictionary)
	dst.SessionId = src.SessionId
}

func (s *meta) Dictionary() interface{} {
	return s.dic
}

func (s *meta) Hydrate() error {
	return s.dic.Hydrate(s.dic)
}

func (s *meta) Dehydrate() (md metadata.MD, err error) {
	if s.data == nil {
		s.data = &metadata.MD{}
	}
	err = s.dic.Dehydrate(s.dic)
	if err != nil {
		return
	}
	return *s.data, nil
}

func (s *meta) uint64toBase64(v uint64) string {
	var data [8]byte
	binary.LittleEndian.PutUint64(data[:], v)
	return base64.StdEncoding.EncodeToString(data[:])
}

func (s *meta) base64ToUint64(v string) (out uint64, err error) {
	var data []byte
	data, err = base64.StdEncoding.DecodeString(v)
	if err != nil {
		return
	}
	err = binary.Read(bytes.NewReader(data), binary.LittleEndian, &out)
	return
}

func (s *meta) registerTypeHandlers(dic z.DictionaryInterface) {
	dic.RegisterTypeHandler(reflect.TypeOf(""), func(op z.DictionaryOp, name, alias string, rv reflect.Value) {
		switch op {
		case z.HydrateOp:
			v := s.data.Get(name)
			if len(v) > 0 {
				rv.SetString(v[0])
			}
		case z.DehydrateOp:
			v := rv.String()
			if len(v) > 0 {
				s.data.Set(name, v)
			}
		}
	})
	dic.RegisterTypeHandler(reflect.TypeOf(z.SessionId(0)), func(op z.DictionaryOp, name, alias string, rv reflect.Value) {
		switch op {
		case z.HydrateOp:
			v := s.data.Get(name)
			if len(v) > 0 {
				sId, _ := s.base64ToUint64(v[0])
				rv.SetUint(sId)
			}
		case z.DehydrateOp:
			v := rv.Uint()
			if v > 0 {
				id := s.uint64toBase64(v)
				s.data.Set(name, id)
			}
		}
	})
}
