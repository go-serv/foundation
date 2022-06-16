package z

import "math/rand"

type (
	ErrCodeType int
)

type UniqueIdInterface interface {
	Generate() UniqueId
}

func (UniqueId) Generate() UniqueId {
	return UniqueId(rand.Uint64())
}

type (
	UniqueId  uint64
	SessionId UniqueId
	TenantId  UniqueId
)

const (
	_ = 1 << (iota * 10)
	KiB
	MiB
	GiB
)
