package net

import (
	"github.com/go-serv/service/pkg/z"
)

type headerFlags z.HeaderFlags32Type

const (
	Encryption headerFlags = 1 << iota
)

type codec struct {
	z.CodecInterface
}
