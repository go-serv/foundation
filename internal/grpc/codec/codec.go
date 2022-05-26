package codec

import (
	i "github.com/go-serv/service/internal"
	"google.golang.org/protobuf/proto"
)

var (
	MarshalOptions   = proto.MarshalOptions{}
	UnmarshalOptions = proto.UnmarshalOptions{}
)

type codec struct {
	name string
	proc *msgproc
}

func (c *codec) Marshal(v interface{}) ([]byte, error) {
	m, ok := v.(proto.Message)
	if !ok {
		return nil, nil
	}
	//
	data, err := MarshalOptions.Marshal(m)
	if err != nil {
		return nil, err
	}
	//
	task, err2 := c.proc.NewPostTask(data, m)
	if err2 != nil {
		return nil, err2
	}
	return task.Execute()
}

func (c *codec) Unmarshal(data []byte, v interface{}) error {
	m, ok := v.(proto.Message)
	if !ok {
		return nil
	}
	task, err := c.proc.NewPreTask(data, m)
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

func (c *codec) Processor() i.MessageProcessorInterface {
	return c.proc
}
