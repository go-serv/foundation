package ftp

import (
	proto "github.com/mesh-master/foundation/internal/autogen/foundation"
	"github.com/mesh-master/foundation/pkg/y/netparcel"
	"github.com/mesh-master/foundation/pkg/z"
	"github.com/mesh-master/foundation/pkg/z/platform"
	"os"
)

type FtpImpl struct {
	PostActions postActionMap
}

type (
	rootDirPostfixFn func() platform.Pathname
	ftpState         int
	fileHandle       z.UniqueId
	postActionMap    map[string]netparcel.FtpPostActionHandlerFn
)

const (
	PendingState ftpState = iota + 1
	TransferInProgressState
	PostProcessingInProgressState
	SuspendedState
	CompletedState
	FailedState
)

func (s ftpState) toProtoState() proto.Ftp_TransferState {
	switch s {
	case PendingState:
		return proto.Ftp_PENDING
	case TransferInProgressState:
		return proto.Ftp_TRANSFER_IN_PROGRESS
	case PostProcessingInProgressState:
		return proto.Ftp_POST_PROCESSING_IN_PROGRESS
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

func (fileRange) New() []fileRange {
	return make([]fileRange, 0, 1000) // with the max chunk size of 4Mb must be enough for most cases
}

func (fr fileRange) Span() int64 {
	return fr.end - fr.start
}

func (fr fileRange) isValid(chunk []byte) bool {
	if fr.start > fr.end || fr.Span() != int64(len(chunk)) {
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
	var totalSpan int64
	for i := 0; i < len(ranges); i++ {
		totalSpan += ranges[i].Span()
	}
	return totalSpan == fileSize
}

type fileMapItem struct {
	info              *proto.Ftp_FileInfo
	zfd               *os.File
	chunksTransferred []fileRange
	completedFlag     bool
}

type fileMap map[fileHandle]*fileMapItem

type ftpContext struct {
	state ftpState
	files fileMap
}

type uploadProfile struct {
	maxFileSize      int64
	filePerms        platform.UnixPerms
	rootDir          platform.Pathname
	rootDirPostfixFn rootDirPostfixFn
}

func (prof uploadProfile) RootDir() platform.Pathname {
	return prof.rootDir.ComposePath(prof.rootDirPostfixFn().String())
}

func (prof uploadProfile) MaxFileSize() int64 {
	return prof.maxFileSize
}

func (prof uploadProfile) FilePerms() platform.UnixPerms {
	return prof.filePerms
}
