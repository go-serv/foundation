package codec

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/runtime"
	"google.golang.org/protobuf/proto"
)

var (
	MarshalOptions   = proto.MarshalOptions{}
	UnmarshalOptions = proto.UnmarshalOptions{}
)

type codec struct {
	name string
}

func clientMiddlewareGroupOnCond(m proto.Message, cond func() (bool, error)) (i.CodecMiddlewareGroupInterface, error) {
	var (
		ok     bool
		err    error
		client i.ClientInterface
		svc    i.ServiceInterface
	)
	ok, err = cond()
	if err != nil {
		return nil, err
	}
	switch ok {
	case true:
		client, err = runtime.Runtime().ClientByMessage(m)
		if err != nil {
			return nil, err
		}
		return client.CodecMiddlewareGroup(), nil
	default:
		svc, err = runtime.Runtime().ServiceByMessage(m)
		if err != nil {
			return nil, err
		}
		return svc.CodecMiddlewareGroup(), nil
	}
}

func (c *codec) PureMarshal(m proto.Message) ([]byte, error) {
	return MarshalOptions.Marshal(m)
}

func (c *codec) Marshal(v interface{}) ([]byte, error) {
	var (
		err  error
		ok   bool
		m    proto.Message
		task i.CodecMwMarshalTaskInterface
		mg   i.CodecMiddlewareGroupInterface
	)
	m, ok = v.(proto.Message)
	if !ok {
		return nil, nil
	}
	//
	mg, err = clientMiddlewareGroupOnCond(m, func() (bool, error) {
		return runtime.Runtime().IsRequestMessage(m)
	})
	if err != nil {
		return nil, err
	}
	//
	task, err = mg.NewMarshalTask(m)
	if err != nil {
		return nil, err
	}
	return task.Execute()
}

func (c *codec) PureUnmarshal(data []byte, m proto.Message) error {
	return UnmarshalOptions.Unmarshal(data, m)
}

func (c *codec) Unmarshal(data []byte, v interface{}) error {
	var (
		err  error
		ok   bool
		m    proto.Message
		mg   i.CodecMiddlewareGroupInterface
		task i.CodecMwUnmarshalTaskInterface
	)
	m, ok = v.(proto.Message)
	if !ok {
		return nil
	}
	//
	mg, err = clientMiddlewareGroupOnCond(m, func() (bool, error) {
		return runtime.Runtime().IsResponseMessage(m)
	})
	if err != nil {
		return err
	}
	task, err = mg.NewUnmarshalTask(data, m)
	if err != nil {
		return err
	}
	//
	return task.Execute()
}

func (c *codec) Name() string {
	return c.name
}

func (c *codec) NewDataFrame() i.DataFrameInterface {
	return NewDataFrame()
}
