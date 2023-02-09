package error

import (
	"fmt"
	"github.com/mesh-master/foundation/pkg/z"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func New(severity z.ErrorSeverityLevel, code codes.Code, msg string, args ...any) *error {
	err := new(error)
	err.level = severity
	if len(args) > 0 {
		err.status = status.New(code, fmt.Sprintf(msg, args))
	} else {
		err.status = status.New(code, msg)
	}
	err.createdAt = time.Now().Unix()
	return err
}
