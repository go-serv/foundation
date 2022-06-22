package client

import (
	"github.com/go-serv/service/internal/ancillary/struc/copyable"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/internal/grpc/call"
	"github.com/go-serv/service/pkg/z"
	"github.com/go-serv/service/pkg/z/platform"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type FtpTransferOptions struct {
	copyable.Shallow
	call.NetOptions
	c            *client
	MaxChunkSize int
	BitrateLimit int
	Temp         bool
}

type transferDir struct {
	recursive bool
	req       *proto.Ftp_NewSession_Request
}

func (f transferDir) walkFn(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}
	switch m := info.Mode(); true {
	case m.IsDir() && f.recursive:
		return filepath.Walk(path, f.walkFn)
	case m.IsRegular():
		return f.handleRegularFile(path, info)
	}
	return err
}

func (f transferDir) handleRegularFile(path string, info fs.FileInfo) (err error) {
	f.req.Files = append(f.req.Files, &proto.Ftp_FileInfo{
		RelPath: path,
		Size:    info.Size(),
	})
	return err
}

func (f FtpTransferOptions) FtpTransferDir(target platform.Pathname, recursive bool, temp bool) (err error) {
	td := transferDir{
		recursive: recursive,
		req:       &proto.Ftp_NewSession_Request{},
	}
	if err = filepath.Walk(target.String(), td.walkFn); err != nil {
		return
	}
	//
	var (
		res *proto.Ftp_NewSession_Response
		fd  *os.File
	)
	td.req.Temp = temp
	if res, err = f.c.FtpNewSession(td.req); err != nil {
		return
	}
	for _, desc := range res.GetDescriptors() {
		path := target.ComposePath(desc.Info.GetRelPath())
		if fd, err = os.OpenFile(path.String(), os.O_RDONLY, os.FileMode(0444)); err != nil {
			return
		}
		if err = f.transferFile(fd, desc.Handle); err != nil {
			return
		}
	}
	return
}

func (f FtpTransferOptions) FtpTransferFile(path platform.Pathname) (err error) {
	var (
		file *os.File
		res  *proto.Ftp_NewSession_Response
	)
	if file, err = os.OpenFile(path.String(), os.O_RDWR, os.FileMode(0754)); err != nil {
		return
	}
	req := &proto.Ftp_NewSession_Request{}
	req.Temp = f.Temp
	if res, err = f.c.stubs.FtpNewSession(f.PrepareContext(), req); err != nil {
		return
	}
	return f.transferFile(file, res.GetDescriptors()[0].Handle)
}

func (f FtpTransferOptions) transferFile(reader io.Reader, fh uint64) (err error) {
	var (
		nRead, off int
	)
	buf := make([]byte, z.GrpcMaxMessageSize)
	for {
		nRead, err = reader.Read(buf)
		switch err {
		case io.EOF:
			break
		case nil:
			req := &proto.Ftp_FileChunk_Request{}
			req.Body = buf[0:nRead]
			req.FileHandle = fh
			req.Range = &proto.Ftp_FileRange{Start: int64(off), End: int64(off + nRead)}
			if _, err = f.c.stubs.FtpTransfer(f.PrepareContext(), req); err != nil {
				return
			}
			off += nRead
		default:
			return
		}
	}
	return
}
