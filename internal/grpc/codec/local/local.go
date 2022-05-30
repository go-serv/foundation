package net

import (
	i "github.com/go-serv/service/internal"
)

type headerFlags i.HeaderFlags32Type

const (
	Encryption headerFlags = 1 << iota
)

type codec struct {
	i.CodecInterface
}

type dataframe struct {
	i.DataFrameInterface
	shmObjName   string
	shmBlockSize int
}
