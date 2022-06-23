package error

import (
	"fmt"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(severity z.ErrorSeverityLevel, code codes.Code, msg string, args ...any) *error {
	err := new(error)
	err.level = severity
	err.status = status.New(code, fmt.Sprintf(msg, args))
	return err
}
