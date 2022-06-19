package netparcel

import (
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/pkg/z/platform"
	"io"
)

type (
	SessionRequest        = proto.Session_Request
	SessionResponse       = proto.Session_Response
	FtpNewSessionRequest  = proto.Ftp_NewSession_Request
	FtpNewSessionResponse = proto.Ftp_NewSession_Response
)

type NetParcelClientInterface interface {
	SecureSession(*SessionRequest) (*SessionResponse, error)
	FtpNewSession(*FtpNewSessionRequest) (*FtpNewSessionResponse, error)
	FtpTransferFileByPathname(platform.Pathname) error
	FtpTransferFile(io.Reader) error
}
