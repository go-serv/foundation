package service

import (
	"github.com/go-serv/foundation/internal/reflect"
	"github.com/go-serv/foundation/pkg/z"
)

var ref z.ReflectInterface

func init() {
	ref = reflect.NewReflection()
}
