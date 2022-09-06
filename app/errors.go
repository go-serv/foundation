package app

import "errors"

var (
	ErrMethodDescriptorNotFound = errors.New("")
	ErrDescriptorNotFound       = errors.New("")
	ErrResolverNotFound         = errors.New("runtime: resolver not found")
)
