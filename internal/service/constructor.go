package service

func newBaseService(name string) baseService {
	s := baseService{name: name}
	s.state = StateInit
	//s.grpcServersJob = job.NewJob(nil)
	return s
}

func NewLocalService(name string) *localService {
	s := &localService{newBaseService(name)}
	return s
}

func NewNetworkService(name string) NetworkServiceInterface {
	s := &networkService{newBaseService(name)}
	return s
}
