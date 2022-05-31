package codec

import (
	in "github.com/go-serv/service/internal"
)

//
// The implementation of the codec middleware.
// Marshaler task: message -> marshaler -> t1 -> t2 -> wire data
// Unmarshaler task: wire data -> t2 -> t1 -> unmarshaler -> message

// Wrapper functions for unmarshal/marshal handlers

type codecMwGroup struct {
	codec          in.CodecInterface
	unmarshalChain []in.CodecMwTaskUnHandler
	marshalChain   []in.CodecMwTaskMarshalHandler
}

type task struct {
	codec         in.CodecInterface
	mwGroup       *codecMwGroup
	df            in.DataFrameInterface
	methodReflect in.MethodReflectInterface
	msgReflect    in.MessageReflectInterface
	data          []byte
}

type unmarshalerTask struct {
	task
}

type marshalerTask struct {
	task
}

func (m *codecMwGroup) AddHandlers(un in.CodecMwTaskUnHandler, marshal in.CodecMwTaskMarshalHandler) {
	m.unmarshalChain = append(m.unmarshalChain, un)
	m.marshalChain = append(m.marshalChain, marshal)
}

func (t *marshalerTask) Execute() (out []byte, err error) {
	var curr in.MwTaskChainElement
	headCall := func(next in.MwTaskChainElement, _ []byte, _ in.MethodReflectInterface, msgReflect in.MessageReflectInterface, _ in.DataFrameInterface) (out []byte, err error) {
		out, err = t.codec.PureMarshal(msgReflect.Value())
		if err != nil {
			return
		}
		_, err = next(out)
		return
	}
	// A stub to get rid of a check for the last element in the middleware chain
	stub := func(_ in.MwTaskChainElement, in []byte, _ in.MethodReflectInterface, _ in.MessageReflectInterface, _ in.DataFrameInterface) (out []byte, err error) {
		t.data = in
		return
	}
	// Build a middleware chain of the added handlers in direct order: first added handler will be called first.
	ch := append([]in.CodecMwTaskMarshalHandler{headCall}, t.mwGroup.marshalChain...)
	ch = append(ch, stub)
	l1 := len(ch)
	for i := l1 - 1; i >= 0; i-- {
		handler := ch[i]
		next := curr
		curr = func(in []byte) (in.MwTaskChainElement, error) {
			out, err = handler(next, in, t.methodReflect, t.msgReflect, t.df)
			if err != nil {
				return nil, err
			}
			return curr, nil
		}
	}
	//
	_, err = curr(t.data)
	if err != nil {
		return
	}
	//
	t.df.WithPayload(t.data)
	out, err = t.df.Compose(nil)
	return
}

func (t *unmarshalerTask) Execute() (err error) {
	var curr in.MwTaskChainElement
	tailCall := func(_ in.MwTaskChainElement, in []byte, _ in.MethodReflectInterface, msgReflect in.MessageReflectInterface, _ in.DataFrameInterface) (out []byte, err error) {
		err = t.codec.PureUnmarshal(in, msgReflect.Value())
		if err != nil {
			return
		}
		out = in
		return
	}
	//
	ch := append([]in.CodecMwTaskUnHandler{tailCall}, t.mwGroup.unmarshalChain...)
	l1 := len(ch)
	for i := 0; i < l1; i++ {
		handler := ch[i]
		next := curr
		curr = func(in []byte) (in.MwTaskChainElement, error) {
			_, err = handler(next, in, t.methodReflect, t.msgReflect, t.df)
			if err != nil {
				return nil, err
			}
			return curr, nil
		}
	}
	//
	_, err = curr(t.data)
	return
}
