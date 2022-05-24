package ancillary

import (
	"fmt"
	"path"
	"runtime"
)

type MethodMustBeImplemented struct{}

// methodName returns the name of the calling method,
// assumed to be two stack frames above.
func methodName() string {
	pc, _, _, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	if f == nil {
		return "unknown method"
	}
	return f.Name()
}

func (m MethodMustBeImplemented) Panic() {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	basename := path.Base(file)
	panic(fmt.Sprintf("[%s:%d]: method '%s' must be implemented", basename, line, f.Name()))
}
