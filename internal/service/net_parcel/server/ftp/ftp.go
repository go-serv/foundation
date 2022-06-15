package ftp

import (
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/pkg/z"
	"github.com/go-serv/service/pkg/z/platform"
)

type FtpImpl struct{}

type (
	ftpState   int
	fileHandle z.UniqueId
)

const (
	PendingState ftpState = iota + 1
	InProgressState
	SuspendedState
	CompletedState
	FailedState
)

func (s ftpState) toProtoState() proto.Ftp_TransferState {
	switch s {
	case PendingState:
		return proto.Ftp_PENDING
	case InProgressState:
		return proto.Ftp_IN_PROGRESS
	case SuspendedState:
		return proto.Ftp_SUSPENDED
	case CompletedState:
		return proto.Ftp_COMPLETED
	case FailedState:
		return proto.Ftp_FAILED
	default:
		panic("unknown FTP state")
	}
}

type fileRange struct {
	start, end int64
}

func (fr fileRange) isValid(data []byte) bool {
	if fr.start > fr.end || (fr.end-fr.start) != int64(len(data)) {
		return false
	} else {
		return true
	}
}

func (fr fileRange) intersects(ranges []fileRange) bool {
	for i := 0; i < len(ranges); i++ {
		switch true {
		case fr.start >= ranges[i].start && fr.start <= ranges[i].end:
			return true
		case fr.end >= ranges[i].start && fr.end <= ranges[i].end:
			return true
		}
	}
	return false
}

func (fr fileRange) spans(fileSize int64, ranges []fileRange) bool {
	var (
		totalSpan int64
	)
	r := append(ranges, fr)
	for i := 0; i < len(r); i++ {
		totalSpan += r[i].end - r[i].start
	}
	return totalSpan == fileSize
}

type fileMapItem struct {
	info        *proto.Ftp_FileInfo
	zfd         platform.FileDescriptor
	transferred []fileRange
}

type fileMap map[fileHandle]*fileMapItem

type ftpContext struct {
	state ftpState
	files fileMap
}
