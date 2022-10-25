package memoize

type MemoizerInterface interface {
	Reset()
	Run(args ...any) (v any, err error)
}
