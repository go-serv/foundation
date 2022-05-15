package service

import (
	"context"
	net_svc "github.com/go-serv/service/internal/autogen/proto/net"
)

type netParcel struct {
	net_svc.NetParcelServer
}

func (s *netParcel) GetCryptoNonce(context.Context, *net_svc.CryptoNonce_Request) (*net_svc.CryptoNonce_Response, error) {
	return nil, nil
}
