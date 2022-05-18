package runtime

var rt *runtime

func init() {
	rt = new(runtime)
	rt.registeredNetServices = make(netServiceRegistry)
}
