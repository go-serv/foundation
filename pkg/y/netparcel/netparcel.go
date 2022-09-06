package netparcel

import (
	proto "github.com/go-serv/foundation/internal/autogen/proto/net"
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/pkg/z/platform"
)

type (
	SessionRequest         = proto.Session_Request
	SessionResponse        = proto.Session_Response
	FtpNewSessionRequest   = proto.Ftp_NewSession_Request
	FtpNewSessionResponse  = proto.Ftp_NewSession_Response
	FtpPostActionHandlerFn func(ctx z.NetServerContextInterface, pathname platform.Pathname) error
)

type NetParcelServiceInterface interface {
	RegisterFtpPostActionHandler(fn FtpPostActionHandlerFn, fileExt string)
}

type NetParcelClientInterface interface {
	SecureSession(*SessionRequest) (*SessionResponse, error)
	FtpNewSession(*FtpNewSessionRequest) (*FtpNewSessionResponse, error)
	FtpTransferDir(target platform.Pathname, recursive bool, temp bool) error
	FtpTransferFile(target platform.Pathname, temp bool, postAction bool) error
}
