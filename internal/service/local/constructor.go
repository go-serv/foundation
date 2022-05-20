package local

import "github.com/go-serv/service/internal/service"

func NewLocalService(name string) *localService {
	s := &localService{service.NewBaseService(name)}
	return s
}
