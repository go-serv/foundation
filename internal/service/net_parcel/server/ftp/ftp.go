package ftp

import (
	"github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/pkg/z"
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

type fileMap map[fileHandle]*net.Ftp_FileInfo

type ftpContext struct {
	state   ftpState
	files   fileMap
	dirname string
}
