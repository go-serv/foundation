package aes

import (
	"crypto/cipher"
)

type aesInfo struct {
	block cipher.Block
	gcm   cipher.AEAD
	nonce []byte
}

func (c *aesInfo) WithNonce(nonce []byte) (err error) {
	c.gcm, err = cipher.NewGCMWithNonceSize(c.block, len(nonce))
	return
}

func (c *aesInfo) Encrypt(data []byte, additional []byte) []byte {
	return c.gcm.Seal(data[:0], c.nonce, data, additional)
}

func (c *aesInfo) Decrypt(data []byte, additional []byte) (plain []byte, err error) {
	plain, err = c.gcm.Open(data[:0], c.nonce, data, additional)
	return
}
