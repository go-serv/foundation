package middleware

import (
	"errors"
	"github.com/go-serv/foundation/addon/sec_chan/internal/codec"
	"github.com/go-serv/foundation/internal/autogen/foundation"
	"github.com/go-serv/foundation/internal/autogen/net/sec_chan"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/protobuf/proto"
)

var (
	ErrDf = errors.New("not a dataframe")
)

func ServerReqHandler(next z.NextHandlerFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	var (
		sess z.SessionInterface
		ok   bool
		df   z.DataFrameInterface
	)
	srvCtx := ctx.(z.NetServerContextInterface)
	sess = srvCtx.Session()
	if sess != nil && sess.BlockCipher() != nil && req.MessageReflection().Bool(sec_chan.E_EncOff) != true {
		if df, ok = srvCtx.Request().Data().(z.DataFrameInterface); !ok {
			err = ErrDf
			return
		}
		df.WithBlockCipher(sess.BlockCipher())
		if err = df.Decrypt(); err != nil {
			return
		}
	}
	_, err = next(req, nil)
	return
}

func ServerResHandler(next z.NextHandlerFn, ctx z.NetContextInterface, res z.ResponseInterface) (err error) {
	var (
		msg proto.Message
		ok  bool
		df  z.DataFrameInterface
	)

	if _, err = next(nil, res); err != nil {
		return
	}

	srvCtx := ctx.(z.NetServerContextInterface)
	if msg, ok = srvCtx.Response().Data().(proto.Message); !ok {
		return
	}
	df = codec.NewDataFrame(msg)
	srvCtx.Response().WithData(df)

	sess := srvCtx.Session()
	if sess != nil && sess.BlockCipher() != nil && res.MessageReflection().Bool(sec_chan.E_EncOff) != true {
		df.WithBlockCipher(sess.BlockCipher())
		return
	}

	if res.MethodReflection().Bool(foundation.E_CloseSession) {
		sess.Close()
	}

	return
}

func ClientReqHandler(next z.NextHandlerFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	var (
		ok bool
		df z.DataFrameInterface
	)
	cipher := ctx.(z.NetClientContextInterface).Client().BlockCipher()
	if df, ok = req.Data().(z.DataFrameInterface); !ok {
		err = ErrDf
		return
	}
	if cipher != nil && req.MessageReflection().Bool(sec_chan.E_EncOff) != true {
		df.WithBlockCipher(cipher)
	}
	_, err = next(req, nil)
	return
}

func ClientResHandler(next z.NextHandlerFn, ctx z.NetContextInterface, res z.ResponseInterface) (err error) {
	var (
		ok bool
		df z.DataFrameInterface
	)

	if _, err = next(nil, res); err != nil {
		return
	}

	cipher := ctx.(z.NetClientContextInterface).Client().BlockCipher()
	if cipher != nil && res.MessageReflection().Bool(sec_chan.E_EncOff) != true {
		if df, ok = res.Data().(z.DataFrameInterface); !ok {
			err = ErrDf
			return
		}
		df.WithBlockCipher(cipher)
		if err = df.Decrypt(); err != nil {
			return
		}
	}
	return
}
