package ftp

import (
	"context"
	proto "github.com/go-serv/foundation/internal/autogen/foundation"
	grpc_err "github.com/go-serv/foundation/internal/grpc/error"
	"github.com/go-serv/foundation/internal/grpc/session"
	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/pkg/z/platform"
	"google.golang.org/grpc/codes"
	"os"
	"strconv"
)

func transferTotalSize(r *proto.Ftp_NewSession_Request) uint64 {
	return 0
}

func (FtpImpl) FtpNewSession(ctx context.Context, req *proto.Ftp_NewSession_Request) (res *proto.Ftp_NewSession_Response, err error) {
	var (
		sess     z.SessionInterface
		profiles []z.FtpUploadProfileInterface
		pv       any
	)
	plat := runtime.Runtime().Platform()
	netCtx := ctx.(z.NetServerContextInterface)
	sess = netCtx.Session()
	if sess == nil { // Create an insecure session
		lifetime := uint16(req.GetLifetime())
		if lifetime == 0 {
			return nil, grpc_err.New(z.ErrSeverityLow, codes.InvalidArgument, "session lifetime must be specified")
		}
		sess = session.NewSession(lifetime)
		netCtx.WithSession(sess)
	}
	if sess.State() != session.New {
		return nil, grpc_err.New(z.ErrSeverityLow, codes.FailedPrecondition, "file transfer session is already in progress")
	}
	if len(req.GetFiles()) == 0 {
		return nil, grpc_err.New(z.ErrSeverityLow, codes.InvalidArgument, "no files to transfer")
	}
	if pv, err = runtime.Runtime().Resolve(z.FtpUploadProfilerResolver); err != nil {
		return nil, grpc_err.New(z.ErrSeverityLow, codes.FailedPrecondition, "no FTP upload profiles")
	}
	profiles = pv.([]z.FtpUploadProfileInterface)
	profileIdx := req.GetUploadProfile()
	if int(profileIdx) >= len(profiles) {
		return nil, grpc_err.New(z.ErrSeverityLow, codes.FailedPrecondition, "profile index out of range")
	}
	profile := profiles[profileIdx]
	// Out of space check
	availableSpace := plat.AvailableDiskSpace(profile.RootDir())
	if transferTotalSize(req) > availableSpace {
		return nil, grpc_err.New(z.ErrSeverityCritical, codes.FailedPrecondition, "out of disk space")
	}
	//
	dirname := profile.RootDir()
	if !dirname.DirExists() {
		if err = plat.CreateDir(dirname, profile.FilePerms()); err != nil {
			return nil, grpc_err.New(z.ErrSeverityHigh, codes.FailedPrecondition, err.Error())
		}
	}
	postfix := strconv.FormatUint(uint64(sess.Id()), 16)
	// Add a marker to the directory pathname denoting that the directory is temporary and must be deleted
	// along with the chunksTransferred files, once the session is expired. It's up to a service to move transferred files to
	// another location.
	if req.GetTemp() {
		postfix += "-temp"
	}
	targetDir := dirname.ComposePath(postfix, platform.PathSeparator)
	if !targetDir.IsCanonical() {
		return nil, grpc_err.New(z.ErrSeverityLow, codes.InvalidArgument, "meta characters in path names are not allowed")
	}
	//
	if err = plat.CreateDir(targetDir, profile.FilePerms()); err != nil {
		return nil, grpc_err.New(z.ErrSeverityHigh, codes.FailedPrecondition, err.Error())
	}
	//
	ftpCtx := new(ftpContext)
	ftpCtx.state = PendingState
	ftpCtx.files = make(fileMap)
	res = &proto.Ftp_NewSession_Response{}
	for _, info := range req.GetFiles() {
		fd := &proto.Ftp_FileDescriptor{}
		uv := z.UniqueId(0).Generate()
		fd.Handle = uint64(uv)
		fd.Info = info
		res.Descriptors = append(res.Descriptors, fd)
		// Create relative path directories if necessary
		var zfpath, relpath platform.Pathname
		relpath = platform.Pathname(info.GetRelPath()).Normalize()
		if !relpath.IsFilename() {
			zfbase := targetDir.ComposePath(relpath.Dirname().String())
			if !zfbase.DirExists() {
				if err = plat.CreateDir(zfbase, profile.FilePerms()); err != nil {
					return
				}
			}
		}
		//
		zfpath = targetDir.ComposePath(relpath.String())
		var zfd *os.File
		if zfd, err = plat.CreateZeroFile(zfpath, info.Size, profile.FilePerms()); err != nil {
			return nil, grpc_err.New(z.ErrSeverityHigh, codes.FailedPrecondition, err.Error())
		}
		ftpCtx.files[fileHandle(uv)] = &fileMapItem{
			info:              info,
			zfd:               zfd,
			chunksTransferred: fileRange{}.New(),
		}
	}
	ftpCtx.state = TransferInProgressState
	sess.WithContext(ftpCtx)
	return
}
