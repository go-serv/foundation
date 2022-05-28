package internal

import (
	rt "github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/internal/service/local"
)

func init() {
	svc := new(sample)
	svc.LocalServiceInterface = local.NewService(Name)
	rt.Runtime().RegisterLocalService(svc)
}
