package client

import (
	"context"
	"github.com/mesh-master/foundation/internal/grpc/callctx/net"
	"time"
)

type Options struct {
	cancelFn  context.CancelFunc
	TimeoutMs int
}

type NetOptions struct {
	Options
}

func (o *Options) PrepareContext() (ctx context.Context) {
	if o.TimeoutMs > 0 {
		clientDeadline := time.Now().Add(time.Duration(o.TimeoutMs) * time.Millisecond)
		ctx, o.cancelFn = context.WithDeadline(context.Background(), clientDeadline)
	} else {
		ctx = context.Background()
	}
	return
}

func (o *NetOptions) PrepareContext() context.Context {
	ctx := o.Options.PrepareContext()
	return net.NewClientContext(ctx)
}
