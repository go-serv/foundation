package error

import (
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc/status"
)

type error struct {
	level  z.ErrorSeverityLevel
	status *status.Status
}

func (e error) Error() string {
	return e.status.String()
}
