package aes

import (
	"crypto/cipher"
)

type aesInfo struct {
	block cipher.Block
	gcm   cipher.AEAD
	nonce []byte
}

func (c *aesInfo) WithNonce(nonce []byte) {
	gcm, err := cipher.NewGCMWithNonceSize(c.block, len(nonce))
	if err != nil {
		panic(err)
	}
	c.gcm = gcm
}

func (c *aesInfo) Encrypt(data []byte) []byte {
	return c.gcm.Seal(data[:0], c.nonce, data, nil)
}

func (c *aesInfo) Decrypt(data []byte) ([]byte, error) {
	plain, err := c.gcm.Open(data[:0], c.nonce, data, nil)
	if err != nil {
		return nil, err
	} else {
		return plain, nil
	}
}
