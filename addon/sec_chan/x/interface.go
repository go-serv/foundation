package x

import (
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/pkg/z/ancillary/crypto"
)

type DataFrameInterface interface {
	z.ApiKeyAwareInterface
	Parse(wire []byte) error
	Compose() ([]byte, error)
	WithBlockCipher(cipher crypto.AEAD_CipherInterface)
	Decrypt() error
	Payload() []byte
}
