package callctx

import (
	"context"
	"github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/protobuf/proto"
)

type CallCtx struct {
	context.Context
	methodRef    z.MethodReflectionInterface
	inputMsgRef  z.MessageReflectionInterface
	outputMsgRef z.MessageReflectionInterface
}

func (c *CallCtx) WithInput(msg proto.Message) (err error) {
	reflect := runtime.Runtime().Reflection()
	if c.methodRef == nil {
		if c.methodRef, err = reflect.MethodReflectionFromMessage(msg); err != nil {
			return
		}
	}
	c.inputMsgRef = c.methodRef.FromMessage(msg)
	return
}

func (c *CallCtx) WithOutput(msg proto.Message) (err error) {
	reflect := runtime.Runtime().Reflection()
	if c.methodRef == nil {
		if c.methodRef, err = reflect.MethodReflectionFromMessage(msg); err != nil {
			return
		}
	}
	c.outputMsgRef = c.methodRef.FromMessage(msg)
	return
}

func (c *CallCtx) MethodReflection() z.MethodReflectionInterface {
	return c.methodRef
}

func (c *CallCtx) InputReflection() z.MessageReflectionInterface {
	return c.inputMsgRef
}

func (c *CallCtx) OutputReflection() z.MessageReflectionInterface {
	return c.outputMsgRef
}
