package session

import (
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
	"github.com/go-serv/service/pkg/z"
)

func serverReqHandler(next z.NetChainElementFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	srvCtx := ctx.(z.NetServerContextInterface)
	sess := srvCtx.Session()
	if sess != nil && sess.BlockCipher() != nil && srvCtx.InputReflection().Bool(go_serv.E_EncOff) != true {
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
	if sess != nil && sess.BlockCipher() != nil && srvCtx.OutputReflection().Bool(go_serv.E_EncOff) != true {
		df := srvCtx.Response().DataFrame()
		df.WithBlockCipher(sess.BlockCipher())
		return
	}
	return
}

func clientReqHandler(next z.NetChainElementFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	cipher := ctx.(z.NetClientContextInterface).Client().BlockCipher()
	if cipher != nil && ctx.InputReflection().Bool(go_serv.E_EncOff) != true {
		req.DataFrame().WithBlockCipher(cipher)
	}
	_, err = next(req, nil)
	return
}

func clientResHandler(next z.NetChainElementFn, ctx z.NetContextInterface, res z.ResponseInterface) (err error) {
	_, err = next(nil, res)
	cipher := ctx.(z.NetClientContextInterface).Client().BlockCipher()
	if cipher != nil && ctx.OutputReflection().Bool(go_serv.E_EncOff) != true {
		df := ctx.Response().DataFrame()
		df.WithBlockCipher(cipher)
		if err = df.Decrypt(); err != nil {
			return
		}
	}
	return
}
