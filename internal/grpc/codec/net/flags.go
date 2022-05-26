package codec

import (
	i "github.com/go-serv/service/internal"
)

type headerFlags i.HeaderFlags32Type

const (
	Encryption headerFlags = 1 << iota
)
