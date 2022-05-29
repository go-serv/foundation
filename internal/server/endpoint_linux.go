package server

func (e *localEndpoint) Address() string {
	return "/tmp/." + e.pathname
}
