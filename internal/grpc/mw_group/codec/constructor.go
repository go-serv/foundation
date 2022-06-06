package codec

import (
	"github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/protobuf/proto"
)

func NewCodecMiddlewareGroup(cc z.CodecInterface) *codecMwGroup {
	g := new(codecMwGroup)
	g.codec = cc
	return g
}

func (m *codecMwGroup) NewUnmarshalTask(wire []byte, msg proto.Message) (z.CodecMwUnmarshalTaskInterface, error) {
	t := &unmarshalerTask{}
	t.mwGroup = m
	// Parse incoming data frame
	t.df = m.codec.NewDataFrame()
	if err := t.df.Parse(wire, nil); err != nil {
		return nil, err
	}
	//
	ref := runtime.Runtime().Reflection()
	mReflect, err := ref.MethodReflectionFromMessage(msg)
	if err != nil {
		return nil, err
	}
	t.methodReflect = mReflect
	t.msgReflect = mReflect.FromMessage(msg)
	t.data = t.df.Payload()
	t.codec = m.codec
	return t, nil
}

func (m *codecMwGroup) NewMarshalTask(msg proto.Message) (z.CodecMwMarshalTaskInterface, error) {
	t := &marshalerTask{}
	t.mwGroup = m
	t.df = m.codec.NewDataFrame()
	//
	ref := runtime.Runtime().Reflection()
	mReflect, err := ref.MethodReflectionFromMessage(msg)
	if err != nil {
		return nil, err
	}
	t.methodReflect = mReflect
	t.msgReflect = mReflect.FromMessage(msg)
	t.codec = m.codec
	return t, nil
}
