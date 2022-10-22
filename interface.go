package foundation

import (
	"github.com/go-serv/foundation/pkg/z"
)

type ClientInterface interface {
	z.NetworkClientInterface
	ListServices() error
}
