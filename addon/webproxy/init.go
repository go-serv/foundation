package sec_chan

import (
	"github.com/mesh-master/foundation/addon/sec_chan/internal/codec"
	"github.com/mesh-master/foundation/addon/sec_chan/y"
	"github.com/mesh-master/foundation/internal/autogen/net/sec_chan"
	"github.com/mesh-master/foundation/internal/runtime"
	"github.com/mesh-master/foundation/internal/service"
	"github.com/mesh-master/foundation/pkg/z"
	"github.com/mesh-master/foundation/pkg/z/event"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
)

func init() {
	handler := codec.MessageWrapperHandler()
	encoding.RegisterMessageWrapper(codec.Name, handler)
	encoding.RegisterCodec(codec.NewCodec())
	service.Reflection().AddProtoExtension(sec_chan.E_EncOff)
	runtime.Runtime().RegisterEventHandler(event.NetClientBeforeDial, func(args ...any) bool {
		arg := args[0].(event.NetClientBeforeDialArgs)
		if !arg.TlsEnabled {
			arg.Client.WithDialOption(grpc.WithDefaultCallOptions(grpc.ForceCodec(y.NewCodec())))
		}
		return false
	})
	runtime.Runtime().RegisterEventHandler(event.NetClientNewContext, func(args ...any) bool {
		ctx := args[0].(z.NetClientContextInterface)
		newCtx := encoding.ContextWithMessageWrapper(ctx.Interface(), handler)
		ctx.WithInterface(newCtx)
		return false
	})

}
