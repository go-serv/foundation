package ftp

import (
	"context"
	"fmt"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/internal/grpc/session"
	"github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/pkg/z"
	"github.com/go-serv/service/pkg/z/platform"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
)

func transferTotalSize(r *proto.Ftp_NewSession_Request) uint64 {
	return 0
}

func validatePathnames(targetDir platform.Pathname, r *proto.Ftp_NewSession_Request) bool {
	for _, info := range r.GetFiles() {
		p := targetDir.ComposePath(info.GetPathname())
		if !p.IsCanonical() {
			return false
		}
	}
	return true
}

func (FtpImpl) FtpNewSession(ctx context.Context, req *proto.Ftp_NewSession_Request) (res *proto.Ftp_NewSession_Response, err error) {
	var (
		sess     z.SessionInterface
		profiles []z.FtpUploadProfileInterface
		zfd      platform.FileDescriptor
		pv       any
	)
	plat := runtime.Runtime().Platform()
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
	//
	if sess.State() != session.New {
		return nil, status.Error(codes.FailedPrecondition, "file transfer session is already in progress")
	}
	//
	if len(req.GetFiles()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no files to transfer")
	}
	//
	if pv, err = runtime.Runtime().Resolve(z.FtpUploadProfilerResolver); err != nil {
		return nil, status.Error(codes.FailedPrecondition, "no FTP upload profiles")
	}
	profiles = pv.([]z.FtpUploadProfileInterface)
	profileIdx := req.GetUploadProfile()
	if int(profileIdx) >= len(profiles) {
		return nil, status.Error(codes.FailedPrecondition, "profile index out of range")
	}
	profile := profiles[profileIdx]
	// Out of space check
	availableSpace := plat.AvailableDiskSpace(profile.BaseDir())
	if transferTotalSize(req) > availableSpace {
		return nil, status.Error(codes.FailedPrecondition, "out of disk space")
	}
	// Compose a pathname prefix of the target directory in the format /yyyy/mm/dd/hh
	now := time.Now()
	hour, day, month, year := now.Hour(), now.Day(), now.Month(), now.Year()
	dirname := profile.BaseDir().ComposePath(fmt.Sprintf("%d", year),
		fmt.Sprintf("%.2d", month),
		fmt.Sprintf("%.2d", day),
		fmt.Sprintf("%.2d", hour),
		platform.PathSeparator,
	)
	if !plat.DirectoryExists(dirname) {
		if err = plat.CreateDir(dirname, profile.FilePerms()); err != nil {
			return nil, status.Error(codes.FailedPrecondition, "fs: failed to create directory")
		}
	}
	postfix := strconv.FormatUint(uint64(sess.Id()), 16)
	// Add a marker to the directory pathname denoting that the directory is temporary and must be deleted
	// along with the transferred files, once the session is expired. It's up to a service to move transferred files to
	// another location.
	if req.GetTemp() {
		postfix += "-temp"
	}
	targetDir := dirname.ComposePath(postfix, platform.PathSeparator)
	if !targetDir.IsCanonical() {
		return nil, status.Error(codes.InvalidArgument, "meta characters in path names are not allowed")
	}
	//
	if err = plat.CreateDir(targetDir, 0755); err != nil {
		return nil, status.Error(codes.InvalidArgument, "mkdir")
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
		//
		zfpath := targetDir.ComposePath(info.GetPathname())
		zfd, err = plat.CreateZeroFile(zfpath, int64(info.Size), 0755)
		if err != nil {
			return nil, status.Error(codes.FailedPrecondition, "failed to create file")
		}
		ftpCtx.files[fileHandle(uv)] = &fileMapItem{
			info:        info,
			zfd:         zfd,
			transferred: make([]fileRange, 0, 1000), // with max chunk size of 4Mb must be enough for most cases
		}
	}
	sess.WithContext(ftpCtx)
	return
}
