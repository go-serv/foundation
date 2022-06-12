package ftp

import (
	"context"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/internal/grpc/session"
	"github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"path/filepath"
	"strings"
)

func (FtpImpl) FtpNewSession(ctx context.Context, req *proto.Ftp_NewSession_Request) (res *proto.Ftp_NewSession_Response, err error) {
	var (
		sess   z.SessionInterface
		ftpDir string
	)
	platform := runtime.Runtime().Platform()
	netCtx := ctx.(z.NetServerContextInterface)
	sess = netCtx.Session()
	if sess == nil { // Create an insecure session
		lifetime := uint16(req.GetLifetime())
		if lifetime == 0 {
			return nil, status.Error(codes.InvalidArgument, "session lifetime must be specified")
		}
		sess = session.NewSession(lifetime)
		netCtx.WithSession(sess)
	}
	// If session context is not empty, then user is trying to call the method during the same session several times.
	// Such calls mut be rejected.
	sessCtx := sess.Context()
	if sessCtx != nil {
		return nil, status.Error(codes.FailedPrecondition, "session context is already set")
	}
	//
	if len(req.GetFiles()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no files to transfer")
	}
	// Check if target directory is writable
	ftpDir, err = netCtx.Server().Resolver().FtpRootDir(nil)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "FTP root directory is not set")
	}
	//availSpace := platform.
	_ = platform
	dirname := filepath.FromSlash(ftpDir + "/" + req.GetTargetDir())
	if strings.Contains(dirname, ".") || strings.Contains(dirname, "..") {
		return nil, status.Error(codes.InvalidArgument, "./ or ../ path notations are not allowed")
	}
	//
	ftpCtx := new(ftpContext)
	ftpCtx.state = PendingState
	ftpCtx.dirname = dirname
	ftpCtx.files = make(fileMap)
	res = &proto.Ftp_NewSession_Response{}
	for _, info := range req.GetFiles() {
		fd := &proto.Ftp_FileDescriptor{}
		uv := z.UniqueId(0).Generate()
		fd.Handle = uint64(uv)
		fd.Info = info
		res.Descriptors = append(res.Descriptors, fd)
		ftpCtx.files[fileHandle(uv)] = info
	}
	sess.WithContext(ftpCtx)
	return
}
