package memoize

import "sync"

type memoizer struct {
	sync.Once
	value any
	fn    func(...any) (any, error)
}

func (r *memoizer) Run(args ...any) (v any, err error) {
	if r.value != nil {
		v = r.value
	} else {
		r.Do(func() {
			r.value, err = r.fn(args)
		})
		v = r.value
	}
	return
}
