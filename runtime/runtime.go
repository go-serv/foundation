package runtime

import (
	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/pkg/ancillary/memoize"
)

func RegisterResolver(key any, resolver memoize.MemoizerInterface) {
	runtime.Runtime().RegisterResolver(key, resolver)
}

func Resolve(key any, args ...any) (v any, err error) {
	return runtime.Runtime().Resolve(key, args...)
}
