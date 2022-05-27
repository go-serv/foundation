package codec

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/runtime"
	"google.golang.org/protobuf/proto"
)

//
// The implementation of pre and post message processors.
// Pre-processor: wire data -> t1 -> t2 -> unmarshaler -> message
// Post-processor: message -> marshaler -> t2 -> t1 -> wire data

type msgproc struct {
	codec     i.CodecInterface
	preChain  []i.MsgProcTaskHandler
	postChain []i.MsgProcTaskHandler
}

type msgprocTask struct {
	proc       *msgproc
	df         i.DataFrameInterface
	methodDesc i.MethodDescriptorInterface
	data       []byte
}

type unmarshalerTask struct {
	msgprocTask
}

type marshalerTask struct {
	msgprocTask
}

func (m *msgproc) NewUnmarshalTask(wire []byte, msg proto.Message) (i.MessageProcessTaskInterface, error) {
	t := &unmarshalerTask{}
	t.proc = m
	// Parse incoming data frame
	t.df = m.codec.NewDataFrame()
	if err := t.df.Parse(wire); err != nil {
		return nil, err
	}
	//
	md, err := runtime.Runtime().MethodDescriptorByMessage(msg)
	if err != nil {
		return nil, err
	}
	t.methodDesc = md
	//
	t.data = t.df.Payload()
	return t, nil
}

func (m *msgproc) NewMarshalTask(wire []byte, msg proto.Message) (i.MessageProcessTaskInterface, error) {
	t := &marshalerTask{}
	t.proc = m
	// Parse incoming data frame
	t.df = m.codec.NewDataFrame()
	//
	md, err := runtime.Runtime().MethodDescriptorByMessage(msg)
	if err != nil {
		return nil, err
	}
	t.methodDesc = md
	//
	t.data = wire
	return t, nil
}

func (m *msgproc) AddHandlers(pre i.MsgProcTaskHandler, post i.MsgProcTaskHandler) {
	m.preChain = append(m.preChain, pre)
	m.postChain = append(m.postChain, post)
}

func (t *unmarshalerTask) Execute() ([]byte, error) {
	var outer, inner, curr i.MsgProcTaskHandler
	tailCall := func(next i.MsgProcTaskHandler, in []byte, md i.MethodDescriptorInterface, df i.DataFrameInterface) ([]byte, error) {
		return in, nil
	}
	ch := append(t.proc.preChain, tailCall)
	l1 := len(ch)
	curr = ch[l1-1]
	//
	for ii := l1 - 2; ii >= 0; ii-- {
		outer, inner = ch[ii], curr
		curr = func(next i.MsgProcTaskHandler, in []byte, md i.MethodDescriptorInterface, df i.DataFrameInterface) ([]byte, error) {
			return outer(inner, in, md, df)
		}
	}
	return curr(inner, t.data, t.methodDesc, t.df)
}

func (t *marshalerTask) Execute() (out []byte, err error) {
	var outer, inner, curr i.MsgProcTaskHandler
	tailCall := func(next i.MsgProcTaskHandler, out []byte, md i.MethodDescriptorInterface, df i.DataFrameInterface) ([]byte, error) {
		return out, nil
	}
	// Iterate over the post-processor tasks in reverse order
	ch := append([]i.MsgProcTaskHandler{tailCall}, t.proc.postChain...)
	l1 := len(ch)
	curr = ch[0]
	for ii := 1; ii < l1; ii++ {
		outer, inner = ch[ii], curr
		curr = func(next i.MsgProcTaskHandler, out []byte, md i.MethodDescriptorInterface, df i.DataFrameInterface) ([]byte, error) {
			return outer(inner, out, md, df)
		}
	}
	out, err = curr(inner, t.data, t.methodDesc, t.df)
	if err != nil {
		return nil, err
	}
	//
	t.df.WithPayload(out)
	return t.df.Compose()
}
