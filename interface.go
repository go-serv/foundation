package foundation

import (
	"github.com/mesh-master/foundation/pkg/z"
)

type ClientInterface interface {
	z.NetworkClientInterface
	ListServices() error
}
