package resolver

func NewResolver(handler func(...any) (any, error)) *resolver {
	r := new(resolver)
	r.handler = handler
	return r
}
