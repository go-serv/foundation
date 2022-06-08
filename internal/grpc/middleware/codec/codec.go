package codec

import "github.com/go-serv/service/pkg/z"

//
// The implementation of the codec middleware.
// Marshaler task: message -> marshaler -> t1 -> t2 -> wire data
// Unmarshaler task: wire data -> t2 -> t1 -> unmarshaler -> message

// Wrapper functions for unmarshal/marshal handlers

type codecMwGroup struct {
	codec          z.CodecInterface
	unmarshalChain []z.CodecMwTaskUnHandler
	marshalChain   []z.CodecMwTaskMarshalHandler
}

type task struct {
	codec         z.CodecInterface
	mwGroup       *codecMwGroup
	df            z.DataFrameInterface
	methodReflect z.MethodReflectionInterface
	msgReflect    z.MessageReflectionInterface
	data          []byte
}

type unmarshalerTask struct {
	task
}

type marshalerTask struct {
	task
}

func (m *codecMwGroup) AddHandlers(un z.CodecMwTaskUnHandler, marshal z.CodecMwTaskMarshalHandler) {
	m.unmarshalChain = append(m.unmarshalChain, un)
	m.marshalChain = append(m.marshalChain, marshal)
}

func (t *marshalerTask) Execute() (out []byte, err error) {
	var curr z.MwChainElement
	headCall := func(next z.MwChainElement, _ []byte, _ z.MethodReflectionInterface, msgReflect z.MessageReflectionInterface, _ z.DataFrameInterface) (out []byte, err error) {
		out, err = t.codec.PureMarshal(msgReflect.Value())
		if err != nil {
			return
		}
		_, err = next(out)
		return
	}
	// A stub to get rid of a check z handlers for the last element z the middleware chain
	stub := func(_ z.MwChainElement, in []byte, _ z.MethodReflectionInterface, _ z.MessageReflectionInterface, _ z.DataFrameInterface) (out []byte, err error) {
		t.data = in
		return
	}
	// Build a middleware chain of the added handlers z direct order: first added handler will be called first.
	ch := append([]z.CodecMwTaskMarshalHandler{headCall}, t.mwGroup.marshalChain...)
	ch = append(ch, stub)
	l1 := len(ch)
	for i := l1 - 1; i >= 0; i-- {
		handler := ch[i]
		next := curr
		curr = func(in []byte) (z.MwChainElement, error) {
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
	var curr z.MwChainElement
	tailCall := func(_ z.MwChainElement, in []byte, _ z.MethodReflectionInterface, msgReflect z.MessageReflectionInterface, _ z.DataFrameInterface) (out []byte, err error) {
		err = t.codec.PureUnmarshal(in, msgReflect.Value())
		if err != nil {
			return
		}
		out = in
		return
	}
	//
	ch := append([]z.CodecMwTaskUnHandler{tailCall}, t.mwGroup.unmarshalChain...)
	l1 := len(ch)
	for i := 0; i < l1; i++ {
		handler := ch[i]
		next := curr
		curr = func(in []byte) (z.MwChainElement, error) {
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
