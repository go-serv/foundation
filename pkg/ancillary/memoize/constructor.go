package memoize

import "sync"

func NewMemoizer(handler func(...any) (any, error)) *memoizer {
	m := new(memoizer)
	m.run = new(sync.Once)
	m.fn = handler
	return m
}
