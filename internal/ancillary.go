package internal

type DictionaryItem interface{}

type SymCipherInterface interface {
	WithNonce([]byte)
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}
