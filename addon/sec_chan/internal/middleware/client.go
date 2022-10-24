package middleware

import (
	"github.com/go-serv/foundation/addon/sec_chan/internal/codec"
	"github.com/go-serv/foundation/addon/sec_chan/x"
	"github.com/go-serv/foundation/internal/autogen/net/sec_chan"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/protobuf/proto"
)

func ClientReqHandler(next z.NextHandlerFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	var (
		msg proto.Message
		df  x.DataFrameInterface
	)
	client := ctx.(z.NetClientContextInterface).Client()
	hasTls := client.Endpoint().(z.NetEndpointInterface).TlsConfig() != nil
	if !hasTls {
		msg = req.Data().(proto.Message)
		df = codec.NewDataFrame(msg)
		encOff := req.MethodReflection().Bool(sec_chan.E_EncOff)
		if !encOff {
			cipher := ctx.(z.NetClientContextInterface).Client().BlockCipher()
			df.WithBlockCipher(cipher)
		}
		req.WithData(df)
	}
	_, err = next(req, nil)
	return
}

func ClientResHandler(next z.NextHandlerFn, ctx z.NetContextInterface, res z.ResponseInterface) (err error) {
	var (
		df x.DataFrameInterface
	)
	client := ctx.(z.NetClientContextInterface).Client()
	hasTls := client.Endpoint().(z.NetEndpointInterface).TlsConfig() != nil
	if !hasTls {
		encOff := res.MethodReflection().Bool(sec_chan.E_EncOff)
		if !encOff {
			cipher := ctx.(z.NetClientContextInterface).Client().BlockCipher()
			df = res.Data().(x.DataFrameInterface)
			df.WithBlockCipher(cipher)
			if err = df.Decrypt(); err != nil {
				return
			}
		}
	}
	_, err = next(nil, res)
	return
}
