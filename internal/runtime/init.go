package runtime

var rt *runtime

func init() {
	rt = new(runtime)
	//
	rt.netServices = make(registry, 0)
	rt.netClients = make(registry, 0)
	//
	rt.localService = make(registry, 0)
	rt.localClients = make(registry, 0)
}
