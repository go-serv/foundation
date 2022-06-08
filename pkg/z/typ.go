package z

import "math/rand"

type UniqueIdInterface interface {
	Generate() UniqueId
}

func (UniqueId) Generate() UniqueId {
	return UniqueId(rand.Uint64())
}

type (
	UniqueId      uint64
	SessionId     UniqueId
	TransactionId UniqueId
)