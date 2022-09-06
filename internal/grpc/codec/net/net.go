package net

import (
	"github.com/go-serv/foundation/pkg/z"
)

type headerFlags z.HeaderFlagsType

const (
	Encryption headerFlags = 1 << iota
)

type codec struct {
	z.CodecInterface
}
