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

func (c *codec) Marshal(v interface{}) ([]byte, error) {
	var (
		err  error
		data []byte
		ok   bool
		m    proto.Message
		task i.CodecMwTaskInterface
		mg   i.CodecMiddlewareGroupInterface
	)
	m, ok = v.(proto.Message)
	if !ok {
		return nil, nil
	}
	//
	data, err = MarshalOptions.Marshal(m)
	if err != nil {
		return nil, err
	}
	//
	mg, err = clientMiddlewareGroupOnCond(m, func() (bool, error) {
		return runtime.Runtime().IsRequestMessage(m)
	})
	if err != nil {
		return nil, err
	}
	task, err = mg.NewMarshalTask(data, m)
	//
	if err != nil {
		return nil, err
	}
	// Invoke marshaler task handlers
	return task.Execute()
}

func (c *codec) Unmarshal(data []byte, v interface{}) error {
	var (
		err  error
		ok   bool
		m    proto.Message
		mg   i.CodecMiddlewareGroupInterface
		task i.CodecMwTaskInterface
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
	data, err = task.Execute()
	if err != nil {
		return err
	}
	return UnmarshalOptions.Unmarshal(data, m)
}

func (c *codec) Name() string {
	return c.name
}

func (c *codec) NewDataFrame() i.DataFrameInterface {
	return NewDataFrame()
}
