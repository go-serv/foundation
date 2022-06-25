package crypto

type AEAD_CipherInterface interface {
	WithNonce([]byte) error
	Encrypt(in []byte, additional []byte) []byte
	Decrypt(in []byte, additional []byte) (out []byte, err error)
}

type PubKeyExchangeInterface interface {
	ComputeKey(pubKey []byte) (privKey []byte, err error)
	PublicKey() []byte
}
