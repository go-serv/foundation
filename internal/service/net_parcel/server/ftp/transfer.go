package ftp

import (
	"context"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/pkg/z"
	"github.com/go-serv/service/pkg/z/platform"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ftp FtpImpl) FtpTransfer(ctx context.Context, req *proto.Ftp_FileChunk_Request) (res *proto.Ftp_FileChunk_Response, err error) {
	var (
		start, end int64
		sess       z.SessionInterface
	)
	netCtx := ctx.(z.NetServerContextInterface)
	sess = netCtx.Session()
	transferCtx := sess.Context().(*ftpContext)
	//
	if transferCtx.state != TransferInProgressState {
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
	if !fr.isValid(chunk) || fr.intersects(item.chunksTransferred) {
		return nil, status.Error(codes.InvalidArgument, "invalid file range")
	}
	_, err = item.zfd.WriteAt(chunk, start)
	if err != nil {
		return
	}
	item.chunksTransferred = append(item.chunksTransferred, fr)
	// Check if file transfer has been completed.
	if fr.spans(item.info.Size, item.chunksTransferred) {
		// Call a post action handler if necessary.
		if item.info.GetPostAction() {
			path := platform.Pathname(item.zfd.Name())
			ext := path.MultiExt()
			if handler, has := ftp.PostActions[ext]; has {
				transferCtx.state = PostProcessingInProgressState
				//go func() {
				if err = handler(netCtx, path); err != nil {
					transferCtx.state = FailedState
				} else {
					item.completedFlag = true
				}
				//}()
			}
		} else {
			item.completedFlag = true
		}
	}
	res = &proto.Ftp_FileChunk_Response{}
	// Check if current transfer session has been completed.
	var nCompleted int
	for _, file := range transferCtx.files {
		if file.completedFlag {
			nCompleted++
		}
	}
	if nCompleted == len(transferCtx.files) {
		transferCtx.state = CompletedState
	}
	//
	res.State = transferCtx.state.toProtoState()
	return
}
