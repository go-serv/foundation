package client

import (
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/pkg/z/platform"
	"io"
)

type NetParcelClientInterface interface {
	SecureSession(*proto.Session_Request) (*proto.Session_Response, error)
	FtpNewSession(*proto.Ftp_NewSession_Request) (*proto.Ftp_NewSession_Response, error)
	FtpTransferFileByPathname(platform.Pathname) error
	FtpTransferFile(io.Reader) error
}
