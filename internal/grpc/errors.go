package grpc

import "errors"

var (
	ErrInvalidProtoMessage = errors.New("grpc: not a proto.Message type")
)
