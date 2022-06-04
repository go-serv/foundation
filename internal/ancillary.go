package internal

type SymCipherInterface interface {
	WithNonce([]byte)
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}
