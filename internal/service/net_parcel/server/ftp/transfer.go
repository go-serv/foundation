package ftp

import (
	"context"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (FtpImpl) FtpTransfer(ctx context.Context, req *proto.Ftp_FileChunk_Request) (res *proto.Ftp_FileChunk_Response, err error) {
	var (
		start, end int64
		sess       z.SessionInterface
	)
	netCtx := ctx.(z.NetServerContextInterface)
	sess = netCtx.Session()
	transferCtx := sess.Context().(*ftpContext)
	//
	if transferCtx.state != InProgressState {
		return nil, status.Error(codes.InvalidArgument, "file transfer was terminated")
	}
	//
	h := fileHandle(req.GetFileHandle())
	if _, ok := transferCtx.files[h]; !ok {
		return nil, status.Error(codes.InvalidArgument, "invalid file handler")
	}
	//
	item := transferCtx.files[h]
	start = req.GetRange().GetStart()
	chunk := req.GetBody()
	if req.GetRange().End > 0 {
		end = req.GetRange().GetEnd()
	} else {
		end = start + int64(len(chunk))
	}
	//
	fr := fileRange{start, end}
	if !fr.isValid(chunk) || fr.intersects(item.transferred) {
		return nil, status.Error(codes.InvalidArgument, "invalid file range")
	}
	_, err = item.zfd.WriteAt(chunk, int64(start))
	if err != nil {
		return
	}
	// Check if file transferring is completed.
	if fr.spans(item.info.Size, item.transferred) {
		transferCtx.state = CompletedState
	}
	item.transferred = append(item.transferred, fr)
	//
	res = &proto.Ftp_FileChunk_Response{}
	res.State = transferCtx.state.toProtoState()
	return
}
