package memoize

import "sync"

type memoizer struct {
	sync.Once
	value any
	fn    func(...any) (any, error)
}

func (r *memoizer) Run(args ...interface{}) (v interface{}, err error) {
	if r.value != nil {
		v = r.value
	} else {
		r.Do(func() {
			r.value, err = r.fn(args)
		})
		if err != nil {
			return
		}
		v = r.value
	}
	return
}
