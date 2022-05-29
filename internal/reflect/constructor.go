package reflect

func NewReflection() *reflect {
	ref := new(reflect)
	ref.services = make(serviceMapType, 0)
	return ref
}
