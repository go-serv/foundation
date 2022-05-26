package codec

import (
	"github.com/go-serv/service/internal/grpc/descriptor"
	"google.golang.org/protobuf/proto"
)

//
// The implementation of pre and post message processors.
// Pre-processor: wire data -> t1 -> t2 -> unmarshaler -> message
// Post-processor: message -> marshaler -> t2 -> t1 -> wire data

type TaskHandler func(next TaskHandler, in []byte, msg descriptor.MessageDescriptorInterface, df DataFrameInterface) ([]byte, error)

type msgproc struct {
	codec     CodecInterface
	preChain  []TaskHandler
	postChain []TaskHandler
}

type msgprocTask struct {
	proc *msgproc
	df   DataFrameInterface
	msg  descriptor.MessageDescriptorInterface
	data []byte
}

type msgprocPreTask struct {
	msgprocTask
}

type msgprocPostTask struct {
	msgprocTask
}

func (m *msgproc) NewPreTask(wire []byte, msg proto.Message) (*msgprocPreTask, error) {
	t := &msgprocPreTask{}
	t.proc = m
	// Parse incoming data frame
	t.df = m.codec.NewDataFrame()
	if err := t.df.Parse(wire); err != nil {
		return nil, err
	}
	t.data = t.df.Payload()
	t.msg = descriptor.NewMessageDescriptor(msg)
	return t, nil
}

func (m *msgproc) NewPostTask(wire []byte, msg proto.Message) (*msgprocPostTask, error) {
	t := &msgprocPostTask{}
	t.proc = m
	// Parse incoming data frame
	t.df = m.codec.NewDataFrame()
	t.df.AttachData(wire)
	t.msg = descriptor.NewMessageDescriptor(msg)
	return t, nil
}

func (m *msgproc) AddHandlers(pre TaskHandler, post TaskHandler) {
	m.preChain = append(m.preChain, pre)
	m.postChain = append(m.postChain, post)
}

func (t *msgprocPreTask) Execute() ([]byte, error) {
	var outer, inner, curr TaskHandler
	tailCall := func(next TaskHandler, in []byte, msg descriptor.MessageDescriptorInterface, df DataFrameInterface) ([]byte, error) {
		return nil, nil
	}
	ch := append(t.proc.preChain, tailCall)
	l1 := len(ch)
	curr = t.proc.preChain[l1-1]
	//
	for i := l1 - 2; i >= 0; i-- {
		outer, inner = t.proc.preChain[i], curr
		curr = func(next TaskHandler, in []byte, msg descriptor.MessageDescriptorInterface, df DataFrameInterface) ([]byte, error) {
			return outer(inner, in, msg, df)
		}
	}
	return curr(inner, t.data, t.msg, t.df)
}

func (t *msgprocPostTask) Execute() ([]byte, error) {
	var outer, inner, curr TaskHandler
	tailCall := func(next TaskHandler, out []byte, msg descriptor.MessageDescriptorInterface, df DataFrameInterface) ([]byte, error) {
		return nil, nil
	}
	// Iterate over the post-processor tasks in reverse order
	ch := append([]TaskHandler{tailCall}, t.proc.postChain...)
	l1 := len(ch)
	curr = t.proc.preChain[0]
	for i := 1; i < l1; i++ {
		outer, inner = t.proc.preChain[i], curr
		curr = func(next TaskHandler, out []byte, msg descriptor.MessageDescriptorInterface, df DataFrameInterface) ([]byte, error) {
			return outer(inner, out, msg, df)
		}
	}
	return curr(inner, t.data, t.msg, t.df)
}
