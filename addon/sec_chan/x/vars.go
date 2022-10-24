package x

type KeyExchangeAlgoType int

const (
	KeyExchangeDH KeyExchangeAlgoType = iota
	KeyExchangeECDH
	KeyExchangeRSA
	KeyExchangePSK
)

type (
	secChanMwKey   struct{}
	pskResolverKey struct{}
)

var (
	PskResolverKey pskResolverKey // pre-shared key
	SecChanMwKey   secChanMwKey
)
