package z

type SessionInterface interface {
	Id() SessionId
	CryptoNonce() []byte
	WithCryptoNonce([]byte)
	EncKey() []byte
	WithEncKey([]byte)
}
