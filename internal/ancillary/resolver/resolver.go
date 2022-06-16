package resolver

import "sync"

type resolver struct {
	sync.Mutex
	value   any
	handler func(...any) (any, error)
}

func (r *resolver) Run(args ...interface{}) (v interface{}, err error) {
	if r.value != nil {
		v = r.value
	} else {
		r.Lock()
		defer r.Unlock()
		r.value, err = r.handler(args)
		if err != nil {
			return
		}
		v = r.value
	}
	return
}
