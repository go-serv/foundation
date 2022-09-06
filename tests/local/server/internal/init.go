package internal

import (
	rt "github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/service/local"
)

func init() {
	svc := new(sample)
	svc.LocalServiceInterface = local.NewService(Name)
	rt.Runtime().RegisterLocalService(svc)
}
