package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
)

func NewCipher(key []byte, nonce []byte) (*aesInfo, error) {
	// Normalize key length
	if len(key) != 32 {
		normalizedKey := sha256.Sum256(key)
		key = normalizedKey[:]
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// gcm or Galois/Counter Mode, is a mode of operation for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		return nil, err
	}
	return &aesInfo{block, gcm, nonce}, nil
}
