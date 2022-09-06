package memoize

func NewMemoizer(handler func(...any) (any, error)) *memoizer {
	r := new(memoizer)
	r.fn = handler
	return r
}
