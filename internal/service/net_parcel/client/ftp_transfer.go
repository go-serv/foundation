package client

import (
	"github.com/go-serv/service/internal/ancillary/struc/copyable"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/internal/grpc/call"
	"github.com/go-serv/service/pkg/z"
	"github.com/go-serv/service/pkg/z/platform"
	"io"
	"os"
)

type FtpTransferOptions struct {
	copyable.Shallow
	call.NetOptions
	c              *client
	MaxChunkSize   int
	BandwidthLimit int
}

func (f FtpTransferOptions) FtpTransferFileByPathname(path platform.Pathname) (err error) {
	var (
		file *os.File
	)
	if file, err = os.OpenFile(path.String(), os.O_RDONLY, os.FileMode(0444)); err != nil {
		return
	}
	return f.FtpTransferFile(file)
}

func (f FtpTransferOptions) FtpTransferFile(reader io.Reader) (err error) {
	var (
		nRead, off int
	)
	buf := make([]byte, 0, z.GrpcMaxMessageSize)
	for {
		nRead, err = reader.Read(buf)
		switch err {
		case io.EOF:
			break
		case nil:
			req := &proto.Ftp_FileChunk_Request{}
			req.Body = buf[0:nRead]
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
