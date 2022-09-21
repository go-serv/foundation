package middleware

import (
	"github.com/go-serv/foundation/internal/autogen/go_serv/net/ext"
	"github.com/go-serv/foundation/pkg/z"
)

func serverReqHandler(next z.NetChainElementFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	srvCtx := ctx.(z.NetServerContextInterface)
	sess := srvCtx.Session()
	if sess != nil && sess.BlockCipher() != nil && req.MessageReflection().Bool(ext.E_EncOff) != true {
		df := srvCtx.Request().DataFrame()
		df.WithBlockCipher(sess.BlockCipher())
		if err = df.Decrypt(); err != nil {
			return
		}
	}
	_, err = next(req, nil)
	return
}

func serverResHandler(next z.NetChainElementFn, ctx z.NetContextInterface, res z.ResponseInterface) (err error) {
	_, err = next(nil, res)
	srvCtx := ctx.(z.NetServerContextInterface)
	sess := srvCtx.Session()
	if sess != nil && sess.BlockCipher() != nil && res.MessageReflection().Bool(ext.E_EncOff) != true {
		df := srvCtx.Response().DataFrame()
		df.WithBlockCipher(sess.BlockCipher())
		return
	}
	return
}

func clientReqHandler(next z.NetChainElementFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	cipher := ctx.(z.NetClientContextInterface).Client().BlockCipher()
	if cipher != nil && req.MessageReflection().Bool(ext.E_EncOff) != true {
		req.DataFrame().WithBlockCipher(cipher)
	}
	_, err = next(req, nil)
	return
}

func clientResHandler(next z.NetChainElementFn, ctx z.NetContextInterface, res z.ResponseInterface) (err error) {
	_, err = next(nil, res)
	cipher := ctx.(z.NetClientContextInterface).Client().BlockCipher()
	if cipher != nil && res.MessageReflection().Bool(ext.E_EncOff) != true {
		df := ctx.Response().DataFrame()
		df.WithBlockCipher(cipher)
		if err = df.Decrypt(); err != nil {
			return
		}
	}
	return
}
