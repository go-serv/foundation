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
	name        string
	clientProc  *msgproc
	serviceProc *msgproc
}

func (c *codec) Marshal(v interface{}) ([]byte, error) {
	var (
		err  error
		data []byte
		ok   bool
		m    proto.Message
		proc i.MessageProcessorInterface
		task i.MessageProcessTaskInterface
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
	// Check if data are being marshaled by a service client
	// to select the right message processor
	ok, err = runtime.Runtime().IsRequestMessage(m)
	if err != nil {
		return nil, err
	}
	switch ok {
	case true:
		proc = c.clientProc
	default:
		proc = c.serviceProc
	}
	// Invoke marshaler task handlers
	task, err = proc.NewMarshalTask(data, m)
	if err != nil {
		return nil, err
	}
	return task.Execute()
}

func (c *codec) Unmarshal(data []byte, v interface{}) error {
	var (
		err  error
		ok   bool
		m    proto.Message
		proc i.MessageProcessorInterface
		task i.MessageProcessTaskInterface
	)
	m, ok = v.(proto.Message)
	if !ok {
		return nil
	}
	// Check if data are being unmarshaled by a service client
	ok, err = runtime.Runtime().IsResponseMessage(m)
	if err != nil {
		return err
	}
	switch ok {
	case true:
		proc = c.clientProc
	default:
		proc = c.serviceProc
	}
	// Invoke unmarshaler task handlers
	task, err = proc.NewUnmarshalTask(data, m)
	if err != nil {
		return err
	}
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

func (c *codec) ServiceProcessor() i.MessageProcessorInterface {
	return c.serviceProc
}

func (c *codec) ClientProcessor() i.MessageProcessorInterface {
	return c.clientProc
}
