package codec

import (
	i "github.com/go-serv/service/internal"
)

//
// The implementation of the codec middleware.
// Unmarshaler task: wire data -> t1 -> t2 -> unmarshaler -> message
// Marshaler task: message -> marshaler -> t2 -> t1 -> wire data

type codecMwGroup struct {
	codec          i.CodecInterface
	unmarshalChain []i.UnmarshalMwTaskHandler
	marshalChain   []i.MarshalMwTaskHandler
}

type task struct {
	mwGroup       *codecMwGroup
	df            i.DataFrameInterface
	methodReflect i.MethodReflectInterface
	msgRefect     i.MessageReflectInterface
	data          []byte
}

type unmarshalerTask struct {
	task
}

type marshalerTask struct {
	task
}

func (m *codecMwGroup) AddHandlers(un i.UnmarshalMwTaskHandler, marshal i.MarshalMwTaskHandler) {
	m.unmarshalChain = append(m.unmarshalChain, un)
	m.marshalChain = append(m.marshalChain, marshal)
}

//func (t *unmarshalerTask) Execute() ([]byte, error) {
//	var outer, inner, curr i.UnmarshalMwTaskHandler
//	stubCall := func(next i.UnmarshalMwTaskHandler, in []byte, md i.MethodDescriptorInterface, df i.DataFrameInterface) ([]byte, error) {
//		return in, nil
//	}
//	ch := append(t.mwGroup.unmarshalChain, stubCall)
//	l1 := len(ch)
//	curr = ch[l1-1]
//	//
//	for ii := l1 - 2; ii >= 0; ii-- {
//		outer, inner = ch[ii], curr
//		curr = func(next i.UnmarshalMwTaskHandler, in []byte, md i.MethodDescriptorInterface, df i.DataFrameInterface) ([]byte, error) {
//			return outer(inner, in, md, df)
//		}
//	}
//	return curr(inner, t.data, t.methodDesc, t.df)
//}
//
//func (t *marshalerTask) Execute() (out []byte, err error) {
//	var outer, inner, curr i.MarshalMwTaskHandler
//	stubCall := func(next i.MarshalMwTaskHandler, out []byte, md i.MethodDescriptorInterface, df i.DataFrameInterface) ([]byte, error) {
//		return out, nil
//	}
//	// Iterate over the post-processor tasks in reverse order
//	ch := append([]i.MarshalMwTaskHandler{stubCall}, t.mwGroup.marshalChain...)
//	l1 := len(ch)
//	curr = ch[0]
//	for ii := 1; ii < l1; ii++ {
//		outer, inner = ch[ii], curr
//		curr = func(next i.MarshalMwTaskHandler, out []byte, md i.MethodDescriptorInterface, df i.DataFrameInterface) ([]byte, error) {
//			return outer(inner, out, md, df)
//		}
//	}
//	out, err = curr(inner, t.data, t.methodDesc, t.df)
//	if err != nil {
//		return nil, err
//	}
//	//
//	t.df.WithPayload(out)
//	return t.df.Compose()
//}

func (t *unmarshalerTask) Execute() (out []byte, err error) {
	ch := t.mwGroup.unmarshalChain
	l1 := len(ch)
	//
	for ii := 0; ii < l1; ii++ {
		handler := ch[ii]
		t.data, err = handler(t.data, t.methodReflect, t.msgRefect, t.df)
		if err != nil {
			return
		}
	}
	out = t.data
	return
}

func (t *marshalerTask) Execute() ([]byte, error) {
	var err error
	ch := t.mwGroup.marshalChain
	l1 := len(ch)
	//
	for ii := l1 - 1; ii >= 0; ii-- {
		handler := ch[ii]
		t.data, err = handler(t.data, t.methodReflect, t.msgRefect, t.df)
		if err != nil {
			return nil, err
		}
	}
	t.df.WithPayload(t.data)
	return t.df.Compose(nil)
}
