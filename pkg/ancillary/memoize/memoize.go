package memoize

import "sync"

type memoizer struct {
	run   *sync.Once
	value any
	fn    func(...any) (any, error)
}

func (m *memoizer) Reset() {
	m.run = new(sync.Once)
	m.value = nil
}

func (m *memoizer) Run(args ...any) (v any, err error) {
	if m.value != nil {
		v = m.value
	} else {
		m.run.Do(func() {
			m.value, err = m.fn(args)
		})
		v = m.value
	}
	return
}
