package client

import (
	"context"
	"github.com/go-serv/service/internal/autogen/proto/net"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/pkg/z"
	"github.com/go-serv/service/pkg/z/platform"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"io"
	"os"
)

var serviceName = protoreflect.FullName(proto.NetParcel_ServiceDesc.ServiceName)

type client struct {
	z.NetworkClientInterface
	stubs net.NetParcelClient
}

func (c *client) NewClient(cc grpc.ClientConnInterface) {
	c.stubs = net.NewNetParcelClient(cc)
}

func (c *client) TransferFileByPathname(path platform.Pathname, maxChunkSize int) (err error) {
	var (
		file *os.File
	)
	if file, err = os.OpenFile(path.String(), os.O_RDONLY, os.FileMode(0444)); err != nil {
		return
	}
	return c.TransferFile(file, maxChunkSize)
}

func (c *client) TransferFile(reader io.Reader, maxChunkSize int) (err error) {
	var (
		nRead, off int
	)
	buf := make([]byte, 0, maxChunkSize)
	for {
		nRead, err = reader.Read(buf)
		switch err {
		case io.EOF:
			break
		case nil:
			req := &proto.Ftp_FileChunk_Request{}
			req.Body = buf[0:nRead]
			req.Range = &proto.Ftp_FileRange{Start: int64(off), End: int64(off + nRead)}
			if _, err = c.FtpTransfer(context.Background(), req); err != nil {
				return
			}
			off += nRead
		default:
			return
		}
	}
	return
}
