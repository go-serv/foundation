package codec

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/runtime"
	"google.golang.org/protobuf/proto"
)

func NewCodecMiddlewareGroup(cc i.CodecInterface) *codecMwGroup {
	g := new(codecMwGroup)
	g.codec = cc
	return g
}

func (m *codecMwGroup) NewUnmarshalTask(wire []byte, msg proto.Message) (i.CodecMwTaskInterface, error) {
	t := &unmarshalerTask{}
	t.mwGroup = m
	// Parse incoming data frame
	t.df = m.codec.NewDataFrame()
	if err := t.df.Parse(wire); err != nil {
		return nil, err
	}
	//
	ref := runtime.Runtime().Reflection()
	mReflect, err := ref.MethodReflectionFromMessage(msg)
	if err != nil {
		return nil, err
	}
	t.methodReflect = mReflect
	t.msgRefect = mReflect.FromMessage(msg)
	t.data = t.df.Payload()
	return t, nil
}

func (m *codecMwGroup) NewMarshalTask(wire []byte, msg proto.Message) (i.CodecMwTaskInterface, error) {
	t := &marshalerTask{}
	t.mwGroup = m
	// Parse incoming data frame
	t.df = m.codec.NewDataFrame()
	//
	ref := runtime.Runtime().Reflection()
	mReflect, err := ref.MethodReflectionFromMessage(msg)
	if err != nil {
		return nil, err
	}
	t.methodReflect = mReflect
	t.msgRefect = mReflect.FromMessage(msg)
	t.data = wire
	return t, nil
}
