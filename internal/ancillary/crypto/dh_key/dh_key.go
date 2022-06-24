package dh_key

import (
	"github.com/monnand/dhkx"
)

type DhKeyExchange interface {
	GetPublicKey() []byte
	ComputeKey([]byte) []byte
}

type dhKey struct {
	group   *dhkx.DHGroup
	privKey *dhkx.DHKey
}

// PublicKey returns a public key to exchange from a generated private one.
func (c *dhKey) PublicKey() []byte {
	return c.privKey.Bytes()
}

// ComputeKey computes a private key from the given public one.
func (c *dhKey) ComputeKey(pubKey []byte) ([]byte, error) {
	var (
		privKey *dhkx.DHKey
		err     error
	)
	if privKey, err = c.group.ComputeKey(dhkx.NewPublicKey(pubKey), c.privKey); err != nil {
		return nil, err
	}
	return privKey.Bytes(), nil
}
