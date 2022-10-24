package middleware

import (
	"errors"
	"github.com/go-serv/foundation/addon/sec_chan/internal/codec"
	"github.com/go-serv/foundation/addon/sec_chan/x"
	"github.com/go-serv/foundation/internal/autogen/foundation"
	"github.com/go-serv/foundation/internal/autogen/net/sec_chan"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var (
	ErrDf            = errors.New("not a dataframe")
	errEncryptionReq = status.Error(codes.FailedPrecondition, "gRPC call requires message to be encrypted")
)

func ServerReqHandler(next z.NextHandlerFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	var (
		sess z.SessionInterface
		ok   bool
		df   x.DataFrameInterface
	)
	srvCtx := ctx.(z.NetServerContextInterface)
	// If incoming message was encrypted by the client, an object implementing DataFrameInterface will be created by the
	// registered message wrapper.
	// The modified implementation of grpc-go provides MessageWrapperFromIncomingContext method to determine a wrapper to use.
	// See https://github.com/go-serv/grpc-go/blob/master/encoding/encoding.go for more details.
	if df, ok = srvCtx.Request().Data().(x.DataFrameInterface); ok {
		//srvCtx.ToggleMessageWrapperUse()
		encOff := req.MethodReflection().Bool(sec_chan.E_EncOff)
		if encOff {
			goto next
		}
		sess = srvCtx.Session()
		// Check preconditions.
		if sess == nil || sess.BlockCipher() != nil {
			if req.ServiceReflection().Bool(foundation.E_EnforceSecureChannel) {
				return errEncryptionReq
			} else {
				goto next
			}
		}
		// Decrypt message.
		df.WithBlockCipher(sess.BlockCipher())
		if err = df.Decrypt(); err != nil {
			return
		}
	}
next:
	_, err = next(req, nil)
	return
}

func ServerResHandler(next z.NextHandlerFn, ctx z.NetContextInterface, res z.ResponseInterface) (err error) {
	var (
		msg proto.Message
		ok  bool
		df  x.DataFrameInterface
	)
	srvCtx := ctx.(z.NetServerContextInterface)
	if msg, ok = srvCtx.Response().Data().(proto.Message); !ok {
		return
	}
	df = codec.NewDataFrame(msg)
	srvCtx.Response().WithData(df)

	if res.MethodReflection().Bool(sec_chan.E_EncOff) {
		return
	}

	sess := srvCtx.Session()
	if sess == nil || sess.BlockCipher() == nil {
		return errEncryptionReq
	}
	df.WithBlockCipher(sess.BlockCipher())
	_, err = next(nil, res)
	return
}
