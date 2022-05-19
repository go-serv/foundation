package ancillary

import (
	"fmt"
	"path"
	"runtime"
)

type MethodMustBeImplemented struct {
}

func (m MethodMustBeImplemented) Panic() {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	basename := path.Base(file)
	panic(fmt.Sprintf("[%s:%d]: method '%s' must be implemented", basename, line, f.Name()))
}
