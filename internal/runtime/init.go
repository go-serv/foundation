package runtime

import (
	"github.com/go-serv/foundation/internal/platform"
	"github.com/go-serv/foundation/pkg/z"
)

var rt *runtime

func init() {
	rt = new(runtime)
	rt.platform = platform.NewPlatform(0)
	rt.services = make(map[string]z.ServiceInterface)
	rt.clients = make(map[string]z.ClientInterface)
	rt.resolvers = make(resolversMapTyp, 0)
	rt.eventsMap = make(eventsMapTyp, 0)
}
