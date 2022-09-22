package middleware

import (
	"errors"
	"github.com/go-serv/foundation/addon/sec-chan/internal/codec"
	"github.com/go-serv/foundation/internal/autogen/go_serv/net/ext"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/protobuf/proto"
)

var (
	ErrDf = errors.New("not a dataframe")
)

func serverReqHandler(next z.MiddlewareChainElementFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	var (
		ok bool
		df z.DataFrameInterface
	)
	srvCtx := ctx.(z.NetServerContextInterface)
	sess := srvCtx.Session()
	if sess != nil && sess.BlockCipher() != nil && req.MessageReflection().Bool(ext.E_EncOff) != true {
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

func serverResHandler(next z.MiddlewareChainElementFn, ctx z.NetContextInterface, res z.ResponseInterface) (err error) {
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
	if sess != nil && sess.BlockCipher() != nil && res.MessageReflection().Bool(ext.E_EncOff) != true {
		df.WithBlockCipher(sess.BlockCipher())
		return
	}
	return
}

func clientReqHandler(next z.MiddlewareChainElementFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	var (
		ok bool
		df z.DataFrameInterface
	)
	cipher := ctx.(z.NetClientContextInterface).Client().BlockCipher()
	if df, ok = req.Data().(z.DataFrameInterface); !ok {
		err = ErrDf
		return
	}
	if cipher != nil && req.MessageReflection().Bool(ext.E_EncOff) != true {
		df.WithBlockCipher(cipher)
	}
	_, err = next(req, nil)
	return
}

func clientResHandler(next z.MiddlewareChainElementFn, ctx z.NetContextInterface, res z.ResponseInterface) (err error) {
	var (
		ok bool
		df z.DataFrameInterface
	)

	if _, err = next(nil, res); err != nil {
		return
	}

	cipher := ctx.(z.NetClientContextInterface).Client().BlockCipher()
	if cipher != nil && res.MessageReflection().Bool(ext.E_EncOff) != true {
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
