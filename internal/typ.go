package internal

import "math/rand"

type UniqueIdInterface interface {
	Generate() UniqueId
}

func (u *UniqueId) Generate() {
	*u = UniqueId(rand.Uint64())
}

type UniqueId uint64

type SessionId UniqueId

type TransactionId UniqueId
